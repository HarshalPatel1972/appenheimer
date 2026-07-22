<script lang="ts">
	import { searchStore } from '$lib/state/search.svelte';
	import SearchCanvas from './SearchCanvas.svelte';
	import EmptyState from './EmptyState.svelte';
	import LoadingState from './LoadingState.svelte';
	import ErrorState from './ErrorState.svelte';
	import TypingState from './TypingState.svelte';
</script>

<div class="search-view">
	{#if searchStore.status === 'idle'}
		<EmptyState />
	{:else if searchStore.status === 'typing'}
		<TypingState />
	{:else if searchStore.status === 'loading'}
		<LoadingState />
	{:else if searchStore.status === 'error'}
		<ErrorState message={searchStore.error} />
	{:else if searchStore.status === 'empty'}
		<EmptyState />
	{/if}
	
	<!-- Canvas is always mounted to maintain width/height bounds, only renders when success -->
	<SearchCanvas />
</div>

<style>
	.search-view {
		width: 100vw;
		height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		position: relative;
	}
</style>
