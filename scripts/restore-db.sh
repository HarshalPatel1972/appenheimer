#!/bin/bash
# PostgreSQL Restore Script
if [ -z "$1" ]; then
    echo "Usage: ./restore-db.sh <backup-file>"
    exit 1
fi
echo "Restoring database from $1..."
docker-compose cp $1 postgres:/tmp/restore.sql
docker-compose exec postgres pg_restore -U appenheimer -d appenheimer -c /tmp/restore.sql
docker-compose exec postgres rm /tmp/restore.sql
echo "Restore complete."
