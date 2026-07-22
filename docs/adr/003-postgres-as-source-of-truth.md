# ADR 003: PostgreSQL as Source of Truth

## Context
We needed a persistent storage layer to replace the initial in-memory MVP dataset.

## Decision
We selected a highly normalized PostgreSQL schema as the absolute source of truth.

## Rationale
- **Data Integrity**: Foreign keys, constraints, and transactions guarantee state consistency.
- **Query Complexity**: It handles our complex taxonomy mapping (Platforms, Pricing, Categories) effortlessly via junction tables.
- **Separation of Concerns**: PostgreSQL will serve as the persistent backend, acting as the ingestion source for the future Meilisearch indexing layer which will exclusively handle text search.
