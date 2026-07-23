<script lang="ts">
	import { cubicIn } from 'svelte/easing';
	
	const sampleKeywords = ['design', 'video editor', 'windows', 'free', 'prototyping', 'IDE', 'music production'];
	
	function attract(node: HTMLElement, { duration = 400 }) {
		return {
			duration,
			easing: cubicIn,
			css: (t: number) => `
				transform: translate(0, 0) scale(${t});
				opacity: ${t};
			`
		};
	}
</script>

<div class="typing-state">
	<div class="orbit-center">
		{#each sampleKeywords as kw, i}
			<div 
				class="orbit-container" 
				style="animation-delay: {-i * 1.5}s;"
				out:attract
			>
				<span class="kw" style="animation-delay: {-i * 0.5}s;">{kw}</span>
			</div>
		{/each}
	</div>
</div>

<style>
	.typing-state {
		position: absolute;
		top: 50px;
		left: 50%;
		transform: translateX(-50%);
		width: 100%;
		height: 100%;
		pointer-events: none;
		z-index: 50;
	}

	.orbit-center {
		position: absolute;
		top: 50px; /* Adjust relative to search bar */
		left: 50%;
		width: 0;
		height: 0;
	}

	.orbit-container {
		position: absolute;
		top: 0;
		left: 0;
		/* Animate the rotation to orbit around the center */
		animation: orbit 20s linear infinite;
	}

	.kw {
		display: inline-block;
		background: rgba(255,255,255,0.05);
		border: 1px solid rgba(255,255,255,0.1);
		padding: 6px 14px;
		border-radius: 99px;
		color: var(--text-muted);
		font-size: 0.9rem;
		white-space: nowrap;
		/* Translate outwards to create the radius, then counter-rotate to stay upright */
		animation: counter-orbit 20s linear infinite;
		transform-origin: center;
	}

	@keyframes orbit {
		0% { transform: rotate(0deg) translateX(300px); }
		100% { transform: rotate(360deg) translateX(300px); }
	}

	@keyframes counter-orbit {
		0% { transform: rotate(0deg); }
		100% { transform: rotate(-360deg); }
	}
</style>
