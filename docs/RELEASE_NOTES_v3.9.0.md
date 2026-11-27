# Mistral Go SDK v3.9.0

This release ports the changes from the official Mistral Python SDK 3.4.0 through 3.9.0. Go 1.26 or newer is required. The module remains `github.com/ZaguanLabs/mistral-go/v3`.

- Manage webhooks, retrieve their signing secrets, inspect trigger events and manage trigger handlers. Billing and subscription requests use a dashboard context API key.
- Create and edit scorers, bulk-insert dataset records, discover pipeline result fields, tag workflows and request deployment archive upload URLs.
- Stream text into realtime speech synthesis and receive audio over the same WebSocket. Existing audio sessions can be updated with the REST API. Applications continue to provide their own voice media transport and device integration.
- Send single or batched multimodal embeddings and use the updated moderation, conversation, caching, retrieval and workflow models.
- Configure workflow payload compression, full or selective AES-GCM encryption, key rotation and blob offloading. Schedule creation signs canonical JSON with SHA-256 or HMAC. GCS supports a configurable API root and application default credentials. Other storage providers can implement `WorkflowBlobStore`.
- Configure an optional OTLP exporter per client or supply an application tracer provider. Dedicated exports support redaction and rotated service-account tokens. Application providers manage their own export pipeline.
- Batch submission preserves output and error files by default. Enable `CleanupResultFiles` to delete them after successful materialization. Retries respect the polling deadline.

The source was the official PyPI distribution `mistralai-3.9.0.tar.gz`, SHA-256 `7937d8c1422b1219d59943670960319508c471ab4cb0787d56cf8ec477b930f7`. The distribution was built with `python -m build` and the generated source archive was unpacked for comparison.

The offline race suite covers HTTP routes, WebSockets, telemetry export, GCS requests, and workflow encoding against Python-generated fixtures. Live Mistral, cloud and voice-device integrations were outside the test run.
