import type { AppResult, SearchFilters, AvailableFilters } from '../types/search';

export class SearchState {
	query = $state('');
	filters = $state<SearchFilters>({});
	activeFiltersOrder = $state<{ type: string, value: string }[]>([]);
	
	results = $state<AppResult[]>([]);
	availableFilters = $state<AvailableFilters | null>(null);
	status = $state<'idle' | 'typing' | 'loading' | 'success' | 'error'>('idle');
	error = $state('');

	setTyping() {
		this.status = 'typing';
	}
}

export const searchStore = new SearchState();

// Simple session sync
if (typeof sessionStorage !== 'undefined') {
	const saved = sessionStorage.getItem('appenheimer_search');
	if (saved) {
		try {
			const parsed = JSON.parse(saved);
			if (parsed.query) searchStore.query = parsed.query;
			if (parsed.filters) searchStore.filters = parsed.filters;
		} catch (e) {}
	}
	
	$effect.root(() => {
		$effect(() => {
			const state = { query: searchStore.query, filters: searchStore.filters };
			sessionStorage.setItem('appenheimer_search', JSON.stringify(state));
		});
	});
}
