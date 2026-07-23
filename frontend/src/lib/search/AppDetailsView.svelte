<script lang="ts">
	import type { AppResult, AppDetails } from '$lib/types/search';
	import { getIconUrl } from '$lib/utils/icons';
	
	let { app, details }: { app: AppResult, details: AppDetails | null } = $props();
	let iconError = $state(false);
</script>

<div class="scattered-details-container">
	<!-- Top Left: Overview -->
	<div class="quadrant top-left">
		<span class="quad-tag">01 // OVERVIEW</span>
		<div class="developer">{details?.developer || 'Official App'}</div>
		<p class="description">{details?.longDescription || app.description}</p>
	</div>

	<!-- Top Right: Platforms & Pricing -->
	<div class="quadrant top-right">
		<span class="quad-tag">02 // SPECIFICATIONS</span>
		<div class="meta-group">
			<h4>Platforms</h4>
			<div class="pill-list">
				{#if details?.platforms}
					{#each details.platforms as p}
						<span class="pill">{p}</span>
					{/each}
				{:else}
					<span class="placeholder">Web, Desktop</span>
				{/if}
			</div>
		</div>

		<div class="meta-group">
			<h4>Pricing</h4>
			<div class="pill-list">
				{#if details?.pricing}
					{#each details.pricing as p}
						<span class="pill price">{p}</span>
					{/each}
				{:else}
					<span class="placeholder">Free / Subscription</span>
				{/if}
			</div>
		</div>
	</div>

	<!-- Bottom Left: Categories & Tags -->
	<div class="quadrant bottom-left">
		<span class="quad-tag">03 // TAXONOMY</span>
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
				<span class="placeholder">Software, Productivity</span>
			{/if}
		</div>
	</div>

	<!-- Bottom Right: Access & System Status -->
	<div class="quadrant bottom-right">
		<span class="quad-tag">04 // SYSTEM STATUS</span>
		<div class="access-box">
			{#if details?.websiteUrl}
				<a href={details.websiteUrl} target="_blank" rel="noopener noreferrer" class="link-btn">
					Launch App ↗
				</a>
			{/if}
			<div class="status-row">
				<span class="status-label">Status</span>
				<span class="status-value status-{app.status?.toLowerCase() || 'operational'}">● {app.status || 'Operational'}</span>
			</div>
		</div>
	</div>
</div>

<style>
	.scattered-details-container {
		position: absolute;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		pointer-events: none;
		z-index: 150;
	}

	.quadrant {
		position: absolute;
		background: var(--bg-surface);
		border: 2px solid var(--border-subtle);
		box-shadow: 6px 6px 0 rgba(0,0,0,1);
		padding: 20px 24px;
		box-sizing: border-box;
		color: var(--text-main);
		pointer-events: auto;
		display: flex;
		flex-direction: column;
		gap: 12px;
		animation: fadeIn 0.4s ease forwards;
	}

	@keyframes fadeIn {
		from { opacity: 0; transform: scale(0.95) translateY(6px); }
		to   { opacity: 1; transform: scale(1) translateY(0); }
	}

	.quad-tag {
		font-size: 0.65rem;
		font-weight: 800;
		font-family: var(--font-mono);
		color: var(--color-primary);
		letter-spacing: 0.08em;
	}

	.top-left {
		top: 72px;
		left: 24px;
		max-width: 360px;
		animation-delay: 0.05s;
		border-top: 4px solid var(--color-primary);
	}

	.top-right {
		top: 72px;
		right: 24px;
		max-width: 300px;
		animation-delay: 0.15s;
		border-top: 4px solid var(--border-subtle);
	}

	.bottom-left {
		bottom: 24px;
		left: 24px;
		max-width: 360px;
		animation-delay: 0.1s;
		border-top: 4px solid var(--border-subtle);
	}

	.bottom-right {
		bottom: 24px;
		right: 24px;
		max-width: 280px;
		animation-delay: 0.2s;
		border-top: 4px solid var(--color-primary);
	}

	.developer {
		font-size: 0.75rem;
		font-weight: 700;
		color: var(--text-muted);
		text-transform: uppercase;
		letter-spacing: 0.05em;
	}

	.description {
		font-size: 0.95rem;
		line-height: 1.5;
		margin: 0;
		color: var(--text-main);
		font-family: var(--font-mono);
	}

	.meta-group {
		display: flex;
		flex-direction: column;
		gap: 6px;
	}

	h4 {
		margin: 0;
		font-size: 0.75rem;
		text-transform: uppercase;
		font-weight: 700;
		color: var(--text-muted);
		letter-spacing: 0.05em;
	}

	.pill-list {
		display: flex;
		flex-wrap: wrap;
		gap: 6px;
	}

	.pill {
		background: var(--bg-canvas);
		border: 1px solid var(--border-subtle);
		padding: 4px 10px;
		font-size: 0.75rem;
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
		font-size: 0.8rem;
		color: var(--text-muted);
		font-style: italic;
	}

	.access-box {
		display: flex;
		flex-direction: column;
		gap: 16px;
	}

	.link-btn {
		display: inline-block;
		background: var(--text-main);
		color: var(--bg-surface);
		padding: 10px 18px;
		text-decoration: none;
		font-weight: 700;
		font-size: 0.85rem;
		text-transform: uppercase;
		letter-spacing: 0.05em;
		text-align: center;
		transition: transform 0.2s, box-shadow 0.2s;
		box-shadow: 4px 4px 0 var(--color-primary);
	}

	.link-btn:hover {
		transform: translate(-2px, -2px);
		box-shadow: 6px 6px 0 var(--color-primary);
	}

	.status-row {
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	.status-label {
		font-size: 0.75rem;
		font-weight: 700;
		color: var(--text-muted);
		text-transform: uppercase;
	}

	.status-value {
		font-size: 0.85rem;
		font-weight: 700;
		font-family: var(--font-mono);
	}

	.status-operational { color: #10b981; }
	.status-degraded { color: #f59e0b; }
	.status-outage { color: #ef4444; }
	.status-maintenance { color: #3b82f6; }

	@media (max-width: 900px) {
		.scattered-details-container {
			position: fixed;
			top: 80px;
			left: 0;
			width: 100vw;
			height: calc(100vh - 80px);
			overflow-y: auto;
			display: flex;
			flex-direction: column;
			align-items: center;
			gap: 16px;
			padding: 16px;
			box-sizing: border-box;
			z-index: 100;
		}

		.quadrant {
			position: relative;
			top: auto !important;
			left: auto !important;
			right: auto !important;
			bottom: auto !important;
			width: 100%;
			max-width: 500px !important;
		}
	}
</style>
