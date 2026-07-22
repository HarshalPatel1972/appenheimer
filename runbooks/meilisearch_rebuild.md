# Meilisearch Rebuild Runbook

## Objective
Rehydrate the Meilisearch index entirely from the PostgreSQL source of truth. Meilisearch is a derived projection and its volumes can be safely discarded during an outage.

## Procedure
1. If Meilisearch is corrupted, destroy its container and volume:
   `docker-compose rm -f -v meilisearch`
2. Spin up a fresh Meilisearch instance:
   `docker-compose up -d meilisearch`
3. Trigger the `SyncAllApps` CLI command to bulk rebuild the projection:
   `./scripts/rebuild-search.sh`
4. Wait for the task to complete (monitor logs: `docker-compose logs -f worker`).

## Validation
- Execute a search query on the frontend and verify populated results.
- `GET /metrics` shows index latency dropping to normal.
