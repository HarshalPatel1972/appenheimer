import { searchStore } from '$lib/state/search.svelte';
import { http } from '$lib/api/http';

let timeoutId: ReturnType<typeof setTimeout>;

export function executeSearch() {
	if (timeoutId) clearTimeout(timeoutId);
	
	searchStore.status = 'loading';
	
	timeoutId = setTimeout(async () => {
		try {
			const data = await http('/api/v1/search', {
				method: 'POST',
				body: JSON.stringify({
					query: searchStore.query,
					filters: searchStore.filters
				})
			});
			
			searchStore.results = data.data || [];
			searchStore.availableFilters = data.availableFilters || null;
			searchStore.status = 'success';
		} catch (err) {
			searchStore.status = 'error';
			searchStore.error = 'Search failed';
		}
	}, 150);
}
