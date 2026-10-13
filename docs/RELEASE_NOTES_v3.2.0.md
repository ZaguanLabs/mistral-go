# Mistral Go SDK v3.2.0

Targets official Python SDK 3.2.0; module imports remain `/v3`.

- Adds evaluation project/list/get APIs, v2 evaluations, run search/get, record search/get and filtered run statistics.
- Adds optimization list/get/trials, pipeline-result search and `/v1/whoami`.
- Adds generated typed evaluation models, filters, statistics, sampling strategies, resource chunks and principal variants. Open unions retain native Go `any`/map values and arbitrary future variants.
- Adds min tokens, repetition penalty, top-k and completion/prompt logprob controls to chat and agents, including streamed calls and typed logprob responses.
- Adds speech IPA phoneme hints, voice kind and workflow registration name filtering. Dataset/workflow response maps retain added fields.
- Stops sending removed connector auth-direction and system-prompt-route fields. OAuth server metadata is already optional in Go; native MCP sessions return caller errors directly without Python task groups.

## Migration

`UpdatePipelineRequest` fields are optional pointers: `Name: StringPtr("name")`, `Enabled: BoolPtr(false)`, and pointers to selector/definition slices. Nil leaves a field unchanged; a pointer to an empty slice clears the list. Judge definitions accept `Sampling` (probabilistic or attribute-hash strategies). Pipeline kinds use `detection-v1`, `moderation-v1`, `judge-v1` and `export-v1`.

Standalone pipeline-config endpoints removed by Python remain deprecated legacy Go compatibility methods. Use pipeline CRUD for current operations.

## Source and validation

Official PyPI sdist SHA-256: `b3f2139bde538fe8ed4a19e2adb624fd46dfd175b609231c0836bba21d738730`. Pristine source was built with `python -m build` and the generated sdist unpacked under ignored `docs/client-python/mistralai-3.2.0/dist/`.

Reviewed public method and model contracts against 3.0.0. Tests cover v2 endpoint paths, pagination, escaped IDs, repeated optimization tags, zero/false chat parameters, response logprobs, sampling discriminators and nested run metadata. Complete offline suite passes with the race detector, build, vet and diff checks. Live credential-dependent API tests were not run.
