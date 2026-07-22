# ADR 001: SQLC over GORM

## Context
When building the PostgreSQL database layer for Appenheimer, we needed to select a database access tool for Go.

## Decision
We chose `sqlc` over `GORM`.

## Rationale
- **Type Safety**: `sqlc` generates type-safe Go structs and interfaces directly from raw SQL files.
- **Performance**: No reflection overhead at runtime.
- **Transparency**: Developers retain full control over the exact SQL being executed, avoiding N+1 hidden query issues common in ORMs.
- **Maintainability**: SQL is the universal language of relational data; tying our data access to raw SQL prevents vendor/ORM lock-in.
