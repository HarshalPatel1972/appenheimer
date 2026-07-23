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

	const sampleKeywords = ['design', 'video editor', 'windows', 'free', 'prototyping', 'IDE', 'music production'];
</script>

<div class="empty-state">
	{#if activeFilterCount > 0 && searchStore.results.length === 0 && searchStore.status === 'success'}
		<h3>No apps match these filters.</h3>
		<p>Try removing some filters to see more results.</p>
		<button class="clear-btn" onclick={clearFilters}>Clear Filters</button>
	{:else}
		<div class="hero-copy">
			<h2>Find software.</h2>
			<h2 class="dim">Not websites.</h2>
			<h2 class="dim">Not articles.</h2>
			<h2 class="highlight">Just the right tool.</h2>
		</div>
		
		<div class="floating-keywords">
			{#each sampleKeywords as kw, i}
				<span class="kw" style="animation-delay: {i * 0.4}s">{kw}</span>
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

	.hero-copy {
		display: flex;
		flex-direction: column;
		gap: 8px;
		margin-bottom: 48px;
	}

	.hero-copy h2 {
		margin: 0;
		font-size: 2.5rem;
		font-weight: 700;
		letter-spacing: -0.03em;
		color: var(--text-main);
	}

	.hero-copy .dim {
		color: var(--text-muted);
		opacity: 0.5;
	}

	.hero-copy .highlight {
		color: var(--color-primary);
		margin-top: 16px;
	}

	.floating-keywords {
		display: flex;
		gap: 16px;
		flex-wrap: wrap;
		max-width: 500px;
		justify-content: center;
		opacity: 0.6;
	}

	.kw {
		background: rgba(255,255,255,0.03);
		border: 1px solid rgba(255,255,255,0.1);
		padding: 8px 16px;
		border-radius: 99px;
		color: var(--text-muted);
		animation: float 4s ease-in-out infinite alternate;
	}

	@keyframes float {
		0% { transform: translateY(0px); }
		100% { transform: translateY(-10px); }
	}

	h3 {
		margin: 0 0 8px 0;
		color: var(--text-main);
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
