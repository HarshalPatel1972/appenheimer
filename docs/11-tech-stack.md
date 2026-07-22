# Tech Stack

> Every technology in Appenheimer is chosen for long-term maintainability, simplicity, and reliability.

---

# Purpose

This document records the official technologies used by Appenheimer and the reasoning behind each decision.

Technology choices may evolve over time, but changes should be intentional and documented.

---

# Design Philosophy

When choosing technology, prioritize:

- Simplicity
- Maintainability
- Reliability
- Performance
- Long-term stability

Popularity alone is never a reason to adopt a technology.

---

# Frontend

Framework

- SvelteKit

Language

- TypeScript

Build Tool

- Vite

Reason

- Small runtime
- Excellent performance
- Great developer experience
- SSR support
- Strong animation capabilities

---

# Backend

Language

- Go

Reason

- Simple
- Fast
- Easy to maintain
- Excellent concurrency
- Single binary deployment

---

# Database

- PostgreSQL

Reason

- Reliable
- Mature
- Excellent indexing
- Strong relational model

---

# Search Engine

- Meilisearch

Reason

- Deterministic search
- Fast indexing
- Typo tolerance
- Synonyms
- Lightweight operations

---

# Database Access

- pgx
- sqlc

Reason

- Full SQL control
- Compile-time safety
- No ORM magic

---

# Cache

- Redis

Reason

- Sessions
- Caching
- Background jobs

---

# Background Jobs

- Asynq

Reason

- Mature
- Go-native
- Redis-backed

---

# Object Storage

- Cloudflare R2

Reason

- S3 compatible
- No egress fees
- Easy integration

---

# CDN

- Cloudflare

Reason

- Global CDN
- DNS
- Security
- Performance

---

# Monitoring

- Sentry
- Prometheus
- Grafana

Reason

- Error reporting
- Metrics
- Observability

---

# Testing

Frontend

- Vitest
- Playwright

Backend

- Go testing package

Reason

- Fast
- Reliable
- Production-ready

---

# Infrastructure

- Docker
- Docker Compose

Reason

- Reproducible environments
- Simple deployments
- Easy onboarding

---

# Future Changes

Technology may change as the project evolves.

However, any replacement should provide clear long-term benefits and be recorded in the Decision Log.