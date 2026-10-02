# Mistral Go SDK v3.0.0

Updated against the official Python SDK 3.0.0 source distribution.

## New functionality

- Service-account authentication reads `MISTRAL_SA_TOKEN_PATH` on each request, including realtime handshakes. Caller authorization and explicit API keys take precedence; invalid configured token files fail before the request is sent.
- Connector HTTP clients enforce the configured origin and connector path, disable redirects, attach optional credential names, and expose gateway-authored failures through `ConnectorsGatewayError`. Connector MCP sessions use the official Go MCP SDK and initialize before being returned.
- Managed indexes add name/status/creator filters, navigation, source reading, grep, chunk retrieval, reciprocal-rank-fusion retrievers, and nearest-neighbour candidate limits.
- Observability adds pipeline CRUD and pagination, dataset imports from spans with mapping contracts and import counters, span-evaluation aggregation, and raw attribute keys in field definitions.
- Workflows add deployment unharden, owner filtering, build/runtime log selection, and Mistral Cloud backend build settings and secret bindings.
- Voice search supports page tokens, repeated gender/language filters, and query text. Service-account listing adds search and sort order.
- Structured completion parsing concatenates text chunks and preserves the original mixed content. Empty or non-text-only responses leave the parsed result unset.

## Migration

Use Go 1.25 or newer and change imports to `github.com/ZaguanLabs/mistral-go/v3/sdk`. The minimum Go version follows the official MCP Go SDK dependency. Existing `/v2` consumers retain the 2.x API.

Pipeline-config requests again use singular `Definition`. A legacy one-element `Definitions` list is converted; multiple definitions must use the new pipeline APIs. Pipeline responses retain their individual configurations.

`GetConnector(id, fetchUserData)` replaces the obsolete customer-data and connection-secrets arguments. Tool listing no longer sends `page`; connector query filters no longer send `active`. `DeleteConnectorCredentials` accepts the user, workspace, or organization scope.

Deployment build settings belong in `DeploymentMistralCloudBackendSpec`, including `BuildDirectory`, `DockerfilePath`, and `Secrets`. Deprecated top-level entrypoint, working-directory, CPU/memory, and response commit-SHA fields are no longer serialized. Use the structured commit metadata for commit information.

Chat and agent completion requests accept function, image-generation, document-library, and connector tools. Removed web-search and code-interpreter variants now return a validation error; use connector tools for these services. Other agent APIs retain their own tool contracts.

HTTP connector callers close response bodies and call `Close` when finished. MCP callers close the returned `mcp.ClientSession`; Go contexts provide cancellation. Gateway request timeouts follow the configured Go client timeout.

## Source and validation

Official source: [mistralai 3.0.0 on PyPI](https://pypi.org/project/mistralai/3.0.0/).

The downloaded `mistralai-3.0.0.tar.gz` SHA-256 is `60789d0efd23f779d18922e4a6a4f54af428dff05deab0be9fc963acf7160f7f`. The source was built unmodified with `python -m build`; its generated source archive was unpacked beneath `docs/client-python/mistralai-3.0.0/dist/`. Downloaded sources and build artifacts remain local and ignored by Git.

Source comparison covered public resource changes, generated model changes, authentication hooks, connector helpers, and structured-response parsing. Existing map-based responses preserve added deployment and workflow metadata. Python-only transport/import changes map to the native Go HTTP and MCP implementations.

Validation: local HTTP contract tests, token precedence and rotation tests, gateway boundary/error tests, an initialized local MCP session, the complete offline suite with the race detector, `go build ./...`, `go vet ./...`, and `git diff --check`. Credential-dependent live API tests were not run.
