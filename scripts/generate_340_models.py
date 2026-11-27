#!/usr/bin/env python3
"""Generate Go models for a pair of unpacked Mistral Python SDK releases."""
import ast,re,json,subprocess,sys
from pathlib import Path
repo = Path(__file__).resolve().parent.parent
previous, version = sys.argv[1:] or ("3.2.0", "3.4.0")
source = repo / f"docs/client-python/mistralai-{version}/src/mistralai/client/models"
old = repo / f"docs/client-python/mistralai-{previous}/src/mistralai/client/models"
if not source.is_dir() or not old.is_dir():
    sys.exit("Both unpacked reference SDKs are required")
output = repo / "sdk" / f"models_{version.replace('.', '')}_generated.go"
existing = set(re.findall(r'^type (\w+)', '\n'.join(
    p.read_text() for p in (repo / 'sdk').glob('*.go')
    if not p.name.endswith('_test.go') and p != output
), re.M))
classes={};aliases={}
for f in sorted(source.glob('*.py')):
 if f.name.endswith('op.py'):continue
 if (old/f.name).exists() and (old/f.name).read_bytes()==f.read_bytes():continue
 for n in ast.parse(f.read_text()).body:
  if isinstance(n,ast.ClassDef) and any(ast.unparse(b)=='BaseModel' for b in n.bases) and n.name not in existing:classes[n.name]=n
  if isinstance(n,ast.Assign) and len(n.targets)==1 and isinstance(n.targets[0],ast.Name):
   name=n.targets[0].id
   if name[0].isupper() and not name.endswith('TypedDict') and name not in existing:aliases[name]=n.value
if version != '3.4.0':
    # Boilerplate serializer changes affect hundreds of upstream files. Select
    # changed wire schemas and the model definitions they reference.
    def schema(node):
        if isinstance(node, ast.ClassDef):
            return [(field.target.id, ast.dump(field.annotation), ast.dump(field.value) if field.value else None)
                    for field in node.body if isinstance(field, ast.AnnAssign) and isinstance(field.target, ast.Name)]
        return ast.dump(node)

    def definitions(folder):
        result = {}
        for file in sorted(folder.glob('*.py')):
            if file.name.endswith("op.py"):
                continue
            for node in ast.parse(file.read_text()).body:
                if isinstance(node, ast.ClassDef) and any(ast.unparse(base) == 'BaseModel' for base in node.bases):
                    result[node.name] = node
                elif isinstance(node, ast.Assign) and isinstance(node.targets[0], ast.Name):
                    name = node.targets[0].id
                    if name[0].isupper() and not name.endswith('TypedDict'):
                        result[name] = node.value
        return result

    before, after = definitions(old), definitions(source)
    selected = {name for name, node in after.items()
                if name not in before or schema(node) != schema(before[name])}
    # Run inputs accept every member of this upstream discriminated union.
    selected.add("InputEntries")
    pending = list(selected)
    while pending:
        node = after[pending.pop()]
        for ref in ast.walk(node):
            if isinstance(ref, ast.Name) and ref.id in after and ref.id not in selected and ref.id not in existing:
                selected.add(ref.id)
                pending.append(ref.id)
    funcs = set(re.findall(r'^func (\w+)\(', '\n'.join(p.read_text() for p in (repo/'sdk').glob('*.go')), re.M))
    selected -= existing | funcs
    classes = {name: after[name] for name in sorted(selected) if isinstance(after[name], ast.ClassDef)}
    aliases = {name: after[name] for name in sorted(selected) if not isinstance(after[name], ast.ClassDef)}
known=existing|classes.keys()|aliases.keys()
def args(n):return list(n.slice.elts) if isinstance(n.slice,ast.Tuple) else [n.slice]
def typ(n):
 if isinstance(n,ast.Name):return {'str':'string','int':'int','float':'float64','bool':'bool','datetime':'time.Time','Any':'any','UnrecognizedStr':'string'}.get(n.id,n.id if n.id in known else 'any')
 if isinstance(n,ast.Subscript):
  name=ast.unparse(n.value);a=args(n)
  if name in ['Annotated','SerializeAsAny']:return typ(a[0])
  if name in ['Optional','OptionalNullable','Nullable']:
   t=typ(a[0]);return t if t.startswith(('[]','map[')) or t=='any' else '*'+t
  if name=='List':return '[]'+typ(a[0])
  if name=='Dict':return 'map['+typ(a[0])+']'+typ(a[1])
  if name=='Literal':return 'bool' if isinstance(a[0],ast.Constant) and isinstance(a[0].value,bool) else 'int' if isinstance(a[0],ast.Constant) and isinstance(a[0].value,int) else 'string'
  if name=='Union':
   texts=[ast.unparse(x) for x in a]
   return 'string' if any(x.startswith('Literal[') for x in texts) and all(x.startswith('Literal[') or x=='UnrecognizedStr' for x in texts) else 'any'
 if isinstance(n,ast.Call) and ast.unparse(n.func)=='TypeAliasType':return typ(n.args[1])
 return 'any'
