#!/bin/bash
# Logical PostgreSQL Backup Script
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
FILENAME="backup_${TIMESTAMP}.sql"
echo "Creating logical backup: ${FILENAME}..."
docker-compose exec postgres pg_dump -U appenheimer -d appenheimer -F c -f /tmp/${FILENAME}
docker-compose cp postgres:/tmp/${FILENAME} ./backups/${FILENAME}
docker-compose exec postgres rm /tmp/${FILENAME}
# Inject Metadata
echo "{\"id\": \"${TIMESTAMP}\", \"timestamp\": \"$(date -u -Iseconds)\", \"status\": \"completed\"}" > ./backups/${FILENAME}.meta.json
echo "Backup complete: ./backups/${FILENAME}"
