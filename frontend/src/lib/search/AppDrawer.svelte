<script lang="ts">
	import { detailsStore } from '$lib/state/details.svelte';
	import { fly, fade } from 'svelte/transition';
	
	let app = $derived(detailsStore.appDetails);
	let isLoading = $derived(detailsStore.isLoading);
</script>

<div class="drawer" transition:fly={{ x: 400, duration: 300, opacity: 1 }}>
	{#if isLoading || !app}
		<div class="loading-state" in:fade={{ duration: 150 }}>
			<p>Loading details...</p>
		</div>
	{:else}
		<!-- We key the content on the app ID so it crossfades when the ID changes while the drawer stays open -->
		{#key app.id}
			<div class="drawer-content" in:fade={{ duration: 200, delay: 50 }}>
				<header>
					<div class="icon"></div>
					<div class="title-block">
						<h2>{app.name}</h2>
						<span class="developer">{app.developer}</span>
					</div>
				</header>
				
				<div class="tags">
					{#if app.categories}
						{#each app.categories as cat}
							<span class="tag">{cat}</span>
						{/each}
					{/if}
				</div>

				<p class="description">{app.description}</p>
				
				<div class="metadata">
					{#if app.pricing && app.pricing.length > 0}
						<div class="meta-section">
							<h4>Pricing</h4>
							<div class="meta-pills">
								{#each app.pricing as price}
									<span class="pill">{price}</span>
								{/each}
							</div>
						</div>
					{/if}
					{#if app.platforms && app.platforms.length > 0}
						<div class="meta-section">
							<h4>Platforms</h4>
							<div class="meta-pills">
								{#each app.platforms as platform}
									<span class="pill">{platform}</span>
								{/each}
							</div>
						</div>
					{/if}
				</div>

				{#if app.screenshots && app.screenshots.length > 0}
					<div class="screenshots">
						{#each app.screenshots as shot}
							<!-- svelte-ignore a11y_missing_attribute -->
							<img src={shot} loading="lazy" />
						{/each}
					</div>
				{/if}

				<div class="actions">
					<a href={app.websiteUrl} target="_blank" class="primary-btn">Visit Website</a>
				</div>
			</div>
		{/key}
	{/if}
</div>

<style>
	.drawer {
		position: fixed;
		top: 0;
		right: 0;
		width: 400px;
		height: 100vh;
		background: var(--bg-surface);
		border-left: 1px solid var(--border-subtle);
		z-index: 1000;
		box-shadow: -10px 0 40px rgba(0,0,0,0.5);
		display: flex;
		flex-direction: column;
		overflow-y: auto;
		overflow-x: hidden;
	}
	
	@media (max-width: 768px) {
		.drawer {
			width: 100vw;
			border-left: none;
		}
	}
	
	.loading-state {
		display: flex;
		align-items: center;
		justify-content: center;
		height: 100%;
		color: var(--text-muted);
	}

	.drawer-content {
		padding: 32px;
		display: flex;
		flex-direction: column;
		gap: 24px;
	}

	header {
		display: flex;
		gap: 16px;
		align-items: center;
	}

	.icon {
		width: 64px;
		height: 64px;
		border-radius: 16px;
		background: var(--bg-primary);
		flex-shrink: 0;
	}

	.title-block h2 {
		margin: 0 0 4px 0;
		font-size: 1.5rem;
		color: var(--text-main);
	}

	.developer {
		color: var(--text-muted);
		font-size: 0.9rem;
	}

	.tags {
		display: flex;
		gap: 8px;
		flex-wrap: wrap;
	}

	.tag {
		padding: 4px 10px;
		border-radius: 6px;
		background: var(--bg-primary);
		color: var(--text-muted);
		font-size: 0.8rem;
		border: 1px solid var(--border-subtle);
		text-transform: capitalize;
	}

	.description {
		margin: 0;
		line-height: 1.6;
		color: var(--text-muted);
		font-size: 0.95rem;
	}

	.metadata {
		display: flex;
		gap: 32px;
	}

	.meta-section h4 {
		margin: 0 0 8px 0;
		font-size: 0.85rem;
		text-transform: uppercase;
		letter-spacing: 0.5px;
		color: var(--text-muted);
	}

	.meta-pills {
		display: flex;
		gap: 6px;
		flex-wrap: wrap;
	}

	.pill {
		padding: 2px 8px;
		border-radius: 4px;
		background: rgba(255,255,255,0.05);
		color: var(--text-main);
		font-size: 0.85rem;
	}

	.screenshots {
		display: flex;
		flex-direction: column;
		gap: 16px;
	}

	.screenshots img {
		width: 100%;
		border-radius: 8px;
		background: var(--bg-primary);
		min-height: 200px;
		object-fit: cover;
		border: 1px solid var(--border-subtle);
	}

	.actions {
		margin-top: 16px;
	}

	.primary-btn {
		display: block;
		text-align: center;
		background: var(--color-primary);
		color: white;
		text-decoration: none;
		padding: 12px 24px;
		border-radius: 8px;
		font-weight: 600;
		transition: background 0.2s ease;
	}

	.primary-btn:hover {
		background: #3b82f6;
	}
</style>
