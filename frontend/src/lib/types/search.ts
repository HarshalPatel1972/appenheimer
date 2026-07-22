export interface AppResult {
	id: string;
	name: string;
	description?: string;
	icon?: string;
	categories?: string[];
	status?: 'Operational' | 'Degraded' | 'Maintenance' | 'Outage' | 'Unknown';
}

export interface AppDetails extends AppResult {
	developer: string;
	websiteUrl: string;
	pricing: string[];
	platforms: string[];
	screenshots: string[];
}

export interface SearchFilters {
	platforms?: string[];
	pricing?: string[];
	categories?: string[];
}
export interface AvailableFilters {
	platforms: string[];
	pricing: string[];
	categories: string[];
}
export interface SearchResponse {
	data: AppResult[];
	availableFilters: AvailableFilters;
}
