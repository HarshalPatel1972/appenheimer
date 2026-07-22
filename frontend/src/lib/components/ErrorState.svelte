<script lang="ts">
	let { message = "An error occurred" } = $props<{ message?: string }>();
</script>

<div class="error-state">
	<p class="error-msg">{message}</p>
</div>

<style>
	.error-state {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100%;
		padding: 2rem;
	}
	
	.error-msg {
		color: #ef4444;
		background: rgba(239, 68, 68, 0.1);
		padding: 1rem 2rem;
		border-radius: 8px;
		border: 1px solid rgba(239, 68, 68, 0.2);
	}
</style>
