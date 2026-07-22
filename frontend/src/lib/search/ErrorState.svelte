<script lang="ts">
	let { message }: { message: string | null } = $props();
</script>

<div class="state-container">
	<p class="error">{message || 'Something went wrong'}</p>
</div>

<style>
	.state-container {
		color: var(--text-main);
	}
	.error {
		color: var(--color-error);
	}
</style>
