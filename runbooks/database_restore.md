# PostgreSQL Database Restore Runbook

## Objective
Restore the primary database from a logical backup (`pg_dump`) or perform Point-in-Time Recovery (PITR) via WAL archives.
**RTO:** 30 minutes | **RPO:** 5 minutes

## Prerequisites
- Access to production environment via SSH.
- Located the desired backup archive in off-site storage (e.g., S3).
- Verified `APPENHEIMER_DB_PASSWORD`.

## Procedure (Logical Restore)
1. Stop the application containers to prevent split-brain writes:
   `./scripts/stop.sh`
2. Wipe the existing database volume and recreate the Postgres container:
   `docker-compose rm -f -v postgres`
   `docker-compose up -d postgres`
3. Execute the restore script against the fresh database:
   `./scripts/restore-db.sh /path/to/backup.sql`
4. Run the integrity validation script:
   `./scripts/verify-backup.sh`
5. If validation passes, restart the application cluster:
   `./scripts/start.sh`

## Validation Checks
- Output of `./scripts/health.sh` reports `postgres` as `up`.
- Application loads without 500 errors.
