<script lang="ts">
	import type { AppResult, AppDetails } from '$lib/types/search';
	import { getIconUrl } from '$lib/utils/icons';
	
	let { app, details }: { app: AppResult, details: AppDetails | null } = $props();
	let iconError = $state(false);
</script>

<div class="app-details-radial">
	<div class="center-piece">
		<div class="logo-container">
			{#if !iconError}
				<img src={getIconUrl(app.icon, app.name)} alt={app.name} class="logo-img" onerror={() => iconError = true} />
			{:else}
				<div class="logo-fallback">{app.name.charAt(0)}</div>
			{/if}
		</div>
		<h2>{app.name}</h2>
		<p class="developer">{details?.developer || 'Unknown Developer'}</p>
		<p class="description">{details?.longDescription || app.description}</p>
	</div>
	
	<!-- Top Left: Platforms -->
	<div class="quadrant top-left">
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

	<!-- Top Right: Pricing -->
	<div class="quadrant top-right">
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

	<!-- Bottom Left: Categories & Tags -->
	<div class="quadrant bottom-left">
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

	<!-- Bottom Right: Links & Status -->
	<div class="quadrant bottom-right">
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

<style>
	.app-details-radial {
		width: 100%;
		height: 100%;
		display: grid;
		grid-template-columns: 1fr auto 1fr;
		grid-template-rows: 1fr auto 1fr;
		gap: 24px;
		padding: 32px;
		box-sizing: border-box;
		color: var(--text-main);
		animation: fadeIn 0.4s ease forwards 0.2s;
		opacity: 0;
	}
	
	@keyframes fadeIn {
		to { opacity: 1; }
	}

	.center-piece {
		grid-column: 2;
		grid-row: 1 / -1; /* Spans top to bottom in the center */
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		text-align: center;
		max-width: 320px;
		z-index: 2;
	}

	.logo-container {
		width: 100px;
		height: 100px;
		background: var(--color-primary);
		border-radius: 24px;
		box-shadow: 0 10px 30px rgba(59, 130, 246, 0.3);
		margin-bottom: 16px;
		display: flex;
		align-items: center;
		justify-content: center;
		overflow: hidden;
	}

	.logo-img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.logo-fallback {
		font-size: 3rem;
		font-weight: 800;
		color: #fff;
		text-transform: uppercase;
	}

	.center-piece h2 {
		margin: 0 0 4px 0;
		font-size: 1.8rem;
		font-weight: 700;
		letter-spacing: -0.02em;
	}

	.developer {
		margin: 0 0 16px 0;
		font-size: 0.95rem;
		color: var(--text-muted);
		text-transform: uppercase;
		letter-spacing: 0.05em;
	}

	.description {
		font-size: 1.05rem;
		line-height: 1.5;
		color: rgba(255, 255, 255, 0.85); /* Slightly brighter for contrast */
		margin: 0;
	}

	.quadrant {
		background: rgba(0,0,0,0.2);
		padding: 20px;
		border-radius: 16px;
		border: 1px solid var(--border-subtle);
		width: 100%;
		max-width: 260px;
		display: flex;
		flex-direction: column;
		backdrop-filter: blur(8px);
	}

	.top-left {
		grid-column: 1;
		grid-row: 1;
		align-self: flex-end;
		justify-self: flex-end;
	}
	.top-right {
		grid-column: 3;
		grid-row: 1;
		align-self: flex-end;
		justify-self: flex-start;
	}
	.bottom-left {
		grid-column: 1;
		grid-row: 3;
		align-self: flex-start;
		justify-self: flex-end;
	}
	.bottom-right {
		grid-column: 3;
		grid-row: 3;
		align-self: flex-start;
		justify-self: flex-start;
	}

	h4 {
		margin: 0 0 12px 0;
		color: var(--text-muted);
		font-size: 0.85rem;
		text-transform: uppercase;
		letter-spacing: 0.05em;
	}

	.pill-list {
		display: flex;
		flex-wrap: wrap;
		gap: 8px;
	}

	.pill {
		background: rgba(255,255,255,0.1);
		padding: 6px 12px;
		border-radius: 99px;
		font-size: 0.85rem;
	}

	.pill.price {
		background: rgba(16, 185, 129, 0.2);
		color: #34d399;
		border: 1px solid rgba(16, 185, 129, 0.3);
	}

	.pill.outline {
		background: transparent;
		border: 1px solid rgba(255,255,255,0.2);
		font-size: 0.75rem;
		padding: 4px 10px;
	}

	.placeholder {
		color: var(--text-muted);
		font-style: italic;
		margin: 0;
		font-size: 0.9rem;
	}
	
	.links-container {
		display: flex;
		flex-direction: column;
		gap: 16px;
	}

	.link-btn {
		display: inline-block;
		background: var(--text-main);
		color: var(--bg-surface);
		padding: 10px 16px;
		border-radius: 8px;
		text-decoration: none;
		font-weight: 600;
		text-align: center;
		transition: transform 0.2s;
	}

	.link-btn:hover {
		transform: scale(1.05);
	}

	.status-info {
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	.small-label {
		font-size: 0.8rem;
		color: var(--text-muted);
	}

	.small-text {
		font-size: 0.85rem;
		font-weight: 500;
	}

	.status-operational { color: #34d399; }
	.status-degraded { color: #fbbf24; }
	.status-outage { color: #f87171; }
	.status-maintenance { color: #60a5fa; }
</style>
