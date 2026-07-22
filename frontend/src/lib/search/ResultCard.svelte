<script lang="ts">
	import type { PlacedResult } from '$lib/search/layout/types';
	import { hoverStore } from '$lib/state/hover.svelte';
	import { prefetchAppDetails } from '$lib/services/apps';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	
	let { result }: { result: PlacedResult } = $props();
	
	let cardElement: HTMLElement;
	let prefetchTimeout: ReturnType<typeof setTimeout> | null = null;
	
	function handleClick() {
		const q = $page.url.searchParams.get('q') || '';
		goto(`/?q=${q}&app=${result.app.id}`, { keepFocus: true, noScroll: true, replaceState: false });
	}
	
	function handleEnter() {
		if (cardElement) {
			hoverStore.setHover(result.app, cardElement.getBoundingClientRect());
		}
		if (prefetchTimeout) clearTimeout(prefetchTimeout);
		prefetchTimeout = setTimeout(() => {
			prefetchAppDetails(result.app.id);
		}, 300);
	}
	
	function handleLeave() {
		hoverStore.clearHover(result.app.id);
		if (prefetchTimeout) clearTimeout(prefetchTimeout);
	}
	
	let isHovered = $derived(hoverStore.hoveredApp?.id === result.app.id);
	let isDimmed = $derived(hoverStore.hoveredApp !== null && !isHovered);
	let isDown = $derived(result.app.status === 'Outage' || result.app.status === 'Maintenance');
</script>

<li 
	bind:this={cardElement}
	class="result-card" 
	class:hovered={isHovered}
	class:dimmed={isDimmed}
	class:down={isDown}
	style="transform: translate({result.layout.x}px, {result.layout.y}px) scale({isHovered ? 1.05 : 1});"
	tabindex="0"
	onmouseenter={handleEnter}
	onmouseleave={handleLeave}
	onfocus={handleEnter}
	onblur={handleLeave}
	onclick={handleClick}
>
	<div class="card-content">
		<div class="icon"></div>
		<div class="details">
			<h3>{result.app.name}</h3>
			<p>{result.app.description || 'Application'}</p>
		</div>
	</div>
</li>

<style>
	.result-card {
		position: absolute;
		top: 0;
		left: 0;
		width: 260px;
		height: 84px;
		list-style: none;
		background: var(--bg-surface);
		border: 1px solid var(--border-subtle);
		border-radius: 12px;
		padding: 16px;
		box-sizing: border-box;
		cursor: pointer;
		transition: transform 0.3s cubic-bezier(0.22, 1, 0.36, 1), border-color 0.2s ease, opacity 0.3s ease, filter 0.3s ease, box-shadow 0.3s ease;
		outline: none;
	}
	
	.result-card:focus-visible {
		border-color: var(--color-primary);
		box-shadow: 0 0 0 2px var(--bg-surface), 0 0 0 4px var(--color-primary);
	}
	
	.result-card.hovered {
		border-color: var(--color-primary);
		z-index: 10;
		box-shadow: 0 12px 30px rgba(0,0,0,0.4);
	}
	
	.result-card.dimmed {
		opacity: 0.3;
		filter: blur(3px);
	}
	
	.result-card.down {
		filter: grayscale(1) opacity(0.5);
	}
	
	.result-card.down.hovered {
		filter: grayscale(1) opacity(1);
	}
	
	.card-content {
		display: flex;
		align-items: center;
		gap: 14px;
		height: 100%;
	}
	
	.icon {
		width: 48px;
		height: 48px;
		background: var(--bg-primary);
		border-radius: 10px;
		flex-shrink: 0;
	}
	
	.details {
		flex: 1;
		overflow: hidden;
	}
	
	.details h3 {
		margin: 0;
		font-size: 1rem;
		font-weight: 600;
		color: var(--text-main);
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
	}
	
	.details p {
		margin: 4px 0 0;
		font-size: 0.85rem;
		color: var(--text-muted);
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
	}
</style>
