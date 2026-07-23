<script lang="ts">
	import type { AppResult, AppDetails } from '$lib/types/search';
	
	let { app, details }: { app: AppResult, details: AppDetails | null } = $props();
</script>

<div class="app-details-view">
	<div class="header">
		<div class="logo">
			<!-- Large Logo Placeholder -->
		</div>
	</div>
	
	<div class="grid-layout">
		<!-- Top Left: Platforms -->
		<div class="section platforms">
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
		<div class="section pricing">
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

		<!-- Bottom Left: Categories -->
		<div class="section categories">
			<h4>Categories</h4>
			{#if details?.categories}
				<div class="pill-list">
					{#each details.categories as c}
						<span class="pill">{c}</span>
					{/each}
				</div>
			{:else}
				<p class="placeholder">Loading...</p>
			{/if}
		</div>

		<!-- Bottom Right: Website Links -->
		<div class="section links">
			<h4>Links</h4>
			{#if details?.websiteUrl}
				<a href={details.websiteUrl} target="_blank" rel="noopener noreferrer" class="link-btn">
					Visit Website ↗
				</a>
			{:else}
				<p class="placeholder">Loading...</p>
			{/if}
		</div>
	</div>

	<!-- Bottom Center: Tags & Details -->
	<div class="footer-sections">
		<div class="section description">
			<p>{details?.longDescription || app.description}</p>
		</div>

		<div class="section extra">
			<div class="col">
				<h4>Updates</h4>
				<p class="small-text">{details?.lastUpdate ? new Date(details.lastUpdate).toLocaleDateString() : 'Unknown'}</p>
			</div>
			<div class="col">
				<h4>Developer</h4>
				<p class="small-text">{details?.developer || 'Unknown'}</p>
			</div>
			<div class="col">
				<h4>Tags</h4>
				<div class="pill-list small">
					{#each details?.tags || [] as t}
						<span class="pill outline">{t}</span>
					{/each}
				</div>
			</div>
		</div>
	</div>
</div>

<style>
	.app-details-view {
		width: 100%;
		height: 100%;
		display: flex;
		flex-direction: column;
		padding: 40px;
		box-sizing: border-box;
		overflow-y: auto;
		color: var(--text-main);
		animation: fadeIn 0.4s ease forwards 0.2s;
		opacity: 0;
	}
	
	@keyframes fadeIn {
		to { opacity: 1; }
	}

	.header {
		display: flex;
		justify-content: center;
		margin-bottom: 40px;
	}

	.logo {
		width: 120px;
		height: 120px;
		background: var(--color-primary);
		border-radius: 24px;
		box-shadow: 0 10px 30px rgba(59, 130, 246, 0.3);
	}

	.grid-layout {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 32px;
		margin-bottom: 40px;
	}

	.section {
		background: rgba(0,0,0,0.2);
		padding: 20px;
		border-radius: 16px;
		border: 1px solid var(--border-subtle);
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
		font-size: 0.9rem;
	}

	.pill.price {
		background: rgba(16, 185, 129, 0.2);
		color: #34d399;
		border: 1px solid rgba(16, 185, 129, 0.3);
	}

	.pill.outline {
		background: transparent;
		border: 1px solid rgba(255,255,255,0.2);
		font-size: 0.8rem;
		padding: 4px 10px;
	}

	.placeholder {
		color: var(--text-muted);
		font-style: italic;
		margin: 0;
		font-size: 0.9rem;
	}

	.link-btn {
		display: inline-block;
		background: var(--text-main);
		color: var(--bg-surface);
		padding: 10px 20px;
		border-radius: 8px;
		text-decoration: none;
		font-weight: 600;
		transition: transform 0.2s;
	}

	.link-btn:hover {
		transform: scale(1.05);
	}

	.footer-sections {
		display: flex;
		flex-direction: column;
		gap: 24px;
	}

	.description p {
		font-size: 1.1rem;
		line-height: 1.6;
		margin: 0;
	}

	.extra {
		display: flex;
		gap: 40px;
	}

	.col {
		flex: 1;
	}

	.small-text {
		margin: 0;
		font-size: 0.9rem;
	}
</style>
