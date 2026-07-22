#!/bin/bash
# Smoke Tests
curl -f -s http://localhost/health/ready > /dev/null || exit 1
curl -f -s http://localhost/api/v1/search?q=test > /dev/null || exit 1
echo "Smoke tests passed."
