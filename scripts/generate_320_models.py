import ast,re,json
from pathlib import Path
repo=Path(__file__).resolve().parent.parent; source=repo/'docs/client-python/mistralai-3.2.0/src/mistralai/client/models';old=repo/'docs/client-python/mistralai-3.0.0/src/mistralai/client/models'
existing=set(re.findall(r'^type (\w+)', '\n'.join(p.read_text() for p in (repo/'sdk').glob('*.go') if not p.name.endswith('_test.go') and p.name != 'models_320_generated.go'),re.M))
classes={};aliases={}
for f in sorted(source.glob('*.py')):
 if (old/f.name).exists() or f.name.endswith('op.py'):continue
 for n in ast.parse(f.read_text()).body:
  if isinstance(n,ast.ClassDef) and any(ast.unparse(b)=='BaseModel' for b in n.bases) and n.name not in existing:classes[n.name]=n
  if isinstance(n,ast.Assign) and len(n.targets)==1 and isinstance(n.targets[0],ast.Name):
   name=n.targets[0].id
   if name[0].isupper() and not name.endswith('TypedDict') and name not in existing:aliases[name]=n.value
known=existing|classes.keys()|aliases.keys()
def args(n):return list(n.slice.elts) if isinstance(n.slice,ast.Tuple) else [n.slice]
def typ(n):
 if isinstance(n,ast.Name):return {'str':'string','int':'int','float':'float64','bool':'bool','datetime':'time.Time','Any':'any'}.get(n.id,n.id if n.id in known else 'any')
 if isinstance(n,ast.Subscript):
  name=ast.unparse(n.value);a=args(n)
  if name in ['Annotated','SerializeAsAny']:return typ(a[0])
  if name in ['Optional','OptionalNullable','Nullable']:
   t=typ(a[0]);return t if t.startswith(('[]','map[')) or t=='any' else '*'+t
  if name=='List':return '[]'+typ(a[0])
  if name=='Dict':return 'map['+typ(a[0])+']'+typ(a[1])
  if name=='Literal':return 'bool' if isinstance(a[0],ast.Constant) and isinstance(a[0].value,bool) else 'string'
  if name=='Union':
   texts=[ast.unparse(x) for x in a]
   return 'string' if any(x.startswith('Literal[') for x in texts) and all(x.startswith('Literal[') or x=='UnrecognizedStr' for x in texts) else 'any'
 if isinstance(n,ast.Call) and ast.unparse(n.func)=='TypeAliasType':return typ(n.args[1])
 return 'any'
def fieldname(s):return ''.join({'id':'ID','url':'URL','ipa':'IPA','api':'API','otlp':'OTLP','ns':'NS'}.get(x,x[:1].upper()+x[1:]) for x in s.strip('_').split('_'))
out=['// Code generated from official Mistral Python SDK 3.2.0 models. DO NOT EDIT.','package sdk','import ("encoding/json"; "time")']
for name,n in aliases.items():out.append(f'type {name} = {typ(n)}')
for name,c in classes.items():
 out.append(f'type {name} struct {{');defaults=[];collections=[]
 for n in c.body:
  if not isinstance(n,ast.AnnAssign) or not isinstance(n.target,ast.Name):continue
  wire=n.target.id.rstrip('_');text=ast.unparse(n.annotation)
  m=re.search(r"alias=['\"]([^'\"]+)",text)
  if m:wire=m[1]
  t=typ(n.annotation);fn=fieldname(n.target.id)
  optional=n.value is not None
  if optional and t.startswith(('[]','map[')):collections.append((fn,wire))
  out.append(f'{fn} {t} `json:"{wire}{",omitempty" if optional else ""}"`')
  if isinstance(n.value,ast.Constant) and n.value.value is not None:
   val=n.value.value
   if t.startswith('*'):defaults.append(f'if v.{fn}==nil {{x:={('float64('+str(val)+')') if t=='*float64' else json.dumps(val,ensure_ascii=False)};v.{fn}=&x}}')
   elif t=='string' and isinstance(val,str):defaults.append(f'if v.{fn}=="" {{v.{fn}={json.dumps(val)}}}')
 out.append('}')
 if defaults or collections:
  code=f'func (v {name}) MarshalJSON() ([]byte,error) {{type wire {name};'+';'.join(defaults)+';'
  if collections:
   code+='data,err:=json.Marshal(wire(v));if err!=nil {return nil,err};var fields map[string]json.RawMessage;if err:=json.Unmarshal(data,&fields);err!=nil {return nil,err};'
   for fn,key in collections:code+=f'if v.{fn}!=nil {{data,err:=json.Marshal(v.{fn});if err!=nil {{return nil,err}};fields["{key}"]=data}};'
   code+='return json.Marshal(fields)}'
  else:code+='return json.Marshal(wire(v))}'
  out.append(code)
(repo/'sdk/models_320_generated.go').write_text('\n'.join(out)+'\n')
print('Generated',len(classes),'models and',len(aliases),'aliases')
