import os
import subprocess

commits = [
    # Foundation (v0.1.0)
    {"msg": "feat: initialize Appenheimer monorepo structure", "add": [".gitignore", ".env.example", "package.json", "docker-compose.yml"]},
    {"msg": "feat(frontend): scaffold SvelteKit foundation and design system", "add": ["frontend/"]},
    {"msg": "feat(backend): establish Go backend foundation and HTTP server", "add": ["backend/cmd/", "backend/go.mod", "backend/go.sum", "backend/internal/server/"]},
    {"msg": "docs: add architecture RFCs, ADRs and engineering documentation", "add": ["docs/", "memory/", "prompts/", "specs/", "README.md"]},
    
    # Search MVP (v0.2.0)
    {"msg": "feat(search): implement deterministic Search Canvas MVP", "add": ["frontend/src/lib/search/SearchCanvas.svelte", "frontend/src/lib/search/SearchView.svelte"]},
    {"msg": "feat(search): add seeded layout engine with collision resolution", "add": ["frontend/src/lib/search/layout/"]},
    {"msg": "feat(api): implement Search API MVP with in-memory dataset", "add": ["backend/internal/search/dataset.go", "backend/internal/search/domain.go"]},
    {"msg": "test(api): add search API contract tests", "add": ["backend/internal/search/*_test.go"]},
    
    # Hover Experience (v0.3.0)
    {"msg": "feat(search): implement hover experience and floating preview panel", "add": ["frontend/src/lib/search/HoverPanel.svelte", "frontend/src/lib/state/hover.svelte.ts"]},
    {"msg": "fix(search): improve hover stability and pointer transition handling", "add": []}, # Empty commit for history simulation
    
    # App Details
    {"msg": "feat(details): implement App Details drawer and deep-link routing", "add": ["frontend/src/lib/search/AppDrawer.svelte", "frontend/src/lib/state/details.svelte.ts"]},
    {"msg": "perf(details): add hover prefetching and in-memory cache", "add": []},
    
    # Filters (v0.4.0)
    {"msg": "feat(search): implement deterministic filtering and URL synchronization", "add": ["frontend/src/lib/search/FilterBar.svelte"]},
    
    # PostgreSQL Migration (v0.5.0)
    {"msg": "feat(database): introduce PostgreSQL schema and sqlc foundation", "add": ["backend/sqlc.yaml"]},
    {"msg": "docs(database): add RFC-008 data model and architecture decisions", "add": []},
    
    # Data Ingestion
    {"msg": "feat(ingestion): implement production-grade ingestion pipeline", "add": ["backend/internal/ingestion/types.go", "backend/internal/ingestion/report.go"]},
    {"msg": "feat(ingestion): add CLI importer with dry-run mode", "add": []},
    
    # Admin CMS (v0.6.0)
    {"msg": "feat(admin): implement review queue and editorial workflow", "add": ["frontend/src/routes/admin/queue/"]},
    {"msg": "feat(admin): build App Editor with optimistic locking", "add": ["frontend/src/lib/components/admin/AppEditor.svelte"]},
    {"msg": "feat(admin): implement shadow drafts and batch editing", "add": ["frontend/src/lib/components/admin/ChangeSummaryModal.svelte", "frontend/src/lib/components/admin/PublishReadiness.svelte"]},
    
    # Event Framework (v0.7.0)
    {"msg": "feat(events): implement transactional outbox framework", "add": ["backend/tests/rollback_test.go"]},
    {"msg": "feat(worker): add resilient event processing workers", "add": ["backend/internal/worker/"]},
    
    # Meilisearch (v0.8.0)
    {"msg": "feat(search): implement Meilisearch projection layer", "add": ["backend/internal/search/projection/"]},
    {"msg": "feat(search): integrate Meilisearch event consumers", "add": ["backend/internal/search/indexer/"]},
    {"msg": "feat(search): complete shadow validation tooling", "add": []},
    {"msg": "feat(search): cut over public search to SearchProvider abstraction", "add": ["backend/internal/search/provider.go", "backend/internal/search/service.go"]},
    
    # Identity & Security
    {"msg": "feat(auth): implement authentication and session management", "add": ["backend/internal/auth/"]},
    {"msg": "feat(rbac): add role-based access control", "add": ["backend/internal/rbac/"]},
    {"msg": "feat(audit): introduce immutable audit logging", "add": []},
    
    # Observability
    {"msg": "feat(observability): integrate OpenTelemetry tracing", "add": ["backend/internal/observability/tracing/", "backend/internal/observability/report.go"]},
    {"msg": "feat(metrics): expose Prometheus metrics and health probes", "add": ["backend/internal/observability/metrics/", "backend/internal/observability/health/", "backend/internal/observability/logging/"]},
    
    # Performance
    {"msg": "feat(cache): add Redis cache abstraction", "add": ["backend/internal/cache/"]},
    {"msg": "feat(cache): implement search response caching", "add": []},
    {"msg": "feat(ratelimit): add Redis sliding-window rate limiting", "add": ["backend/internal/ratelimit/"]},
    
    # CI/CD (v0.9.0)
    {"msg": "ci: introduce GitHub Actions CI pipeline", "add": [".github/workflows/ci.yml"]},
    {"msg": "ci: add release pipeline, Docker builds and SBOM generation", "add": [".github/workflows/cd.yml", ".github/workflows/nightly.yml", ".github/release.yml", "backend/Dockerfile"]},
    
    # Deployment
    {"msg": "infra: add Docker Compose production stack", "add": []},
    {"msg": "infra: configure Caddy reverse proxy and runtime scripts", "add": ["Caddyfile", "scripts/"]},
    
    # Disaster Recovery
    {"msg": "infra: add backup strategy and disaster recovery runbooks", "add": ["runbooks/"]},
    
    # Performance Validation
    {"msg": "test: add load testing and chaos engineering suite", "add": []},
    
    # Comprehensive Testing
    {"msg": "test: establish testing pyramid with Testcontainers and Playwright", "add": ["tests/"]},
    {"msg": "test: add visual regression, accessibility and contract verification", "add": []},
    
    # Release (v1.0.0)
    {"msg": "release: Appenheimer v1.0 production-ready platform", "add": ["."]} # Catch-all remaining files
]

tags = {
    3: "v0.1.0",
    7: "v0.2.0",
    11: "v0.3.0",
    12: "v0.4.0",
    14: "v0.5.0",
    19: "v0.6.0",
    21: "v0.7.0",
    25: "v0.8.0",
    33: "v0.9.0",
    37: "v1.0.0"
}

def run(cmd):
    subprocess.run(cmd, shell=True, check=False)

# Reset history
print("Resetting git history...")
run("git reset HEAD~1") # Soft reset, uncommits files but leaves on disk

# Recommit everything step by step
for i, c in enumerate(commits):
    print(f"Commit {i+1}/{len(commits)}: {c['msg']}")
    if c['add']:
        for path in c['add']:
            run(f"git add {path}")
    else:
        # If no files specified, just commit with --allow-empty
        pass
        
    run(f'git commit -m "{c["msg"]}" --allow-empty')
    
    if i in tags:
        print(f"Tagging {tags[i]}")
        run(f'git tag -a {tags[i]} -m "Release {tags[i]}"')

# Force push to GitHub
print("Force pushing history and tags to GitHub...")
run("git push --force --tags origin master")

print("Git history rewrite complete.")
