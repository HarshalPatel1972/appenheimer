<script lang="ts">
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { searchStore } from '$lib/state/search.svelte';
	import { detailsStore } from '$lib/state/details.svelte';
	import { fetchAppDetails } from '$lib/services/apps';
	import { liveSearch, executeSearch } from '$lib/services/search';
	import SearchCanvas from '$lib/search/SearchCanvas.svelte';
	import EmptyState from '$lib/components/EmptyState.svelte';
	import TypingState from '$lib/components/TypingState.svelte';
	import LoadingState from '$lib/components/LoadingState.svelte';
	import ErrorState from '$lib/components/ErrorState.svelte';
	import FilterBar from '$lib/search/FilterBar.svelte';
	import AppDetailsView from '$lib/search/AppDetailsView.svelte';
	import { getIconUrl } from '$lib/utils/icons';
	import { fade } from 'svelte/transition';

	let inputQuery = $state('');
	
	// Sync URL -> State
	// Only overwrites inputQuery when this is a real navigation event (back/forward,
	// direct URL load, filter click) — NOT during live typing. We detect this by
	// checking if the URL's q differs from BOTH searchStore.query (the live-typed value)
	// and inputQuery. If they differ it means an external navigation happened.
	$effect(() => {
		const q = $page.url.searchParams.get('q') || '';
		const appId = $page.url.searchParams.get('app');
		const platform = $page.url.searchParams.get('platform');
		const pricing = $page.url.searchParams.get('pricing');
		const category = $page.url.searchParams.get('category');

		const platforms = platform ? platform.split(',') : [];
		const pricings = pricing ? pricing.split(',') : [];
		const categories = category ? category.split(',') : [];

		let changed = false;

		// Only treat as "changed" if the URL diverges from both the live typed value
		// AND the committed store value — i.e. this is external navigation, not typing.
		const isExternalNav = q !== searchStore.query && q !== inputQuery;
		if (isExternalNav || (q !== searchStore.query && q === '')) {
			inputQuery = q;
			searchStore.query = q;
			changed = true;
		}

		if (JSON.stringify(searchStore.filters.platforms) !== JSON.stringify(platforms)) {
			searchStore.filters.platforms = platforms;
			changed = true;
		}
		if (JSON.stringify(searchStore.filters.pricing) !== JSON.stringify(pricings)) {
			searchStore.filters.pricing = pricings;
			changed = true;
		}
		if (JSON.stringify(searchStore.filters.categories) !== JSON.stringify(categories)) {
			searchStore.filters.categories = categories;
			changed = true;
		}

		if (!q && platforms.length === 0 && pricings.length === 0 && categories.length === 0) {
			searchStore.status = 'idle';
			searchStore.results = [];
		}

		if (changed) {
			const newOrder: {type: string, value: string}[] = [];
			const currentMap = new Map(searchStore.activeFiltersOrder.map(f => [f.type + ':' + f.value, f]));
			
			const addFilters = (type: string, vals: string[]) => {
				for (const v of vals) {
					const key = type + ':' + v;
					if (currentMap.has(key)) {
						newOrder.push(currentMap.get(key)!);
					} else {
						newOrder.push({ type, value: v });
					}
				}
			};
			addFilters('platform', platforms);
			addFilters('pricing', pricings);
			addFilters('category', categories);
			
			searchStore.activeFiltersOrder = newOrder;
			
			if (q.trim().length > 0 || newOrder.length > 0) {
				executeSearch();
			} else {
				searchStore.status = 'idle';
				searchStore.results = [];
			}
		}

		if (appId !== detailsStore.activeAppId) {
			detailsStore.activeAppId = appId;
			if (appId) {
				fetchAppDetails(appId);
			} else {
				detailsStore.appDetails = null;
			}
		}
	});

	// Auto-close if results change and app is not there
	$effect(() => {
		if (detailsStore.activeAppId && searchStore.status === 'success') {
			const found = searchStore.results.find(r => r.id === detailsStore.activeAppId);
			if (!found) {
				closeDrawer();
			}
		}
	});

	let activeApp = $derived(
		detailsStore.activeAppId 
			? searchStore.results.find(r => r.id === detailsStore.activeAppId) || null 
			: null
	);

	function handleInput(e: Event) {
		const target = e.target as HTMLInputElement;
		inputQuery = target.value;
		// Live search: syncs query + debounces API call on every keystroke
		liveSearch(inputQuery);
	}

	function commitSearch() {
		// Flush debounce immediately — search is already in searchStore.query
		executeSearch();
		// Update URL to reflect committed query
		const url = new URL($page.url);
		if (searchStore.query) {
			url.searchParams.set('q', searchStore.query);
		} else {
			url.searchParams.delete('q');
		}
		goto(url, { keepFocus: true, noScroll: true, replaceState: true });
	}

	function handleInputKeydown(e: KeyboardEvent) {
		if (e.key === 'Enter') {
			e.preventDefault();
			commitSearch();
		}
	}

	function closeDrawer() {
		const q = $page.url.searchParams.get('q') || '';
		if (q) {
			goto(`/?q=${q}`, { keepFocus: true, noScroll: true });
		} else {
			goto(`/`, { keepFocus: true, noScroll: true });
		}
	}

	function handleKeydown(e: KeyboardEvent) {
		if (e.key === 'Escape' && detailsStore.activeAppId) {
			closeDrawer();
			return;
		}

		if (detailsStore.activeAppId && searchStore.status === 'success') {
			const currentIndex = searchStore.results.findIndex(r => r.id === detailsStore.activeAppId);
			if (currentIndex === -1) return;

			if (e.key === 'ArrowRight' || e.key === 'ArrowDown') {
				e.preventDefault();
				if (currentIndex < searchStore.results.length - 1) {
					const nextId = searchStore.results[currentIndex + 1].id;
					const q = $page.url.searchParams.get('q') || '';
					goto(`/?q=${q}&app=${nextId}`, { keepFocus: true, noScroll: true, replaceState: true });
				}
			} else if (e.key === 'ArrowLeft' || e.key === 'ArrowUp') {
				e.preventDefault();
				if (currentIndex > 0) {
					const prevId = searchStore.results[currentIndex - 1].id;
					const q = $page.url.searchParams.get('q') || '';
					goto(`/?q=${q}&app=${prevId}`, { keepFocus: true, noScroll: true, replaceState: true });
				}
			}
		}
	}
