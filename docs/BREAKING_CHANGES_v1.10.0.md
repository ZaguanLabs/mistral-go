# Breaking Changes in Python SDK v1.10.0 → Go SDK v2.1.0

## Overview

Python SDK v1.10.0 introduces significant breaking changes to response structures and field names. This document details all breaking changes and how they're addressed in the Go SDK v2.1.0.

## Breaking Changes Summary

### 1. Library API Changes

**Changed Fields:**
- `created` → `created_at` (timestamp field renamed)
- `updated` → `updated_at` (timestamp field renamed)
- Added: `owner_id` (nullable string)
- Added: `owner_type` (string)
- Added: `total_size` (int)
- Added: `nb_documents` (int)
- Added: `chunk_size` (nullable int)
- Added: `emoji` (optional nullable string)
- Added: `generated_description` (optional nullable string)
- Added: `explicit_user_members_count` (optional nullable int)
- Added: `explicit_workspace_members_count` (optional nullable int)
- Added: `org_sharing_role` (optional nullable string)
- Added: `generated_name` (optional nullable string)

**Impact:** All library operations (create, get, update, delete, list)

### 2. Document API Changes

**Changed Fields:**
- `created` → `created_at` (timestamp field renamed)
- `updated` → `updated_at` (timestamp field renamed)
- `status` → `processing_status` (field renamed)
- `size` → nullable (was required, now optional)
- Added: `hash` (nullable string)
- Added: `mime_type` (nullable string)
- Added: `extension` (nullable string)
- Added: `uploaded_by_id` (nullable string)
- Added: `uploaded_by_type` (string)
- Added: `tokens_processing_total` (int)
- Added: `summary` (optional nullable string)
- Added: `last_processed_at` (optional nullable timestamp)
- Added: `number_of_pages` (optional nullable int)
- Added: `tokens_processing_main_content` (optional nullable int)
- Added: `tokens_processing_summary` (optional nullable int)
- Added: `url` (optional nullable string)
- Added: `attributes` (optional map[string]any)

**New Request Fields:**
- `UpdateDocumentRequest.attributes` (optional map[string]any)
- `ListDocuments.filters_attributes` (optional filter parameter)

**Impact:** All document operations (get, list, update, upload)

### 3. Files API Changes

**Changed Fields:**
- `total` → optional/nullable (was required, now optional in ListFilesOut)

**New Request Fields:**
- `ListFiles.include_total` (optional boolean parameter)

**Impact:** ListFiles operation

### 4. Chat API Changes

**New Request Fields:**
- `ChatCompletionRequest.metadata` (optional map[string]any)

**Impact:** Chat completion operations

### 5. FIM API Changes

**New Request Fields:**
- `FIMCompletionRequest.metadata` (optional map[string]any)

**Impact:** FIM completion operations

### 6. Agents API Changes

**New Request Fields:**
- `AgentCompletionRequest.metadata` (optional map[string]any) - Already added
- `CreateAgentRequest.metadata` (optional map[string]any)

**New Operations:**
- `DeleteAgent` - Already added

**New Request Parameters:**
- `GetAgent.agent_version` (optional string)
- `ListAgents` - request structure changed (sources filter)

**Response Changes:**
- Agent responses now include more fields

### 7. Conversations API Changes

**New Operations:**
- `DeleteConversation` - Already added

**New Request Fields:**
- `ListConversations.metadata` (optional map[string]any filter)

**Response Changes:**
- Conversation responses structure changed
- `tool_execution_entry.name` field changed

### 8. Model Capabilities Changes

**New Fields:**
- `classification` (optional bool)
- `moderation` (optional bool)
- `audio` (optional bool)

### 9. Accesses API Changes

**Changed Fields:**
- `org_id` parameter changed
- `share_with_uuid` response field changed

**Impact:** All access operations (list, update_or_create, delete)

## Migration Strategy

Due to the extensive breaking changes, we need to:

1. Update all type definitions
2. Fix all tests to use new field names
3. Update documentation
4. Provide clear migration guide for users
5. Consider version bump to v3.0.0 instead of v2.1.0

## Recommendation

Given the extensive breaking changes, this should be:
- **Go SDK v3.0.0** (major version bump)
- Clear migration guide from v2.x to v3.0.0
- Deprecation notices for old field names

## Status

- ✅ RequestSource type added
- ✅ OCRTableObject added
- ✅ Agent metadata parameter added
- ✅ Delete operations added
- 🔄 Library type - partially updated (needs field name changes)
- 🔄 Document type - partially updated (needs field name changes)
- 🔄 Files API - total field updated
- ❌ Chat metadata - not yet added
- ❌ FIM metadata - not yet added
- ❌ Model capabilities - not yet updated
- ❌ Conversations changes - not yet updated
- ❌ Accesses changes - not yet updated
- ❌ Tests - need extensive updates
