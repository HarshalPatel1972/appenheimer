<script lang="ts">
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { searchStore } from '$lib/state/search.svelte';
	import { detailsStore } from '$lib/state/details.svelte';
	import { fetchAppDetails } from '$lib/services/apps';
	import { executeSearch } from '$lib/services/search';
	import SearchCanvas from '$lib/search/SearchCanvas.svelte';
	import EmptyState from '$lib/components/EmptyState.svelte';
	import TypingState from '$lib/components/TypingState.svelte';
	import LoadingState from '$lib/components/LoadingState.svelte';
	import ErrorState from '$lib/components/ErrorState.svelte';
	import FilterBar from '$lib/search/FilterBar.svelte';
	import { fade } from 'svelte/transition';

	let inputQuery = $state('');
	
	// Sync URL -> State
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
		if (q !== searchStore.query) {
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
			executeSearch();
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

	let queryChips = $derived(searchStore.query.trim() ? searchStore.query.trim().split(/\s+/) : []);

	function handleInput(e: Event) {
		const target = e.target as HTMLInputElement;
		inputQuery = target.value;
		if (inputQuery.trim().length > 0) {
			searchStore.setTyping();
		} else if (queryChips.length === 0) {
			searchStore.status = 'idle';
		}
	}

	function commitSearch() {
		if (inputQuery.trim()) {
			const current = searchStore.query ? searchStore.query.trim() + ' ' : '';
			searchStore.query = current + inputQuery.trim();
			inputQuery = '';
		}
		
		executeSearch();
		
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
		} else if (e.key === 'Backspace' && inputQuery === '' && queryChips.length > 0) {
			removeChip(queryChips.length - 1);
		}
	}

	function removeChip(index: number) {
		const newChips = [...queryChips];
		newChips.splice(index, 1);
		searchStore.query = newChips.join(' ');
		
		executeSearch();
		
		const url = new URL($page.url);
		if (searchStore.query) {
			url.searchParams.set('q', searchStore.query);
		} else {
			url.searchParams.delete('q');
		}
		goto(url, { keepFocus: true, noScroll: true, replaceState: true });
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
	<div class="search-bar-container">
		<input 
			type="text" 
			class="search-input" 
			placeholder={queryChips.length === 0 ? "Search for software..." : "Add another keyword..."} 
			value={inputQuery}
			oninput={handleInput}
			onkeydown={handleInputKeydown}
		/>
		
		{#if queryChips.length > 0}
			<div class="query-chips">
				{#each queryChips as chip, i}
					<button class="chip" onclick={() => removeChip(i)}>
						{chip}
						<span class="remove">✕</span>
					</button>
				{/each}
			</div>
		{/if}

		<FilterBar />
	</div>
	
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

	{#if detailsStore.activeAppId}
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div class="canvas-dimmer" onclick={closeDrawer} transition:fade={{ duration: 200 }}></div>
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

	.search-bar-container {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		z-index: 100;
		width: 100%;
		max-width: 600px;
		padding: 0 16px;
		box-sizing: border-box;
		transition: top 0.4s ease, transform 0.4s ease;
	}

	.search-input {
		width: 100%;
		padding: 16px 24px;
		border-radius: 99px;
		border: 1px solid var(--border-subtle);
		background: rgba(30,30,30,0.8);
		backdrop-filter: blur(10px);
		color: var(--text-main);
		font-size: 1.1rem;
		box-shadow: 0 8px 32px rgba(0,0,0,0.4);
		transition: border-color 0.2s, box-shadow 0.2s;
		outline: none;
	}

	.search-input:focus {
		border-color: var(--color-primary);
		box-shadow: 0 8px 32px rgba(0,0,0,0.6), 0 0 0 2px rgba(59, 130, 246, 0.3);
	}

	.query-chips {
		display: flex;
		gap: 8px;
		flex-wrap: wrap;
		justify-content: center;
		margin-top: 16px;
	}

	.chip {
		display: inline-flex;
		align-items: center;
		gap: 8px;
		padding: 6px 14px;
		background: rgba(59, 130, 246, 0.15);
		border: 1px solid rgba(59, 130, 246, 0.4);
		border-radius: 99px;
		color: var(--text-main);
		font-size: 0.9rem;
		cursor: pointer;
		transition: all 0.2s ease;
	}

	.chip:hover {
		background: rgba(239, 68, 68, 0.15);
		border-color: rgba(239, 68, 68, 0.4);
	}

	.chip .remove {
		font-size: 0.8rem;
		opacity: 0.7;
	}

	.chip:hover .remove {
		opacity: 1;
		color: #f87171;
	}

	.content {
		flex: 1;
		width: 100%;
		height: 100%;
		position: relative;
	}

	.canvas-dimmer {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background: rgba(0,0,0,0.5);
		backdrop-filter: blur(4px);
		z-index: 500;
	}
</style>
