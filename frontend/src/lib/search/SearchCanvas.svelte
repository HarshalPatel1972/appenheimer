<script lang="ts">
	import { searchStore } from '$lib/state/search.svelte';
	import ResultCard from './ResultCard.svelte';
	import HoverPanel from './HoverPanel.svelte';
	import { calculateLayout } from './layout/engine';
	import type { PlacedResult } from './layout/types';
	
	let containerWidth = $state(0);
	let containerHeight = $state(0);
	
	let placedResults = $derived.by(() => {
		if (searchStore.status === 'success' && containerWidth > 0 && containerHeight > 0) {
			return calculateLayout(searchStore.query, searchStore.results, containerWidth, containerHeight);
		}
		return [];
	});
</script>

<div 
	class="canvas-container" 
	bind:clientWidth={containerWidth} 
	bind:clientHeight={containerHeight}
>
	{#if searchStore.status === 'success'}
		<ul role="list" class="canvas-list">
			{#each placedResults as result (result.app.id)}
				<ResultCard {result} />
			{/each}
		</ul>
		<HoverPanel />
	{/if}
</div>

<style>
	.canvas-container {
		width: 100vw;
		height: 100vh;
		position: absolute;
		top: 0;
		left: 0;
		overflow: hidden;
		pointer-events: none;
	}
	
	.canvas-list {
		margin: 0;
		padding: 0;
		width: 100%;
		height: 100%;
		pointer-events: auto;
	}
</style>
