<script lang="ts">
	import '../app.css';
	let { children } = $props();
</script>

<!-- The floating canvas and main content will be rendered here -->
{@render children()}
