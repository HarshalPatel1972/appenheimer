# RFC 010 - Event Processing Architecture

## 1. Context
As Appenheimer transitions from a monolithic prototype into a scalable production platform, we must reliably decouple core database mutations (like publishing an app) from their expensive, asynchronous side-effects (like indexing search, sending emails, or triggering analytics). 

This RFC establishes a generic, resilient **Event Processing Framework** built entirely atop PostgreSQL using the Transactional Outbox pattern. 

## 2. Event Schema
We retain the **`outbox_events`** table name. The explicit naming avoids ambiguity when future architectural concepts (e.g., `inbox_events`, `domain_events`) are introduced.

The schema is:
- `id` (UUID, Primary Key)
- `type` (VARCHAR, e.g., `app.published`)
- `event_version` (INT, enables payload schema evolution without breaking consumers)
- `aggregate_id` (UUID, the ID of the mutated entity, e.g., App ID)
- `payload` (JSONB, fully self-contained immutable state so consumers rarely query the DB)
- `status` (VARCHAR)
- `attempts` (INT, default 0)
- `last_error` (TEXT, nullable)
- `next_attempt_at` (TIMESTAMPTZ, controls exponential backoff logic)
- `locked_at` (TIMESTAMPTZ, nullable)
- `locked_by` (VARCHAR, nullable, e.g., the worker pod ID)
- `created_at` (TIMESTAMPTZ)
- `updated_at` (TIMESTAMPTZ)

## 3. The Event Lifecycle
Every event iterates through a strict, simplified state machine:

```text
[Pending] → [Processing] → [Completed]
                  ↓
                [Dead]
```

- **Pending**: The event was created inside a transaction, or a transient failure occurred and `next_attempt_at` is set in the future.
- **Processing**: A worker has acquired a lock on the event.
- **Completed**: The consumer(s) successfully processed the event.
- **Dead**: The event exceeded the maximum retry limit (e.g., 5 attempts) and requires manual administrator intervention.

## 4. Worker Rules & Architecture
A background Go worker polls the `outbox_events` table.

- **Polling**: Continuously fetches events `WHERE status = 'pending' AND next_attempt_at <= NOW()`.
- **Concurrency & Locking**: 
  - `SELECT ... FOR UPDATE SKIP LOCKED` guarantees safe concurrent execution across multiple worker instances.
  - Updates set `status = 'processing'`, `locked_at = NOW()`, `locked_by = WorkerID`.
- **Lock Recovery**: If a worker crashes mid-processing, an event stays locked. Another routine will automatically reclaim locks where `locked_at` is older than a configurable threshold (e.g., 5 minutes) and revert them to `Pending`.
- **Event Ordering**: Events are strictly ordered by `created_at ASC` **per aggregate**. 
  - `app.updated` strictly executes after `app.published` for App A. 
  - Concurrently, App B can process entirely asynchronously. 

## 5. Processing Guarantees
We explicitly guarantee **At-Least-Once Delivery**. Every downstream consumer (Meilisearch, Analytics, Email) **MUST** be idempotent. 

## 6. Future Consumers
Workers explicitly identify their identity in logs (e.g., `MeilisearchWorker`, `AnalyticsWorker`). This guarantees that as the platform scales, logs remain highly actionable.

## 7. Recovery & Metrics
To maintain platform health:
- **Replay**: Replaying events does *not* reset completed rows. Instead, it generates a fresh synthetic event in the outbox (e.g., `Reindex App -> INSERT new app.published`) to preserve a perfectly immutable audit trail.
- **Metrics**: The framework natively exposes:
  - Total Pending, Processing, Completed, Dead
  - Average Processing Time
  - Retries
  - Queue Depth
  - Oldest Pending Event Age
