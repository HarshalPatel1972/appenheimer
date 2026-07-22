# Redis Recovery Runbook

## Objective
Recover from a Redis cluster failure. Redis operates entirely as a volatile cache and rate limiter. No permanent state is stored here.

## Procedure
1. Restart the Redis container:
   `docker-compose restart redis`
2. Or wipe it if corrupted:
   `docker-compose rm -f -v redis && docker-compose up -d redis`

## Impact Expectations
- **Rate Limits**: All active rate-limit sliding windows will reset to 0. Watch out for potential momentary brute-force vulnerability.
- **Search Cache**: The cache will start cold. Meilisearch will experience a spike in traffic for ~10 minutes as the cache warms up.
