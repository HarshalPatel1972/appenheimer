<script lang="ts">
	import type { PlacedResult } from '$lib/search/layout/types';
	import { hoverStore } from '$lib/state/hover.svelte';
	import { detailsStore } from '$lib/state/details.svelte';
	import { prefetchAppDetails } from '$lib/services/apps';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { elasticOut } from 'svelte/easing';
	import { getIconUrl } from '$lib/utils/icons';
	
	let { result, centerX, centerY }: { result: PlacedResult, centerX: number, centerY: number } = $props();
	
	let cardElement: HTMLElement;
	let prefetchTimeout: ReturnType<typeof setTimeout> | null = null;
	
	function handleClick() {
		const q = $page.url.searchParams.get('q') || '';
		if (isActive) {
			// Toggle off
			goto(`/?q=${q}`, { keepFocus: true, noScroll: true, replaceState: false });
		} else {
			goto(`/?q=${q}&app=${result.app.id}`, { keepFocus: true, noScroll: true, replaceState: false });
		}
	}
	
	function handleEnter() {
		if (isActive) return;
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
	
	// Active: this specific icon is the selected app
	let isActive = $derived(detailsStore.activeAppId === result.app.id);
	let isHovered = $derived(hoverStore.hoveredApp?.id === result.app.id && !isActive);
	// Dimmed: any other icon when something is hovered OR when any app is active
	let isDimmed = $derived(
		(!isActive && !isHovered && hoverStore.hoveredApp !== null) ||
		(!isActive && detailsStore.activeAppId !== null)
	);
	let isDown = $derived(result.app.status === 'Outage' || result.app.status === 'Maintenance');

	let repulsionX = $state(0);
	let repulsionY = $state(0);
	let iconError = $state(false);
	
	// Repulsion effect from hovered icon — only when nothing is active
	$effect(() => {
		if (detailsStore.activeAppId) {
			repulsionX = 0;
			repulsionY = 0;
			return;
		}
		if (hoverStore.hoveredApp && !isHovered && hoverStore.anchorRect && cardElement) {
			const hoveredCenterX = hoverStore.anchorRect.x + hoverStore.anchorRect.width / 2;
			const hoveredCenterY = hoverStore.anchorRect.y + hoverStore.anchorRect.height / 2;
			const myRect = cardElement.getBoundingClientRect();
			const myCenterX = myRect.x + myRect.width / 2;
			const myCenterY = myRect.y + myRect.height / 2;
			
			const dx = myCenterX - hoveredCenterX;
			const dy = myCenterY - hoveredCenterY;
			const distance = Math.sqrt(dx * dx + dy * dy);
			
			if (distance < 120 && distance > 0) {
				const force = (120 - distance) / 120;
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

	// ── Position logic ──────────────────────────────────────────────────────
	// When active: icon animates from wherever it was floating to sit at the
	// left edge of the original 600px centred search-bar block.
	// The centre-stage in SearchPage has max-width 600px centred, so its left
	// edge is at  centerX - 300.  We want the 80px icon top-left there.
	const ACTIVE_SIZE = 80;
	const SEARCH_BLOCK_HALF = 300; // half of 600px centre-stage

	let targetX = $derived(
		isActive
			? Math.max(8, centerX - SEARCH_BLOCK_HALF)   // left edge of search block
			: result.layout.x + repulsionX
	);
	let targetY = $derived(
		isActive
			? centerY - ACTIVE_SIZE / 2                   // vertically centred
			: result.layout.y + repulsionY
	);

	let opacity = $derived(isActive ? 1 : (detailsStore.activeAppId ? 0 : 1));
	let scaleVal = $derived(isHovered && !isActive ? 1.15 : 1);

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
	class:active={isActive}
	style="transform: translate({targetX}px, {targetY}px) scale({scaleVal}); opacity: {opacity}; pointer-events: {detailsStore.activeAppId && !isActive ? 'none' : 'auto'}; z-index: {isActive ? 210 : 'auto'};"
	tabindex={detailsStore.activeAppId && !isActive ? -1 : 0}
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
		/* Smooth position, opacity, size and filter all at once */
		transition:
			transform 0.55s cubic-bezier(0.22, 1, 0.36, 1),
			opacity 0.3s ease,
			filter 0.25s ease,
			width 0.45s cubic-bezier(0.22, 1, 0.36, 1),
			height 0.45s cubic-bezier(0.22, 1, 0.36, 1);
		outline: none;
		user-select: none;
	}
	
	/* ── Active: grows to prominent badge ─────────────────────────────── */
	.result-icon-item.active {
		width: 80px;
		height: 92px;
		cursor: pointer;
	}

	.result-icon-item.active .icon-wrapper {
		width: 80px;
		height: 80px;
		border: 2px solid var(--border-subtle);
		box-shadow: 6px 6px 0 rgba(0,0,0,1);
	}

	.result-icon-item.active .icon-label {
		font-size: 0.7rem;
		font-weight: 800;
		max-width: 90px;
		letter-spacing: 0.04em;
	}

	/* ── Hovered ───────────────────────────────────────────────────────── */
	.result-icon-item:focus-visible .icon-wrapper {
		border-color: var(--color-primary);
		box-shadow: 0 0 0 3px var(--color-primary);
	}
	
	.result-icon-item.hovered .icon-wrapper {
		border-color: var(--color-primary);
		box-shadow: 0 4px 12px rgba(217, 56, 30, 0.3);
		transform: translateY(-2px);
	}

	/* ── Dimmed: blur + grayscale for spotlight effect (Issue 11) ──────── */
	.result-icon-item.dimmed {
		filter: blur(1.5px) grayscale(1) opacity(0.25);
	}
	
	.result-icon-item.down {
		filter: grayscale(1) opacity(0.3);
	}
	
	/* ── Icon box ──────────────────────────────────────────────────────── */
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
		transition: transform 0.2s, box-shadow 0.2s, border-color 0.2s, width 0.45s cubic-bezier(0.22, 1, 0.36, 1), height 0.45s cubic-bezier(0.22, 1, 0.36, 1);
	}
	
	.icon-img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.icon-fallback {
		font-size: 1.4rem;
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
		transition: font-size 0.35s ease, max-width 0.35s ease;
	}
</style>
