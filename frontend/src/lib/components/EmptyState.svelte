<script lang="ts">
	import { searchStore } from '$lib/state/search.svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	let { onChipClick }: { onChipClick: (keyword: string) => void } = $props();

	let activeFilterCount = $derived(searchStore.activeFiltersOrder.length);

	function clearFilters() {
		const url = new URL($page.url);
		url.searchParams.delete('platform');
		url.searchParams.delete('pricing');
		url.searchParams.delete('category');
		goto(url, { keepFocus: true, noScroll: true });
	}

	const sampleKeywords = ['design', 'video editor', 'windows', 'free', 'prototyping', 'IDE', 'music production', 'browser', 'office'];
</script>

<div class="empty-state">
	{#if activeFilterCount > 0 && searchStore.results.length === 0 && searchStore.status === 'success'}
		<h3>No apps match these filters.</h3>
		<p>Try removing some filters to see more results.</p>
		<button class="clear-btn" onclick={clearFilters}>Clear Filters</button>
	{:else if searchStore.status === 'success' && searchStore.results.length === 0 && searchStore.query}
		<div class="no-results">
			<span class="no-results-code">404</span>
			<p class="no-results-msg">No results for "<strong>{searchStore.query}</strong>"</p>
			<p class="no-results-hint">Try a broader term or different spelling.</p>
		</div>
	{:else if searchStore.status === 'idle'}
		<div class="floating-keywords">
			{#each sampleKeywords as kw, i}
				<button
					class="kw"
					style="animation-delay: {i * 0.35}s"
					onclick={() => onChipClick(kw)}
					title="Search for {kw}"
				>
					{kw}
				</button>
			{/each}
		</div>
	{/if}
</div>

<style>
	.empty-state {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		height: 100%;
		text-align: center;
	}

	.floating-keywords {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		display: flex;
		gap: 16px;
		flex-wrap: wrap;
		max-width: 560px;
		justify-content: center;
		/* Push below the search bar (which is at center) */
		margin-top: 130px;
	}

	/* Clickable keyword chip */
	.kw {
		background: var(--bg-surface);
		border: 2px solid var(--border-subtle);
		padding: 8px 16px;
		font-family: var(--font-mono);
		font-weight: 600;
		font-size: 0.8rem;
		text-transform: uppercase;
		letter-spacing: 0.05em;
		box-shadow: 3px 3px 0 rgba(0,0,0,1);
		color: var(--text-muted);
		animation: floatChip 4s ease-in-out infinite alternate;
		cursor: pointer;
		transition: color 0.2s, border-color 0.2s, box-shadow 0.2s, transform 0.15s;
	}

	.kw:hover {
		color: var(--text-main);
		border-color: var(--color-primary);
		box-shadow: 4px 4px 0 var(--color-primary);
		transform: translate(-1px, -1px);
	}

	@keyframes floatChip {
		0%   { transform: translateY(0px) rotate(-0.5deg); }
		100% { transform: translateY(-8px) rotate(0.5deg); }
	}

	.kw:hover {
		animation-play-state: paused;
	}

	/* No-results */
	.no-results {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 8px;
		transform: translateY(-80px);
	}

	.no-results-code {
		font-size: 5rem;
		font-weight: 800;
		font-family: var(--font-mono);
		color: var(--border-subtle);
		letter-spacing: -0.04em;
		line-height: 1;
	}

	.no-results-msg {
		margin: 0;
		font-size: 1rem;
		font-weight: 600;
		color: var(--text-main);
	}

	.no-results-hint {
		margin: 0;
		font-size: 0.85rem;
		color: var(--text-muted);
	}

	h3 {
		margin: 0 0 8px 0;
		color: var(--text-main);
		font-weight: 700;
		text-transform: uppercase;
		font-family: var(--font-mono);
	}

	p {
		margin: 0;
		font-size: 1rem;
		max-width: 300px;
		line-height: 1.5;
		color: var(--text-muted);
	}

	.clear-btn {
		margin-top: 24px;
		padding: 10px 20px;
		background: var(--bg-surface);
		border: 2px solid var(--border-subtle);
		color: var(--text-main);
		box-shadow: 4px 4px 0 rgba(0,0,0,1);
		cursor: pointer;
		font-weight: 700;
		text-transform: uppercase;
		font-family: var(--font-mono);
		letter-spacing: 0.05em;
		transition: all 0.2s ease;
	}

	.clear-btn:hover {
		transform: translate(-2px, -2px);
		box-shadow: 6px 6px 0 var(--color-primary);
	}
</style>
