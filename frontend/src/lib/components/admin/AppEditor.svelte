<script lang="ts">
	import { onMount } from 'svelte';
	import ChangeSummaryModal from './ChangeSummaryModal.svelte';
	import PublishReadiness from './PublishReadiness.svelte';
	
	export let appData = {};
	export let readonly = false;
	
	let draftData = { ...appData };
	let originalDataStr = JSON.stringify(appData);
	
	let isDirty = false;
	let showSummary = false;
	
	$: isDirty = JSON.stringify(draftData) !== originalDataStr;
	
	// LocalStorage Recovery
	onMount(() => {
		if (!readonly) {
			const recovered = localStorage.getItem(`draft_${appData.id}`);
			if (recovered && recovered !== originalDataStr) {
				if (confirm("Recover unsaved changes?")) {
					draftData = JSON.parse(recovered);
				}
			}
			
			// Auto-snapshot to localStorage
			const interval = setInterval(() => {
				if (isDirty) localStorage.setItem(`draft_${appData.id}`, JSON.stringify(draftData));
			}, 5000);
			return () => clearInterval(interval);
		}
	});

	// beforeunload Guard
	const handleUnload = (e) => {
		if (isDirty && !readonly) {
			e.preventDefault();
			e.returnValue = '';
		}
	};

	async function handleSaveClick() {
		if (!isDirty) return;
		showSummary = true;
	}

	async function confirmSave() {
		// PATCH /admin/apps/{id}
		// On success:
		isDirty = false;
		originalDataStr = JSON.stringify(draftData);
		localStorage.removeItem(`draft_${appData.id}`);
		showSummary = false;
	}
</script>

<svelte:window on:beforeunload={handleUnload} />

<div class="editor-container">
	<div class="header">
		<h2>{readonly ? 'Reviewing' : 'Editing'}: {draftData.name}</h2>
		<div class="metadata">
			<span>Version: {draftData.version}</span>
			<span>Status: {draftData.visibility_status}</span>
		</div>
		{#if readonly && draftData.visibility_status === 'published'}
			<button class="btn" on:click={() => {/* POST shadow-draft */}}>Edit App</button>
		{:else if !readonly}
			<button class="btn btn-primary" disabled={!isDirty} on:click={handleSaveClick}>
				Save Draft {isDirty ? '*' : ''}
			</button>
		{/if}
	</div>

	<div class="layout">
		<div class="form-sections">
			<!-- General -->
			<section class="card">
				<h3>General</h3>
				<input type="text" bind:value={draftData.name} disabled={readonly} placeholder="App Name" />
				<textarea bind:value={draftData.description} disabled={readonly} placeholder="Description"></textarea>
			</section>
			
			<!-- Sub-resources (Taxonomies, Links, Media, Aliases) stubbed -->
		</div>

		<div class="sidebar">
			<PublishReadiness {draftData} />
			
			<div class="preview-section">
				<h3>Live Preview</h3>
				<!-- Reuse ResultCard / AppDrawer here passing draftData -->
				<p class="text-muted">Preview renders public components with local state.</p>
			</div>
		</div>
	</div>
</div>

{#if showSummary}
	<ChangeSummaryModal 
		original={JSON.parse(originalDataStr)} 
		current={draftData} 
		onConfirm={confirmSave} 
		onCancel={() => showSummary = false} 
	/>
{/if}
