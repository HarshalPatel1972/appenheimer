<script lang="ts">
	import type { PlacedResult } from '$lib/search/layout/types';
	import { hoverStore } from '$lib/state/hover.svelte';
	import { detailsStore } from '$lib/state/details.svelte';
	import { prefetchAppDetails } from '$lib/services/apps';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { elasticOut } from 'svelte/easing';
	import { getIconUrl } from '$lib/utils/icons';
	import AppDetailsView from './AppDetailsView.svelte';
	
	let { result, centerX, centerY }: { result: PlacedResult, centerX: number, centerY: number } = $props();
	
	let cardElement: HTMLElement;
	let prefetchTimeout: ReturnType<typeof setTimeout> | null = null;
	
	function handleClick() {
		const q = $page.url.searchParams.get('q') || '';
		if (detailsStore.activeAppId === result.app.id) {
			goto(`/?q=${q}`, { keepFocus: true, noScroll: true, replaceState: false });
		} else {
			goto(`/?q=${q}&app=${result.app.id}`, { keepFocus: true, noScroll: true, replaceState: false });
		}
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
	
	let isActive = $derived(detailsStore.activeAppId === result.app.id);
	let isHovered = $derived(hoverStore.hoveredApp?.id === result.app.id && !isActive);
	let isDimmed = $derived((hoverStore.hoveredApp !== null && !isHovered) || (detailsStore.activeAppId !== null && !isActive));
	let isDown = $derived(result.app.status === 'Outage' || result.app.status === 'Maintenance');

	let repulsionX = $state(0);
	let repulsionY = $state(0);
	let iconError = $state(false);
	
	$effect(() => {
		if (hoverStore.hoveredApp && !isHovered && hoverStore.anchorRect && cardElement) {
			const hoveredCenterX = hoverStore.anchorRect.x + hoverStore.anchorRect.width / 2;
			const hoveredCenterY = hoverStore.anchorRect.y + hoverStore.anchorRect.height / 2;
			const myRect = cardElement.getBoundingClientRect();
			const myCenterX = myRect.x + myRect.width / 2;
			const myCenterY = myRect.y + myRect.height / 2;
			
			const dx = myCenterX - hoveredCenterX;
			const dy = myCenterY - hoveredCenterY;
			const distance = Math.sqrt(dx * dx + dy * dy);
			
			if (distance < 350 && distance > 0) {
				const force = (350 - distance) / 350;
				repulsionX = (dx / distance) * force * 12;
				repulsionY = (dy / distance) * force * 12;
			} else {
				repulsionX = 0;
				repulsionY = 0;
			}
		} else {
			repulsionX = 0;
			repulsionY = 0;
		}
	});

	let targetX = $derived(result.layout.x + repulsionX);
	let targetY = $derived(result.layout.y + repulsionY);

	function constellation(node: HTMLElement, { delay = 0, duration = 1000 }) {
		const dx = centerX - result.layout.x;
		const dy = centerY - result.layout.y;
		
		return {
			delay,
			duration,
			easing: elasticOut,
			css: (t: number, u: number) => `
				transform: translate(${result.layout.x + dx * u}px, ${result.layout.y + dy * u}px) scale(${t});
				opacity: ${t};
			`
		};
	}
</script>

<li 
	bind:this={cardElement}
	class="result-icon-item" 
	class:hovered={isHovered}
	class:dimmed={isDimmed}
	class:down={isDown}
	style="transform: translate({targetX}px, {targetY}px) scale({isHovered ? 1.15 : 1}); opacity: {detailsStore.activeAppId ? 0 : 1}; pointer-events: {detailsStore.activeAppId ? 'none' : 'auto'};"
	tabindex="0"
	in:constellation={{ delay: Math.random() * 150 }}
	onmouseenter={handleEnter}
	onmouseleave={handleLeave}
	onfocus={handleEnter}
	onblur={handleLeave}
	onclick={handleClick}
>
	<div class="icon-wrapper">
		{#if !iconError}
			<img src={getIconUrl(result.app.icon, result.app.name)} alt={result.app.name} class="icon-img" onerror={() => iconError = true} />
		{:else}
			<div class="icon-fallback">{result.app.name.charAt(0)}</div>
		{/if}
	</div>
	<span class="icon-label">{result.app.name}</span>
</li>

<style>
	.result-icon-item {
		position: absolute;
		top: 0;
		left: 0;
		width: 56px;
		height: 64px;
		list-style: none;
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 4px;
		cursor: pointer;
		transition: transform 0.4s cubic-bezier(0.22, 1, 0.36, 1), 
					opacity 0.3s ease, 
					filter 0.3s ease;
		outline: none;
		user-select: none;
	}
	
	.result-icon-item:focus-visible .icon-wrapper {
		border-color: var(--color-primary);
		box-shadow: 0 0 0 3px var(--color-primary);
	}
	
	.result-icon-item.hovered .icon-wrapper {
		border-color: var(--color-primary);
		box-shadow: 0 4px 12px rgba(217, 56, 30, 0.3);
		transform: translateY(-2px);
	}
	
	.result-icon-item.dimmed {
		filter: grayscale(1) opacity(0.3);
	}
	
	.result-icon-item.down {
		filter: grayscale(1) opacity(0.3);
	}
	
	.icon-wrapper {
		width: 44px;
		height: 44px;
		flex-shrink: 0;
		overflow: hidden;
		background: #ffffff;
		border: 1.5px solid var(--border-subtle);
		display: flex;
		align-items: center;
		justify-content: center;
		box-shadow: 2px 2px 0 rgba(0,0,0,1);
		transition: transform 0.2s, box-shadow 0.2s, border-color 0.2s;
	}
	
	.icon-img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.icon-fallback {
		font-size: 1.2rem;
		font-weight: 800;
		color: var(--text-main);
		text-transform: uppercase;
	}
	
	.icon-label {
		font-size: 0.65rem;
		font-weight: 700;
		color: var(--text-main);
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
		max-width: 64px;
		text-align: center;
		text-transform: uppercase;
		letter-spacing: 0.03em;
		font-family: var(--font-mono);
	}
</style>
