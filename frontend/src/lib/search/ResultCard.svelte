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

	let targetX = $derived(isActive ? (centerX - 400) : (result.layout.x + repulsionX));
	let targetY = $derived(isActive ? (centerY - 300) : (result.layout.y + repulsionY));

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
	style="transform: translate({targetX}px, {targetY}px) scale({isHovered ? 1.05 : 1});"
	tabindex="0"
	in:constellation={{ delay: Math.random() * 200 }}
	onmouseenter={handleEnter}
	onmouseleave={handleLeave}
	onfocus={handleEnter}
	onblur={handleLeave}
	onclick={handleClick}
>
	{#if isActive}
		<AppDetailsView app={result.app} details={detailsStore.appDetails} />
	{:else}
		<div class="card-content">
			<div class="icon-container">
				{#if !iconError}
					<img src={getIconUrl(result.app.icon, result.app.name)} alt={result.app.name} class="icon-img" onerror={() => iconError = true} />
				{:else}
					<div class="icon-fallback">{result.app.name.charAt(0)}</div>
				{/if}
			</div>
			<div class="details">
				<h3>{result.app.name}</h3>
				<p>{result.app.description || 'Application'}</p>
			</div>
		</div>
	{/if}
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
		box-shadow: 0 2px 8px rgba(0,0,0,0.05);
		transition: transform 0.4s cubic-bezier(0.22, 1, 0.36, 1), 
					width 0.4s cubic-bezier(0.22, 1, 0.36, 1), 
					height 0.4s cubic-bezier(0.22, 1, 0.36, 1),
					border-radius 0.4s ease,
					border-color 0.2s ease, 
					opacity 0.4s ease, 
					filter 0.4s ease, 
					box-shadow 0.4s ease;
		outline: none;
		overflow: hidden;
	}
	
	.result-card.active {
		width: 800px;
		height: 600px;
		z-index: 1000;
		cursor: default;
		border-radius: 24px;
		box-shadow: 0 24px 64px rgba(0,0,0,0.2);
		background: var(--bg-canvas); /* slightly different background for expanded view */
		border-color: var(--border-subtle);
	}

	.result-card:focus-visible {
		border-color: var(--color-primary);
		box-shadow: 0 0 0 2px var(--bg-surface), 0 0 0 4px var(--color-primary);
	}
	
	.result-card.hovered {
		border-color: var(--color-primary);
		z-index: 10;
		box-shadow: 0 12px 30px rgba(0,0,0,0.15);
	}
	
	.result-card.dimmed {
		opacity: 0.15;
		filter: blur(5px) grayscale(0.5);
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
	
	.icon-container {
		width: 48px;
		height: 48px;
		flex-shrink: 0;
		border-radius: 10px;
		overflow: hidden;
		background: var(--bg-primary);
		display: flex;
		align-items: center;
		justify-content: center;
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
