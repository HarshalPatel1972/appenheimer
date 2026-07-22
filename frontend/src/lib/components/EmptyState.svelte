<script lang="ts">
	import { searchStore } from '$lib/state/search.svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	let activeFilterCount = $derived(searchStore.activeFiltersOrder.length);

	function clearFilters() {
		const url = new URL($page.url);
		url.searchParams.delete('platform');
		url.searchParams.delete('pricing');
		url.searchParams.delete('category');
		goto(url, { keepFocus: true, noScroll: true });
	}
</script>

<div class="empty-state">
	{#if activeFilterCount > 0 && searchStore.results.length === 0 && searchStore.status === 'success'}
		<h3>No apps match these filters.</h3>
		<p>Try removing some filters to see more results.</p>
		<button class="clear-btn" onclick={clearFilters}>Clear Filters</button>
	{:else}
		<div class="icon"></div>
		<h2>Discover the best software</h2>
		<p>Search for design tools, code editors, and more.</p>
	{/if}
</div>

<style>
	.empty-state {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		height: 100%;
		color: var(--text-muted);
		text-align: center;
	}

	.icon {
		width: 64px;
		height: 64px;
		background: var(--bg-surface);
		border-radius: 16px;
		margin-bottom: 24px;
		border: 1px solid var(--border-subtle);
	}

	h2, h3 {
		margin: 0 0 8px 0;
		color: var(--text-main);
	}

	p {
		margin: 0;
		font-size: 1rem;
		max-width: 300px;
		line-height: 1.5;
	}

	.clear-btn {
		margin-top: 24px;
		padding: 10px 20px;
		background: var(--bg-surface);
		border: 1px solid var(--border-subtle);
		color: var(--text-main);
		border-radius: 8px;
		cursor: pointer;
		font-weight: 500;
		transition: all 0.2s ease;
	}

	.clear-btn:hover {
		background: rgba(255,255,255,0.1);
	}
</style>
