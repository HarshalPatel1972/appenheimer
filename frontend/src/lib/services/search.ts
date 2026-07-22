import { searchStore } from '$lib/state/search.svelte';

let timeoutId: ReturnType<typeof setTimeout>;

export function executeSearch() {
	if (timeoutId) clearTimeout(timeoutId);
	
	searchStore.status = 'loading';
	
	timeoutId = setTimeout(async () => {
		try {
			const res = await fetch('/api/v1/search', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({
					query: searchStore.query,
					filters: searchStore.filters
				})
			});
			
			if (res.ok) {
				const data = await res.json();
				searchStore.results = data.data || [];
				searchStore.availableFilters = data.availableFilters || null;
				searchStore.status = 'success';
			} else {
				searchStore.status = 'error';
				searchStore.error = 'Search failed';
			}
		} catch (err) {
			searchStore.status = 'error';
			searchStore.error = 'Network error';
		}
	}, 150);
}
