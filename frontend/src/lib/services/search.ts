import { searchStore } from '$lib/state/search.svelte';
import { http } from '$lib/api/http';

let timeoutId: ReturnType<typeof setTimeout>;

async function runSearch(query: string) {
	try {
		const data = await http('/api/v1/search', {
			method: 'POST',
			body: JSON.stringify({
				query,
				filters: searchStore.filters
			})
		});
		// Only apply results if query hasn't changed underneath us
		if (searchStore.query === query) {
			searchStore.results = data.data || [];
			searchStore.availableFilters = data.availableFilters || null;
			searchStore.status = 'success';
		}
	} catch (err) {
		if (searchStore.query === query) {
			searchStore.status = 'error';
			searchStore.error = 'Search failed';
		}
	}
}

/**
 * Live search — debounced 300ms, fires on every keystroke.
 * Syncs inputQuery into searchStore.query immediately so URL and
 * state never diverge, then schedules the API call.
 */
export function liveSearch(inputQuery: string) {
	if (timeoutId) clearTimeout(timeoutId);

	const trimmed = inputQuery.trim();
	searchStore.query = trimmed;

	if (trimmed.length === 0 && searchStore.activeFiltersOrder.length === 0) {
		searchStore.status = 'idle';
		searchStore.results = [];
		return;
	}

	searchStore.status = 'loading';

	timeoutId = setTimeout(() => runSearch(trimmed), 300);
}

/**
 * Explicit search — fires immediately (used on Enter / suggestion click).
 * Still uses the current searchStore.query.
 */
export function executeSearch() {
	if (timeoutId) clearTimeout(timeoutId);

	const trimmed = searchStore.query.trim();

	if (trimmed.length === 0 && searchStore.activeFiltersOrder.length === 0) {
		searchStore.status = 'idle';
		searchStore.results = [];
		return;
	}

	searchStore.status = 'loading';
	runSearch(trimmed);
}
