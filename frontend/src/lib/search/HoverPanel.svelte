<script lang="ts">
	import { hoverStore } from '$lib/state/hover.svelte';
	import { calculateHoverPosition } from './layout/hover';
	
	let panelWidth = 320;
	let panelHeight = 160;
	
	let app = $derived(hoverStore.hoveredApp);
	let rect = $derived(hoverStore.anchorRect);
	
	let position = $derived.by(() => {
		if (rect && typeof window !== 'undefined') {
			return calculateHoverPosition(rect, panelWidth, panelHeight, window.innerWidth, window.innerHeight);
		}
		return { x: 0, y: 0 };
	});
</script>

{#if app && rect}
<div 
	class="hover-panel" 
	onmouseenter={() => hoverStore.keepHoverAlive(app.id)}
	onmouseleave={() => hoverStore.clearHover(app.id)} 
	style="transform: translate({position.x}px, {position.y}px); width: {panelWidth}px;"
>
	<div class="panel-header">
		<h3>{app.name}</h3>
		<span class="status-badge" data-status={app.status}>{app.status}</span>
	</div>
	
	<p class="desc">{app.description}</p>
	
	{#if app.categories && app.categories.length > 0}
		<div class="tags">
			{#each app.categories as cat}
				<span class="tag">{cat}</span>
			{/each}
		</div>
	{/if}
</div>
{/if}

<style>
	.hover-panel {
		position: absolute;
		top: 0;
		left: 0;
		background: var(--bg-surface);
		border: 2px solid var(--border-subtle);
		padding: 16px;
		box-sizing: border-box;
		box-shadow: 4px 4px 0 rgba(0,0,0,1);
		z-index: 1000;
		pointer-events: auto;
		transition: transform 0.25s cubic-bezier(0.22, 1, 0.36, 1);
		animation: fadeIn 0.15s ease-out;
	}
	
	@keyframes fadeIn {
		from { opacity: 0; transform: scale(0.96); }
		to { opacity: 1; transform: scale(1); }
	}
	
	.panel-header {
		display: flex;
		justify-content: space-between;
		align-items: flex-start;
		margin-bottom: 8px;
	}
	
	.panel-header h3 {
		margin: 0;
		font-size: 1rem;
		font-weight: 700;
		text-transform: uppercase;
		color: var(--text-main);
	}
	
	.desc {
		font-size: 0.85rem;
		color: var(--text-muted);
		margin: 0 0 12px 0;
		line-height: 1.4;
		font-family: var(--font-mono);
	}
	
	.tags {
		display: flex;
		flex-wrap: wrap;
		gap: 6px;
	}
	
	.tag {
		background: var(--bg-canvas);
		color: var(--text-main);
		font-size: 0.7rem;
		font-weight: 600;
		font-family: var(--font-mono);
		padding: 3px 6px;
		text-transform: uppercase;
		border: 1px solid var(--border-subtle);
	}
	
	.status-badge {
		font-size: 0.7rem;
		padding: 2px 6px;
		font-weight: 700;
		font-family: var(--font-mono);
		text-transform: uppercase;
		letter-spacing: 0.5px;
		border: 1px solid var(--border-subtle);
	}
	
	.status-badge[data-status="Operational"] { background: #10b981; color: #fff; }
	.status-badge[data-status="Degraded"] { background: #f59e0b; color: #fff; }
	.status-badge[data-status="Outage"] { background: #ef4444; color: #fff; }
	.status-badge[data-status="Maintenance"] { background: #3b82f6; color: #fff; }
</style>
