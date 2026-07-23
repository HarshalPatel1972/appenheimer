import { detailsStore } from '$lib/state/details.svelte';
import type { AppDetails } from '$lib/types/search';
import { http } from '$lib/api/http';

const CACHE_TTL_MS = 10 * 60 * 1000; // 10 minutes
const cache = new Map<string, { data: AppDetails, expiresAt: number }>();
const fetching = new Set<string>();

export async function fetchAppDetails(id: string) {
	detailsStore.isLoading = true;
	
	const now = Date.now();
	if (cache.has(id)) {
		const entry = cache.get(id)!;
		if (now < entry.expiresAt) {
			detailsStore.appDetails = entry.data;
			detailsStore.isLoading = false;
			return;
		} else {
			cache.delete(id);
		}
	}

	if (fetching.has(id)) {
		while (fetching.has(id)) {
			await new Promise(r => setTimeout(r, 50));
		}
		if (cache.has(id)) {
			const entry = cache.get(id)!;
			if (now < entry.expiresAt) {
				detailsStore.appDetails = entry.data;
				detailsStore.isLoading = false;
				return;
			}
		}
	}

	fetching.add(id);
	try {
		const data = await http(`/api/v1/apps/${id}`);
		cache.set(id, { data, expiresAt: Date.now() + CACHE_TTL_MS });
		if (detailsStore.activeAppId === id) {
			detailsStore.appDetails = data;
		}
	} catch (err) {
		console.error("Failed to fetch app details", err);
		if (detailsStore.activeAppId === id) {
			detailsStore.appDetails = null;
		}
	} finally {
		fetching.delete(id);
		if (detailsStore.activeAppId === id) {
			detailsStore.isLoading = false;
		}
	}
}

export async function prefetchAppDetails(id: string) {
	const now = Date.now();
	if (cache.has(id)) {
		if (now < cache.get(id)!.expiresAt) return;
		cache.delete(id);
	}
	if (fetching.has(id)) return;
	
	fetching.add(id);
	try {
		const data = await http(`/api/v1/apps/${id}`);
		cache.set(id, { data, expiresAt: Date.now() + CACHE_TTL_MS });
	} catch (err) {
		// ignore
	} finally {
		fetching.delete(id);
	}
}
