# Mistral Go SDK v2.10.1

Updated against the official Python SDK 2.10.1 source distribution.

- Managed indexes: create, list with page tokens, get, update schema, delete, ingest documents, delete documents, and search.
- Typed index schemas, field definitions, embedding configurations, boolean/range filters, keyword/vector retrievers, and response models. Search chunk extension fields are retained.
- Evaluation pipeline `definitions` lists and registered judge `slug`/`mapping` replace the previous single definition and inline judge fields.
- Initial service-account roles and organization-wide listing without a workspace filter.
- Workflow trace context is carried in both body and header. Explicit values win; absent values receive a sampled trace ID. Workflow events expose chain-run IDs.
- Agent ownership and connector MCP capability fields; offset voice pagination is deprecated upstream. Existing map responses retain new deployment, voice, and workflow trace metadata.

## Migration

Use `Definitions: []PipelineConfigDefinition{JudgeDefinition{Slug: "registered-judge"}}` in pipeline requests. Legacy `Definition` is wrapped into a list. Pipeline `Slug`, `Group`, and `DefinitionHash`, and judge `Model`/`Prompt`, are retained for source compatibility but no longer serialized.

## Validation

The official PyPI sdist was checksum verified, built with `python -m build`, and its generated sdist unpacked. Local HTTP contract tests, the full offline suite with the race detector, `go build ./...`, `go vet ./...`, and `git diff --check` pass. Credential-dependent live API tests were not run.
