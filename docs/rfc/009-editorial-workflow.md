# RFC 009 - Editorial Workflow

## 1. Context
As Appenheimer shifts from an automated scraper mindset to a high-quality curation platform, the Admin System becomes the operational heart of the product. Before building the UI or API routes, we must definitively map the lifecycle of an application from Ingestion to Publication. This document outlines the state machine, permissions, and operational consequences of the editorial workflow.

## 2. The Editorial State Machine
We expand the `visibility_status` ENUM to cleanly capture the editorial funnel:

```text
[Draft] → [Review] → [Needs Changes] → [Published] → [Archived]
```

- **Draft**: The default state for all newly ingested YAML records or manually created entries.
- **Review**: The record is ready for a senior editor/admin to review.
- **Needs Changes**: The review failed; pushed back to the submitter's queue.
- **Published**: Live on the public platform. Publishing is itself the final approval.
- **Archived**: Soft-deleted or permanently retired from public view.

## 3. Operational Rules

### Roles & Permissions
We define four strict roles to govern the platform:
1. **Viewer**: Read-only access to internal systems.
2. **Contributor**: Can create and edit `Draft` records, and submit them for `Review`.
3. **Editor**: Can review drafts, push to `Needs Changes`, or `Publish`.
4. **Administrator**: Full system access, including taxonomy management and user roles.

### Editing Published Applications (Shadow Drafts)
- **Can published apps be edited directly?** No.
- **What happens when editing a published app?** A "Shadow Draft" is generated. This creates a duplicate record in the `apps` table with `visibility_status = 'draft'` and a `source_app_id` linking back to the live Published app. Once the Shadow Draft is published, its values overwrite the Published record, and the Shadow Draft is archived.
- **Limit**: A published app can have **only one active shadow draft**. If an editor attempts to edit a published app that already has a shadow draft, they are seamlessly redirected to the existing shadow draft.

### Concurrency, Audit, and Locking
- **Optimistic Locking**: The `version BIGINT` column on the `apps` table strictly prevents simultaneous overwrites.
- **Auditability**: Every editable entity includes an `updated_by (UUID)` and `updated_at` timestamp. 

### Publishing Atomicity
Publishing is strictly **one transaction**:
1. Validate requirements.
2. Update database records (e.g. overwriting the live app and archiving the shadow draft).
3. Commit transaction.
Only *after* the database commit succeeds does the system emit a `PublishEvent`.

### History and Reversions
Revisions are first-class citizens. When we reach Phase 3, we will introduce robust relational tables for `app_revisions` rather than brittle JSON snapshots.

## 4. Platform Side-Effects

### Meilisearch & The Transactional Outbox
We implement the **Transactional Outbox Pattern** to ensure absolute consistency between PostgreSQL and Meilisearch.
- **Workflow**: `Publish Transaction → outbox_events Table → Worker → Meilisearch`.
- **Benefits**: This guarantees zero lost events, allows robust retries, prevents inconsistent search states if Meilisearch is down, and easily scales to support analytics webhooks later.

## 5. Required Schema Alterations
1. Expand `visibility_status` ENUM (`review`, `needs_changes`).
2. Add `source_app_id UUID (Nullable)` to the `apps` table.
3. Add `updated_by UUID (Nullable)` to the `apps` table.
4. Create an `outbox_events` table for the Transactional Outbox pattern.
