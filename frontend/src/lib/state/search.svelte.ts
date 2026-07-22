import type { AppResult, SearchFilters, AvailableFilters } from '../types/search';

export class SearchState {
	query = $state('');
	filters = $state<SearchFilters>({});
	activeFiltersOrder = $state<{ type: string, value: string }[]>([]);
	
	results = $state<AppResult[]>([]);
	availableFilters = $state<AvailableFilters | null>(null);
	status = $state<'idle' | 'typing' | 'loading' | 'success' | 'error'>('idle');
	error = $state('');
}

export const searchStore = new SearchState();
