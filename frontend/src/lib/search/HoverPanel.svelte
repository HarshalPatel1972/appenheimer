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
		border: 1px solid var(--border-subtle);
		border-radius: 12px;
		padding: 20px;
		box-sizing: border-box;
		box-shadow: 0 16px 40px rgba(0,0,0,0.5);
		z-index: 1000;
		pointer-events: auto;
		transition: transform 0.25s cubic-bezier(0.22, 1, 0.36, 1);
		animation: fadeIn 0.15s ease-out;
	}
	
	@keyframes fadeIn {
		from { opacity: 0; filter: blur(4px); }
		to { opacity: 1; filter: blur(0px); }
	}
	
	.panel-header {
		display: flex;
		justify-content: space-between;
		align-items: flex-start;
		margin-bottom: 8px;
	}
	
	.panel-header h3 {
		margin: 0;
		font-size: 1.15rem;
		font-weight: 600;
		color: var(--text-main);
	}
	
	.desc {
		font-size: 0.95rem;
		color: var(--text-muted);
		margin: 0 0 16px 0;
		line-height: 1.5;
	}
	
	.tags {
		display: flex;
		flex-wrap: wrap;
		gap: 8px;
	}
	
	.tag {
		background: var(--bg-primary);
		color: var(--text-muted);
		font-size: 0.75rem;
		padding: 4px 8px;
		border-radius: 6px;
		text-transform: capitalize;
		border: 1px solid var(--border-subtle);
	}
	
	.status-badge {
		font-size: 0.75rem;
		padding: 3px 8px;
		border-radius: 99px;
		font-weight: 600;
		text-transform: uppercase;
		letter-spacing: 0.5px;
	}
	
	.status-badge[data-status="Operational"] { background: rgba(16, 185, 129, 0.15); color: #34d399; }
	.status-badge[data-status="Degraded"] { background: rgba(245, 158, 11, 0.15); color: #fbbf24; }
	.status-badge[data-status="Outage"] { background: rgba(239, 68, 68, 0.15); color: #f87171; }
	.status-badge[data-status="Maintenance"] { background: rgba(59, 130, 246, 0.15); color: #60a5fa; }
</style>
