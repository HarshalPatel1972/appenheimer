<script lang="ts">
	import type { AppResult, AppDetails } from '$lib/types/search';
	import { getIconUrl } from '$lib/utils/icons';
	
	let { app, details }: { app: AppResult, details: AppDetails | null } = $props();
	let iconError = $state(false);
</script>

<div class="app-details-box">
	<div class="header-info">
		<p class="developer">{details?.developer || 'Unknown Developer'}</p>
		<p class="description">{details?.longDescription || app.description}</p>
	</div>
	
	<div class="meta-grid">
		<div class="meta-section">
			<h4>Platforms</h4>
			{#if details?.platforms}
				<div class="pill-list">
					{#each details.platforms as p}
						<span class="pill">{p}</span>
					{/each}
				</div>
			{:else}
				<p class="placeholder">Loading...</p>
			{/if}
		</div>

		<div class="meta-section">
			<h4>Pricing</h4>
			{#if details?.pricing}
				<div class="pill-list">
					{#each details.pricing as p}
						<span class="pill price">{p}</span>
					{/each}
				</div>
			{:else}
				<p class="placeholder">Loading...</p>
			{/if}
		</div>

		<div class="meta-section full-width">
			<h4>Categories & Tags</h4>
			<div class="pill-list">
				{#if details?.categories}
					{#each details.categories as c}
						<span class="pill">{c}</span>
					{/each}
				{/if}
				{#if details?.tags}
					{#each details.tags as t}
						<span class="pill outline">{t}</span>
					{/each}
				{/if}
				{#if !details?.categories && !details?.tags}
					<p class="placeholder">Loading...</p>
				{/if}
			</div>
		</div>

		<div class="meta-section full-width">
			<h4>Links & Status</h4>
			<div class="links-container">
				{#if details?.websiteUrl}
					<a href={details.websiteUrl} target="_blank" rel="noopener noreferrer" class="link-btn">
						Visit Website ↗
					</a>
				{:else}
					<p class="placeholder">Loading...</p>
				{/if}
				<div class="status-info">
					<span class="small-label">Last Updated:</span>
					<span class="small-text">{details?.lastUpdate ? new Date(details.lastUpdate).toLocaleDateString() : 'Unknown'}</span>
				</div>
				<div class="status-info">
					<span class="small-label">Status:</span>
					<span class="small-text status-{app.status?.toLowerCase() || 'operational'}">{app.status || 'Operational'}</span>
				</div>
			</div>
		</div>
	</div>
</div>

<style>
	.app-details-box {
		position: absolute;
		top: 140px;
		left: 0;
		width: 600px;
		background: var(--bg-surface);
		border: 2px solid var(--border-subtle);
		box-shadow: 6px 6px 0 rgba(0,0,0,1);
		padding: 24px;
		box-sizing: border-box;
		color: var(--text-main);
		animation: slideDown 0.4s cubic-bezier(0.22, 1, 0.36, 1) forwards 0.2s;
		opacity: 0;
		transform: translateY(-20px);
	}
	
	@keyframes slideDown {
		to { 
			opacity: 1; 
			transform: translateY(0);
		}
	}

	.header-info {
		margin-bottom: 24px;
		padding-bottom: 24px;
		border-bottom: 2px solid var(--border-subtle);
	}

	.developer {
		margin: 0 0 12px 0;
		font-size: 0.85rem;
		color: var(--text-muted);
		text-transform: uppercase;
		letter-spacing: 0.1em;
		font-weight: 700;
	}

	.description {
		font-size: 1.1rem;
		line-height: 1.6;
		color: var(--text-main);
		margin: 0;
		font-family: var(--font-mono);
	}

	.meta-grid {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 24px;
	}

	.meta-section {
		display: flex;
		flex-direction: column;
	}

	.full-width {
		grid-column: 1 / -1;
	}

	h4 {
		margin: 0 0 12px 0;
		color: var(--text-main);
		font-size: 0.9rem;
		text-transform: uppercase;
		letter-spacing: 0.05em;
		font-weight: 700;
		border-left: 4px solid var(--color-primary);
		padding-left: 8px;
	}

	.pill-list {
		display: flex;
		flex-wrap: wrap;
		gap: 8px;
	}

	.pill {
		background: var(--bg-canvas);
		border: 1px solid var(--border-subtle);
		padding: 6px 12px;
		font-size: 0.85rem;
		font-family: var(--font-mono);
		font-weight: 600;
		text-transform: uppercase;
	}

	.pill.price {
		background: #10b981;
		color: #ffffff;
		border-color: #059669;
	}

	.pill.outline {
		background: transparent;
		border: 1px dashed var(--border-subtle);
	}

	.placeholder {
		color: var(--text-muted);
		font-style: italic;
		margin: 0;
		font-size: 0.9rem;
	}
	
	.links-container {
		display: flex;
		align-items: center;
		gap: 24px;
		flex-wrap: wrap;
	}

	.link-btn {
		display: inline-block;
		background: var(--text-main);
		color: var(--bg-surface);
		padding: 10px 20px;
		text-decoration: none;
		font-weight: 700;
		text-transform: uppercase;
		letter-spacing: 0.05em;
		transition: transform 0.2s;
		box-shadow: 4px 4px 0 var(--color-primary);
	}

	.link-btn:hover {
		transform: translate(-2px, -2px);
		box-shadow: 6px 6px 0 var(--color-primary);
	}

	.status-info {
		display: flex;
		flex-direction: column;
		gap: 4px;
	}

	.small-label {
		font-size: 0.7rem;
		color: var(--text-muted);
		text-transform: uppercase;
		letter-spacing: 0.05em;
		font-weight: 700;
	}

	.small-text {
		font-size: 0.9rem;
		font-weight: 600;
		font-family: var(--font-mono);
	}

	.status-operational { color: #10b981; }
	.status-degraded { color: #f59e0b; }
	.status-outage { color: #ef4444; }
	.status-maintenance { color: #3b82f6; }
</style>
