import type { AppDetails } from '../types/search';

export class DetailsState {
	activeAppId = $state<string | null>(null);
	appDetails = $state<AppDetails | null>(null);
	isLoading = $state<boolean>(false);
}

export const detailsStore = new DetailsState();
