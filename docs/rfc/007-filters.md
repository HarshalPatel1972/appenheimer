# RFC 007 - Filters

## 1. Context
With the core discovery loop complete, the final piece of the MVP exploration experience is **Filters**. Users need the ability to narrow down the deterministic layout by specific criteria to find the exact software they need.

## 2. Goals
- Provide rapid, client-driven refinement of search results.
- Keep the UI minimal and unobtrusive, preserving the "no unnecessary reading" philosophy.
- Sync filter state perfectly with the URL for shareability.
- Filter operations must remain deterministic and fast.

## 3. Approved Decisions
- **Filter Dimensions**: `Platform`, `Pricing`, and `Category` ONLY for MVP.
- **Filter Logic**: `AND` between groups, `OR` within groups (e.g., `(Windows OR Mac) AND Free`).
- **URL Strategy**: Comma-separated lists (`/?q=design&platform=windows,mac&pricing=free`).
- **Visual Feedback**: The UI must explicitly show the number of active filters (e.g., "Filters (3)").
- **Dead-end Prevention**: If filters result in 0 apps, present a "Clear Filters" button instead of a generic empty state.
- **Query Preservation**: Changing filters does not alter the underlying text search query.

## 4. Technical Implications
### Backend Changes
- **Data Model Pivot**: `Pricing` and `Platforms` must be lifted from the `AppDetails` struct into the base `App` struct in `domain.go` so the search engine can filter them efficiently in-memory before sorting.
- **API Payload**: `POST /api/v1/search` must parse the optional `Filters` object and apply intersection checks during the search loop.

### Frontend Changes
- **Filter State**: `SearchState` stores the active filters and pushes them into the `POST` payload.
- **URL Synchronization**: `SearchPage.svelte` translates URL parameters into state.
- **New Components**: A `FilterBar.svelte` component. An updated `EmptyState` component handling the 0-results-due-to-filters scenario.
