import type { AppDetails } from '../types/search';

export class DetailsState {
	activeAppId = $state<string | null>(null);
	appDetails = $state<AppDetails | null>(null);
	isLoading = $state<boolean>(false);
}

export const detailsStore = new DetailsState();

if (typeof window !== 'undefined' && typeof sessionStorage !== 'undefined') {
	const urlParams = new URLSearchParams(window.location.search);
	const urlAppId = urlParams.get('app');
	
	if (urlAppId) {
		detailsStore.activeAppId = urlAppId;
	} else if (!window.location.search) {
		const savedAppId = sessionStorage.getItem('appenheimer_active_app');
		if (savedAppId) {
			detailsStore.activeAppId = savedAppId;
		}
	}
	
	$effect.root(() => {
		$effect(() => {
			if (detailsStore.activeAppId) {
				sessionStorage.setItem('appenheimer_active_app', detailsStore.activeAppId);
			} else {
				sessionStorage.removeItem('appenheimer_active_app');
			}
		});
	});
}
