# Architecture

> Appenheimer is built as a collection of loosely coupled, replaceable components.

---

# Purpose

This document describes the high-level architecture of Appenheimer.

It focuses on responsibilities and data flow, not implementation details.

---

# Architectural Principles

- Keep components independent.
- Keep data ownership clear.
- Prefer composition over coupling.
- Design for maintainability first.
- Every component should have a single responsibility.

---

# System Overview

The system consists of five major parts:

- Frontend
- Backend API
- Database
- Search Engine
- Supporting Services

Each component has a clearly defined responsibility.

---

# Frontend

Responsibilities:

- Render the user interface.
- Handle user interactions.
- Display search results.
- Manage animations.
- Maintain local UI state.

The frontend should never become the source of truth.

---

# Backend

Responsibilities:

- Business logic.
- Authentication.
- Data validation.
- API endpoints.
- Administrative operations.
- Communication with external services.

The backend owns the application logic.

---

# Database

Responsibilities:

- Store all persistent data.
- Maintain relationships.
- Preserve data integrity.
- Serve as the single source of truth.

Everything important ultimately lives here.

---

# Search Engine

Responsibilities:

- Index searchable data.
- Execute deterministic searches.
- Handle aliases.
- Handle synonyms.
- Handle typo tolerance.
- Return ranked results.

The search engine is an index, not the source of truth.

---

# Cache

Responsibilities:

- Improve performance.
- Reduce repeated work.
- Store temporary data.

Cached data must always be reproducible.

---

# Object Storage

Responsibilities:

- Store static assets.
- Store images.
- Store uploaded media.

Object storage is never responsible for application data.

---

# Background Workers

Responsibilities:

- Long-running tasks.
- Search indexing.
- Image processing.
- Notifications.
- Maintenance jobs.

Background work should never block user requests.

---

# Data Flow

User

↓

Frontend

↓

Backend API

↓

Database

↓

Search Index

↓

Frontend

The backend coordinates the system.

The database owns the data.

The search engine accelerates discovery.

---

# Source of Truth

Exactly one component owns each piece of information.

- Database → Persistent data
- Search Engine → Search index
- Cache → Temporary data
- Frontend → UI state

Ownership must never overlap.

---

# Replaceability

Every major component should be replaceable.

Replacing one technology should require minimal changes to the rest of the system.

This principle influences every architectural decision.

---

# Scalability

The architecture should support gradual growth.

Scale should be achieved by improving individual components rather than redesigning the system.

Premature optimization should be avoided.

---

# Final Principle

> The architecture should remain understandable to a new contributor after reading this document.