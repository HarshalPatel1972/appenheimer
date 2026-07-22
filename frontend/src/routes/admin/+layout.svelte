<script lang="ts">
	let { children } = $props();
</script>

<div class="admin-layout">
	<aside class="sidebar">
		<h2>Admin</h2>
		<nav>
			<a href="/admin">Dashboard</a>
			<a href="/admin/queue">Review Queue</a>
			<a href="/admin/imports">Imports</a>
		</nav>
	</aside>
	<main class="content">
		{@render children()}
	</main>
</div>

<style>
	.admin-layout {
		display: flex;
		height: 100vh;
		background: #f8f9fa;
		color: #1a1a1a;
	}
	.sidebar {
		width: 250px;
		background: #ffffff;
		border-right: 1px solid #eaeaea;
		padding: 2rem;
	}
	.sidebar nav {
		display: flex;
		flex-direction: column;
		gap: 1rem;
		margin-top: 2rem;
	}
	.content {
		flex: 1;
		padding: 3rem;
		overflow-y: auto;
	}
</style>
