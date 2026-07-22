# Database

> The database is the single source of truth for Appenheimer.

---

# Purpose

This document defines how data is stored, owned, and maintained.

The database prioritizes correctness, consistency, and long-term maintainability over convenience.

---

# Core Principles

- Design the data first.
- Normalize where appropriate.
- Duplicate only when justified.
- Every entity has one owner.
- Data integrity is more important than convenience.

---

# Source of Truth

PostgreSQL is the only source of truth.

Everything else is derived from it.

Examples:

- Meilisearch indexes
- Redis cache
- Analytics
- Background jobs

None of these own application data.

---

# Primary Entities

The database will store entities such as:

- Applications
- Categories
- Tags
- Features
- Platforms
- Developers
- Collections
- Users
- Updates
- Aliases
- Synonyms

Each entity should have a clear responsibility.

---

# Relationships

Relationships should be explicit.

Prefer relational modeling over duplicated data.

Many-to-many relationships should use junction tables.

---

# Data Ownership

Each piece of information should exist only once.

For example:

- Application name belongs to the Application entity.
- Developer information belongs to the Developer entity.
- Categories belong to Category.

Other tables reference these entities instead of duplicating them.

---

# Search Data

The database stores canonical information.

Search indexes are generated from database records.

Search-specific optimizations must never modify the canonical data.

---

# Migrations

Every schema change must use versioned migrations.

Schema changes should be:

- repeatable
- reversible whenever possible
- reviewed before execution

---

# Constraints

Use database constraints wherever possible.

Examples include:

- Foreign keys
- Unique constraints
- Check constraints
- Not null constraints

Application code should not replace database integrity.

---

# Indexing

Indexes should exist to support measured query patterns.

Indexes should not be added preemptively.

Review indexes regularly as the application evolves.

---

# Soft Deletes

Use soft deletes only when recovery or history is valuable.

Otherwise prefer permanent deletion.

Every soft delete must have a clear justification.

---

# Auditability

Important administrative actions should be traceable.

The system should preserve sufficient history to understand significant changes.

---

# Backups

Backups are mandatory.

Recovery procedures should be documented and tested.

A backup that has never been restored is not considered verified.

---

# Performance

Optimize queries only after measuring.

Avoid premature denormalization.

Favor correctness first, optimization second.

---

# Future Evolution

The schema should evolve incrementally.

Avoid breaking changes whenever possible.

Migration should always be easier than rebuilding.

---

# Final Principle

> Data is the foundation of Appenheimer. Every architectural decision should protect its integrity.