# Mistral Go SDK v3.4.0

Updates the Go API to the official Python SDK 3.4.0 schemas and service routes.

- Audio agents: create, list, retrieve, update, delete, versions and sessions.
- Audio sessions: create, retrieve and close; native event/transcript/function-handler lifecycle helpers.
- Responses: create, SSE streaming and input-token counting, including open extension fields and explicit empty collections.
- Managed indexes: current `/v1/context/search-indexes` routes and summary retrieval, update and JSONL generation streams.
- Audio: realtime diarization, context bias, language, speaker references and turn detection; transcription speaker references; speech expressiveness and conversation history.
- Models and workflows: current audio capabilities, reasoning usage, retriever defaults, scorer selectors and deployment repository/Dockerfile filters.

Voice media and device integration use application-supplied native `VoiceAgentTransport` and `VoiceAudioInterface` implementations. The Python LiveKit/PortAudio adapters, decorator-based signature inference and Python worker monkeypatch are language/runtime-specific and are not bundled. Local functions declare their parameter schema explicitly. Worker tokens use encrypted endpoints except on loopback. The service API tests do not verify live media interoperability.

Source: official PyPI `mistralai-3.4.0.tar.gz`, SHA-256 `16f87c9826f2a7dcb21b2c59a22db948344791dcdb230771cd938a85e5d9ef42`. Built with `python -m build` and the generated sdist unpacked under ignored `docs/client-python/`.

Validation: offline `go test -race ./...`, `go vet ./...` and `go build ./...`. Eleven existing unconditional live API tests were excluded after an initial run returned HTTP 401; no usable live API credentials were available. Mock HTTP, SSE/JSONL, websocket payload, voice lifecycle, transcript replacement, handler timeout and credential transport checks pass.
