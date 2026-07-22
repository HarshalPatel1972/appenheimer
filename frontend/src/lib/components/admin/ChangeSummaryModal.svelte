<script lang="ts">
	export let original;
	export let current;
	export let onConfirm;
	export let onCancel;
	
	// Complex diff logic (stubbed for now)
	let diffs = [
		"Name updated",
		"Ranking Weight: 85 -> 90"
	];
</script>

<div class="modal">
	<div class="modal-content">
		<h2>Review Changes</h2>
		<ul class="diff-list">
			{#each diffs as diff}
				<li>{diff}</li>
			{/each}
		</ul>
		<div class="actions">
			<button on:click={onCancel}>Cancel</button>
			<button class="btn-primary" on:click={onConfirm}>Confirm & Save</button>
		</div>
	</div>
</div>