def fieldname(s):return ''.join({'id':'ID','url':'URL','ipa':'IPA','api':'API','otlp':'OTLP','ns':'NS'}.get(x,x[:1].upper()+x[1:]) for x in s.strip('_').split('_'))
out=[f'// Code generated from official Mistral Python SDK {version} models. DO NOT EDIT.','package sdk','import ("encoding/json"; "time")']
if version != '3.4.0':out[2]='import ("encoding/json"; "fmt"; "time")'
for name,n in aliases.items():out.append(f'type {name} = {typ(n)}')
for name,c in classes.items():
 out.append(f'type {name} struct {{');defaults=[];collections=[];wire_names=[];nullable=[]
 extras=any(isinstance(n,ast.AnnAssign) and isinstance(n.target,ast.Name) and n.target.id=='__pydantic_extra__' for n in c.body)
 if extras:out.append('AdditionalProperties map[string]any `json:"-"`')
 for n in c.body:
  if not isinstance(n,ast.AnnAssign) or not isinstance(n.target,ast.Name) or n.target.id.startswith('_'):continue
  wire=n.target.id.rstrip('_');text=ast.unparse(n.annotation)
  m=re.search(r"alias=['\"]([^'\"]+)",text)
  if m:wire=m[1]
  wire_names.append(wire)
  t=typ(n.annotation);fn=fieldname(n.target.id)
  optional=n.value is not None
  if version != '3.4.0' and ('Request' in name or name.startswith('Create') or name in ('RealtimeSpeechSessionUpdatePayload','AudioSessionTurnUpdate')) and 'OptionalNullable[' in text:
   nullable.append(wire)
  if optional and t.startswith(('[]','map[')):collections.append((fn,wire))
  out.append(f'{fn} {t} `json:"{wire}{",omitempty" if optional else ""}"`')
  if isinstance(n.value,ast.Constant) and n.value.value is not None:
   val=n.value.value
   if t.startswith('*'):
    literal=json.dumps(val,ensure_ascii=False)
    defaults.append(f'if v.{fn}==nil {{x:={t[1:]}({literal});v.{fn}=&x}}')
   elif t=='string' and isinstance(val,str):defaults.append(f'if v.{fn}=="" {{v.{fn}={json.dumps(val)}}}')
 if nullable:out.append('ExplicitNulls []string `json:"-"`')
 out.append('}')
 if defaults or collections or extras or nullable:
  code=f'func (v {name}) MarshalJSON() ([]byte,error) {{type wire {name};'+';'.join(defaults)+';'
  if collections or extras or nullable:
   code+='data,err:=json.Marshal(wire(v));if err!=nil {return nil,err};var fields map[string]json.RawMessage;if err:=json.Unmarshal(data,&fields);err!=nil {return nil,err};'
   for fn,key in collections:code+=f'if v.{fn}!=nil {{data,err:=json.Marshal(v.{fn});if err!=nil {{return nil,err}};fields["{key}"]=data}};'
   if extras:
    protected=','.join(json.dumps(key) for key in wire_names)
    code+='for key,value:=range v.AdditionalProperties {'+('switch key {case '+protected+':continue};' if protected else '')+'data,err:=json.Marshal(value);if err!=nil {return nil,err};fields[key]=data};'
   if nullable:
    allowed=','.join(json.dumps(key) for key in nullable)
    code+=f'for _,key:=range v.ExplicitNulls {{switch key {{case {allowed}:fields[key]=json.RawMessage("null");default:return nil,fmt.Errorf("field %q does not accept explicit null",key)}}}};'
   code+='return json.Marshal(fields)}'
  else:code+='return json.Marshal(wire(v))}'
  out.append(code)
 if extras:
  code=f'func(v *{name}) UnmarshalJSON(data []byte) error {{type wire {name};var decoded wire;if err:=json.Unmarshal(data,&decoded);err!=nil {{return err}};var fields map[string]json.RawMessage;if err:=json.Unmarshal(data,&fields);err!=nil {{return err}};'
  for key in wire_names:code+='delete(fields,'+json.dumps(key)+');'
  code+=f'*v={name}(decoded);if len(fields)>0 {{v.AdditionalProperties=map[string]any{{}};for key,raw:=range fields {{var value any;if err:=json.Unmarshal(raw,&value);err!=nil{{return err}};v.AdditionalProperties[key]=value}}}};return nil}}'
  out.append(code)
output.write_text('\n'.join(out)+'\n')
print('Generated',len(classes),'models and',len(aliases),'aliases')

subprocess.run(["gofmt","-w",str(output)],check=True)
