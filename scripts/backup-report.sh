#!/bin/bash
# Print Backup Inventory
echo "=== Appenheimer Backup Inventory ==="
ls -l ./backups/*.sql 2>/dev/null || echo "No backups found."
