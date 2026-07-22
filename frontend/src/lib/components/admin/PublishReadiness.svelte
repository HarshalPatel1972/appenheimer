<script lang="ts">
	export let draftData;
	
	$: hasName = !!draftData.name;
	$: hasOrg = !!draftData.organization_id;
	$: hasDesc = draftData.description?.length > 10;
	$: ready = hasName && hasOrg && hasDesc;
</script>

<div class="card readiness">
	<h3>Publish Readiness</h3>
	<ul>
		<li>{hasName ? '✓' : '⚠'} Name</li>
		<li>{hasOrg ? '✓' : '⚠'} Organization</li>
		<li>{hasDesc ? '✓' : '⚠'} Description > 10 chars</li>
	</ul>
	<p class="status {ready ? 'ready' : 'not-ready'}">
		Ready to Publish: {ready ? 'Yes' : 'No'}
	</p>
</div>
