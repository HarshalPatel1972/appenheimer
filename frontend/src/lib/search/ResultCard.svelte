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

	let targetX = $derived(isActive ? (centerX - 60) : (detailsStore.activeAppId ? (result.layout.x < centerX ? -500 : 2500) : (result.layout.x + repulsionX)));
	let targetY = $derived(isActive ? (centerY) : (detailsStore.activeAppId ? result.layout.y : (result.layout.y + repulsionY)));

	function constellation(node: HTMLElement, { delay = 0, duration = 1200 }) {
		const targetX = result.layout.x;
		const targetY = result.layout.y;
		const dx = centerX - targetX;
		const dy = centerY - targetY;
		
		return {
			delay,
			duration,
			easing: elasticOut,
			css: (t: number, u: number) => `
				transform: translate(${targetX + dx * u}px, ${targetY + dy * u}px) scale(${t});
				opacity: ${t};
			`
		};
	}
</script>

<li 
	bind:this={cardElement}
	class="result-card" 
	class:hovered={isHovered}
	class:dimmed={isDimmed}
	class:down={isDown}
	class:active={isActive}
	style="transform: translate({targetX}px, {targetY}px) scale({isHovered ? 1.05 : 1}); opacity: {detailsStore.activeAppId && !isActive ? 0 : 1}; pointer-events: {detailsStore.activeAppId && !isActive ? 'none' : 'auto'};"
	tabindex="0"
	in:constellation={{ delay: Math.random() * 200 }}
	onmouseenter={handleEnter}
	onmouseleave={handleLeave}
	onfocus={handleEnter}
	onblur={handleLeave}
	onclick={handleClick}
>
	<div class="card-content">
		<div class="icon-container">
			{#if !iconError}
				<img src={getIconUrl(result.app.icon, result.app.name)} alt={result.app.name} class="icon-img" onerror={() => iconError = true} />
			{:else}
				<div class="icon-fallback">{result.app.name.charAt(0)}</div>
			{/if}
		</div>
		<div class="title-section">
			<h3>{result.app.name}</h3>
		</div>
	</div>

	{#if isActive}
		<AppDetailsView app={result.app} details={detailsStore.appDetails} />
	{/if}
</li>

<style>
	.result-card {
		position: absolute;
		top: 0;
		left: 0;
		width: 120px;
		height: 120px;
		list-style: none;
		background: var(--bg-surface);
		border: 2px solid var(--border-subtle);
		padding: 16px;
		box-sizing: border-box;
		cursor: pointer;
		box-shadow: 4px 4px 0 rgba(0,0,0,1);
		transition: transform 0.6s cubic-bezier(0.22, 1, 0.36, 1), 
					width 0.4s cubic-bezier(0.22, 1, 0.36, 1), 
					height 0.4s cubic-bezier(0.22, 1, 0.36, 1),
					border-color 0.2s ease, 
					opacity 0.4s ease, 
					filter 0.4s ease, 
					box-shadow 0.4s ease;
		outline: none;
		overflow: hidden;
	}
	
	.result-card.active {
		z-index: 1000;
		cursor: default;
		background: transparent;
		border: none;
		box-shadow: none;
		overflow: visible;
	}

	.result-card:focus-visible {
		border-color: var(--color-primary);
		box-shadow: 4px 4px 0 var(--color-primary);
	}
	
	.result-card.hovered {
		border-color: var(--color-primary);
		z-index: 10;
		box-shadow: 6px 6px 0 var(--color-primary);
	}
	
	.result-card.dimmed {
		filter: grayscale(1);
	}
	
	.result-card.down {
		filter: grayscale(1) opacity(0.5);
	}
	
	.result-card.down.hovered {
		filter: grayscale(1) opacity(1);
	}
	
	.card-content {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		gap: 12px;
		height: 100%;
	}
	
	.icon-container {
		width: 52px;
		height: 52px;
		flex-shrink: 0;
		overflow: hidden;
		background: var(--bg-primary);
		display: flex;
		align-items: center;
		justify-content: center;
		border: 1px solid var(--border-subtle);
	}
	
	.icon-img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.icon-fallback {
		font-size: 1.5rem;
		font-weight: 700;
		color: var(--text-main);
		text-transform: uppercase;
	}
	
	.title-section {
		width: 100%;
		text-align: center;
		overflow: hidden;
	}
	
	.title-section h3 {
		margin: 0;
		font-size: 0.9rem;
		font-weight: 600;
		color: var(--text-main);
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
		text-transform: uppercase;
		letter-spacing: 0.05em;
	}
</style>
