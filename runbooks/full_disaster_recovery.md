# Full Disaster Recovery (Cluster Revival)

## Objective
Restore the entire Appenheimer platform from catastrophic failure (total infrastructure loss).

## Procedure Sequence
1. Provision fresh virtual machines/cluster nodes.
2. Pull latest codebase and Docker Compose definitions.
3. Download the latest PostgreSQL backup archive from cold storage.
4. Run `./scripts/restore-db.sh <backup-file>`.
5. Start PostgreSQL and Redis: `docker-compose up -d postgres redis`.
6. Run Meilisearch Rebuild: `./scripts/rebuild-search.sh`.
7. Start Frontend, Backend, Worker, Caddy: `./scripts/start.sh`.
8. Execute Health Probes: `./scripts/health.sh`.

## Post-Mortem
Once the site is stable, declare an incident post-mortem to analyze the root cause of the initial infrastructure loss.
