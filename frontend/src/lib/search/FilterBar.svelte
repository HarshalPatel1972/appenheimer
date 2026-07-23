<script lang="ts">
	import { searchStore } from '$lib/state/search.svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	
	let available = $derived(searchStore.availableFilters);
	let active = $derived(searchStore.activeFiltersOrder);
	let count = $derived(active.length);
	
	function toggleFilter(type: string, value: string) {
		const url = new URL($page.url);
		let current = url.searchParams.get(type) ? url.searchParams.get(type)!.split(',') : [];
		
		if (current.includes(value)) {
			current = current.filter(c => c !== value);
		} else {
			current.push(value);
		}
		
		if (current.length > 0) {
			url.searchParams.set(type, current.join(','));
		} else {
			url.searchParams.delete(type);
		}
		
		goto(url, { keepFocus: true, noScroll: true, replaceState: true });
	}
</script>

<div class="filter-bar">
	<div class="filter-controls">
		{#if available}
			{#if available.platforms && available.platforms.length > 0}
				<div class="filter-group">
					<span class="label">Platform:</span>
					{#each available.platforms as p}
						<button 
							class="filter-pill" 
							class:active={searchStore.filters.platforms?.includes(p)}
							onclick={() => toggleFilter('platform', p)}>
							{p}
						</button>
					{/each}
				</div>
			{/if}
			{#if available.pricing && available.pricing.length > 0}
				<div class="filter-group">
					<span class="label">Pricing:</span>
					{#each available.pricing as p}
						<button 
							class="filter-pill" 
							class:active={searchStore.filters.pricing?.includes(p)}
							onclick={() => toggleFilter('pricing', p)}>
							{p}
						</button>
					{/each}
				</div>
			{/if}
			{#if available.categories && available.categories.length > 0}
				<div class="filter-group">
					<span class="label">Category:</span>
					{#each available.categories as c}
						<button 
							class="filter-pill" 
							class:active={searchStore.filters.categories?.includes(c)}
							onclick={() => toggleFilter('category', c)}>
							{c}
						</button>
					{/each}
				</div>
			{/if}
		{/if}
	</div>
	
	{#if count > 0}
		<div class="active-filters">
			<span class="count">Filters ({count}):</span>
			<div class="active-list">
				{#each active as f, i}
					<span class="active-item">{f.value}</span>
					{#if i < active.length - 1}
						<span class="dot">•</span>
					{/if}
				{/each}
			</div>
		</div>
	{/if}
</div>

<style>
	.filter-bar {
		display: flex;
		flex-direction: column;
		gap: 12px;
		align-items: center;
		margin-top: 16px;
	}
	
	.filter-controls {
		display: flex;
		gap: 24px;
		flex-wrap: nowrap;
		justify-content: center;
		max-width: 90vw;
		overflow-x: auto;
		scrollbar-width: none; /* Firefox */
		-ms-overflow-style: none; /* IE/Edge */
	}
	.filter-controls::-webkit-scrollbar {
		display: none; /* Chrome/Safari */
	}
	
	.filter-group {
		display: flex;
		align-items: center;
		gap: 8px;
	}
	
	.label {
		font-size: 0.85rem;
		color: var(--text-muted);
		font-weight: 500;
	}
	
	.filter-pill {
		background: var(--bg-surface);
		border: 2px solid var(--border-subtle);
		color: var(--text-muted);
		padding: 4px 12px;
		font-size: 0.8rem;
		font-weight: 600;
		font-family: var(--font-mono);
		cursor: pointer;
		transition: all 0.2s ease;
		text-transform: uppercase;
		box-shadow: 2px 2px 0 rgba(0,0,0,1);
	}
	
	.filter-pill:hover {
		transform: translate(-1px, -1px);
		box-shadow: 3px 3px 0 var(--color-primary);
	}
	
	.filter-pill.active {
		background: var(--text-main);
		border-color: var(--text-main);
		color: var(--bg-surface);
		box-shadow: 2px 2px 0 var(--color-primary);
	}
	
	.active-filters {
		display: flex;
		align-items: center;
		gap: 12px;
		background: var(--bg-surface);
		padding: 8px 16px;
		font-size: 0.85rem;
		font-family: var(--font-mono);
		font-weight: 600;
		border: 2px solid var(--border-subtle);
		box-shadow: 4px 4px 0 rgba(0,0,0,1);
		text-transform: uppercase;
	}
	
	.count {
		color: var(--text-muted);
		font-weight: 600;
	}
	
	.active-list {
		display: flex;
		align-items: center;
		gap: 8px;
	}
	
	.active-item {
		color: var(--text-main);
		text-transform: capitalize;
	}
	
	.dot {
		color: var(--text-muted);
		opacity: 0.5;
	}
</style>