</script>

<svelte:window onkeydown={handleKeydown} />

<main class="app-main">
	<!-- CENTER STAGE (App Badge + Search Bar) -->
	<div class="center-stage" class:is-active={!!detailsStore.activeAppId}>
		{#if detailsStore.activeAppId && activeApp}
			<div class="active-app-badge" transition:fade={{ duration: 200 }}>
				<div class="active-icon-box">
					<img src={getIconUrl(activeApp.icon, activeApp.name)} alt={activeApp.name} class="active-icon-img" />
				</div>
				<span class="active-app-title">{activeApp.name}</span>
				<button class="close-btn" onclick={closeDrawer} aria-label="Close app details">✕</button>
			</div>
		{/if}

		<div class="search-bar-wrapper">
			<input 
				type="text" 
				class="search-input" 
				class:compact={!!detailsStore.activeAppId}
				placeholder="Search for software..." 
				value={inputQuery}
				oninput={handleInput}
				onkeydown={handleInputKeydown}
			/>

			{#if !detailsStore.activeAppId}
				<FilterBar />
			{/if}
		</div>
	</div>
	
	<!-- CANVAS & DRIFTING LOGOS -->
	<div class="content">
		{#if searchStore.status === 'idle'}
			<EmptyState />
		{:else if searchStore.status === 'typing'}
			<TypingState />
		{:else if searchStore.status === 'loading'}
			<LoadingState />
		{:else if searchStore.status === 'error'}
			<ErrorState message={searchStore.error} />
		{:else if searchStore.status === 'success' && searchStore.results.length === 0}
			<EmptyState />
		{:else if searchStore.status === 'success'}
			<SearchCanvas />
		{/if}
	</div>

	<!-- 4-CORNER SCATTERED DETAILS VIEW -->
	{#if detailsStore.activeAppId && activeApp}
		<AppDetailsView app={activeApp} details={detailsStore.appDetails} />
	{/if}
</main>

<style>
	.app-main {
		display: flex;
		flex-direction: column;
		height: 100vh;
		width: 100vw;
		background: var(--bg-canvas);
		overflow: hidden;
		position: relative;
	}

	.center-stage {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		z-index: 200;
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 16px;
		width: 100%;
		max-width: 640px;
		padding: 0 16px;
		box-sizing: border-box;
		transition: all 0.5s cubic-bezier(0.22, 1, 0.36, 1);
	}

	.center-stage.is-active {
		max-width: 720px;
	}

	.active-app-badge {
		display: flex;
		align-items: center;
		gap: 12px;
		background: var(--bg-surface);
		border: 3px solid var(--border-subtle);
		box-shadow: 6px 6px 0 rgba(0,0,0,1);
		padding: 8px 16px;
		flex-shrink: 0;
	}

	.active-icon-box {
		width: 36px;
		height: 36px;
		overflow: hidden;
		display: flex;
		align-items: center;
		justify-content: center;
		border: 1px solid var(--border-subtle);
		background: #ffffff;
	}

	.active-icon-img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.active-app-title {
		font-size: 1.1rem;
		font-weight: 800;
		font-family: var(--font-mono);
		color: var(--text-main);
		text-transform: uppercase;
		letter-spacing: 0.05em;
	}

	.close-btn {
		background: transparent;
		border: 1.5px solid var(--border-subtle);
		color: var(--text-main);
		font-weight: 800;
		font-size: 0.85rem;
		cursor: pointer;
		padding: 2px 8px;
		margin-left: 4px;
		transition: all 0.2s ease;
	}

	.close-btn:hover {
		background: var(--color-primary);
		color: #ffffff;
		border-color: var(--color-primary);
	}

	.search-bar-wrapper {
		display: flex;
		flex-direction: column;
		align-items: center;
		width: 100%;
		transition: all 0.5s cubic-bezier(0.22, 1, 0.36, 1);
	}

	.search-input {
		width: 100%;
		padding: 16px 24px;
		border: 3px solid var(--border-subtle);
		background: var(--bg-surface);
		color: var(--text-main);
		font-size: 1.2rem;
		font-weight: 700;
		box-shadow: 6px 6px 0 rgba(0,0,0,1);
		transition: all 0.3s ease;
		outline: none;
		border-radius: 0;
		box-sizing: border-box;
	}

	.search-input.compact {
		padding: 10px 16px;
		font-size: 1rem;
		max-width: 320px;
	}

	.search-input:focus {
		border-color: var(--color-primary);
		box-shadow: 6px 6px 0 var(--color-primary);
		transform: translate(-2px, -2px);
	}

	.search-input.compact:focus {
		transform: none;
	}

	.content {
		flex: 1;
		width: 100%;
		height: 100%;
		position: relative;
	}
</style>
