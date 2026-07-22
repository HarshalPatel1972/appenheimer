#!/bin/bash
# Trigger Meilisearch Rebuild from PostgreSQL Projection Layer
echo "Triggering SyncAllApps CLI command..."
docker-compose exec worker /api -cli sync-all-apps
echo "Meilisearch rehydration queued in the Outbox."
