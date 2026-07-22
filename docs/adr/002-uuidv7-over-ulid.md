# ADR 002: UUIDv7 over ULID

## Context
Appenheimer required a globally unique identifier for primary keys that also provided chronological sorting to prevent B-tree fragmentation in PostgreSQL.

## Decision
We chose UUIDv7 over ULID.

## Rationale
- **Standards**: UUIDv7 is an official standard, whereas ULID, while popular, is not universally standardized.
- **PostgreSQL Support**: UUIDv7 fits natively into Postgres' `UUID` column type without requiring string-to-bytes conversions.
- **Industry Momentum**: The industry is largely converging on UUIDv7 for time-sorted unique IDs.
