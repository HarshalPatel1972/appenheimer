<script lang="ts">
	import { searchStore } from '$lib/state/search.svelte';
	import { performSearch } from '$lib/services/search';

	function handleKeydown(event: KeyboardEvent) {
		if (event.key === 'Enter') {
			performSearch();
		}
	}
</script>

<div class="search-input-container">
	<input
		type="text"
		placeholder="What do you want to do?"
		bind:value={searchStore.query}
		oninput={() => searchStore.setTyping()}
		onkeydown={handleKeydown}
		class="search-input"
		autofocus
	/>
</div>

<style>
	.search-input-container {
		position: fixed;
		top: 40px;
		left: 50%;
		transform: translateX(-50%);
		z-index: 1000;
		width: 100%;
		max-width: 600px;
		padding: 0 20px;
		box-sizing: border-box;
	}
	.search-input {
		width: 100%;
		padding: 18px 28px;
		font-size: 1.25rem;
		background: var(--bg-surface);
		border: 1px solid var(--border-subtle);
		color: var(--text-main);
		border-radius: 99px;
		outline: none;
		box-shadow: 0 8px 30px rgba(0,0,0,0.4);
		transition: border-color 0.2s ease, box-shadow 0.2s ease;
	}
	.search-input:focus {
		border-color: var(--color-primary);
		box-shadow: 0 8px 30px rgba(59, 130, 246, 0.2);
	}
</style>
