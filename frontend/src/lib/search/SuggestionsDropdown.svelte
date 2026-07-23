<script lang="ts">
	import { searchStore } from '$lib/state/search.svelte';
	import { getIconUrl } from '$lib/utils/icons';

	let { onSelect }: { onSelect: (name: string) => void } = $props();

	let suggestions = $derived(searchStore.results.slice(0, 7));
	let iconErrors = $state<Record<string, boolean>>({});
</script>

{#if suggestions.length > 0}
<div class="suggestions-dropdown" role="listbox" aria-label="Search suggestions">
	{#each suggestions as app, i}
		<button
			class="suggestion-row"
			role="option"
			aria-selected="false"
			onclick={() => onSelect(app.name)}
		>
			<span class="sug-index">{String(i + 1).padStart(2, '0')}</span>
			<div class="sug-icon">
				{#if !iconErrors[app.id]}
					<img
						src={getIconUrl(app.icon, app.name)}
						alt={app.name}
						onerror={() => iconErrors[app.id] = true}
					/>
				{:else}
					<span class="sug-fallback">{app.name.charAt(0)}</span>
				{/if}
			</div>
			<span class="sug-name">{app.name}</span>
			{#if app.categories && app.categories.length > 0}
				<span class="sug-cat">{app.categories[0]}</span>
			{/if}
			<span class="sug-arrow">→</span>
		</button>
	{/each}
</div>
{/if}

<style>
	.suggestions-dropdown {
		position: absolute;
		top: calc(100% + 2px);
		left: 0;
		right: 0;
		background: var(--bg-surface);
		border: 3px solid var(--border-subtle);
		border-top: none;
		box-shadow: 6px 6px 0 rgba(0,0,0,1);
		z-index: 400;
		overflow: hidden;
	}

	.suggestion-row {
		display: flex;
		align-items: center;
		gap: 12px;
		width: 100%;
		padding: 10px 16px;
		background: transparent;
		border: none;
		border-bottom: 1px solid rgba(17,17,17,0.08);
		cursor: pointer;
		text-align: left;
		transition: background 0.1s;
		box-sizing: border-box;
	}

	.suggestion-row:last-child {
		border-bottom: none;
	}

	.suggestion-row:hover {
		background: var(--bg-canvas);
	}

	.suggestion-row:hover .sug-arrow {
		opacity: 1;
		transform: translateX(0);
	}

	.sug-index {
		font-size: 0.6rem;
		font-family: var(--font-mono);
		font-weight: 700;
		color: var(--color-primary);
		width: 18px;
		flex-shrink: 0;
	}

	.sug-icon {
		width: 26px;
		height: 26px;
		flex-shrink: 0;
		overflow: hidden;
		background: #fff;
		border: 1.5px solid var(--border-subtle);
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.sug-icon img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.sug-fallback {
		font-size: 0.8rem;
		font-weight: 800;
		color: var(--text-main);
		text-transform: uppercase;
	}

	.sug-name {
		flex: 1;
		font-size: 0.82rem;
		font-weight: 700;
		color: var(--text-main);
		text-transform: uppercase;
		font-family: var(--font-mono);
		letter-spacing: 0.04em;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	.sug-cat {
		font-size: 0.65rem;
		color: var(--text-muted);
		text-transform: uppercase;
		font-family: var(--font-mono);
		font-weight: 600;
		flex-shrink: 0;
	}

	.sug-arrow {
		font-size: 0.9rem;
		color: var(--color-primary);
		font-weight: 700;
		opacity: 0;
		transform: translateX(-4px);
		transition: all 0.15s;
		flex-shrink: 0;
	}
</style>
