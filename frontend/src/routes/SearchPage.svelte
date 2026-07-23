<script lang="ts">
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { searchStore } from '$lib/state/search.svelte';
	import { detailsStore } from '$lib/state/details.svelte';
	import { fetchAppDetails } from '$lib/services/apps';
	import { liveSearch, executeSearch } from '$lib/services/search';
	import SearchCanvas from '$lib/search/SearchCanvas.svelte';
	import EmptyState from '$lib/components/EmptyState.svelte';
	import SuggestionsDropdown from '$lib/search/SuggestionsDropdown.svelte';
	import LoadingState from '$lib/components/LoadingState.svelte';
	import ErrorState from '$lib/components/ErrorState.svelte';
	import FilterBar from '$lib/search/FilterBar.svelte';
	import AppDetailsView from '$lib/search/AppDetailsView.svelte';
	import { fade } from 'svelte/transition';

	// ── Input state ─────────────────────────────────────────────────────────
	let inputQuery = $state('');
	let inputFocused = $state(false);
	let showSuggestions = $state(false);
	let inputEl: HTMLInputElement;
	// Tracks the last q value the URL effect synced FROM the URL.
	// We only update inputQuery from the URL when THIS changes — not when
	// searchStore.query changes due to live typing (which was the typing bug).
	let lastUrlQ = $state('');

	// ── Animated rotating placeholder ────────────────────────────────────────
	const placeholders = [
		'try: video editor',
		'try: free design tool',
		'try: music production',
		'try: code editor',
		'try: project management',
		'try: browser',
		'try: image editing',
		'try: open source',
		'try: 3D modelling',
		'try: screen recorder',
	];
	let placeholderIdx = $state(0);
	let placeholderVisible = $state(true);

	$effect(() => {
		if (inputFocused || inputQuery.length > 0) return;
		const id = setInterval(() => {
			// Fade out → swap text → fade in
			placeholderVisible = false;
			setTimeout(() => {
				placeholderIdx = (placeholderIdx + 1) % placeholders.length;
				placeholderVisible = true;
			}, 300);
		}, 2800);
		return () => clearInterval(id);
	});

	let currentPlaceholder = $derived(
		(inputFocused || inputQuery.length > 0) ? 'Search for software...' : placeholders[placeholderIdx]
	);

	// ── Derived booleans ─────────────────────────────────────────────────────
	// Show suggestion dropdown when focused + typing + results exist
	let canShowSuggestions = $derived(
		showSuggestions &&
		inputFocused &&
		inputQuery.trim().length > 0 &&
		(searchStore.status === 'success' || searchStore.status === 'loading') &&
		searchStore.results.length > 0
	);

	// Fade detail quadrants when user re-types while an app is active
	let isTypingWhileActive = $derived(
		!!detailsStore.activeAppId && inputFocused && inputQuery.trim().length > 0
	);

	// Whether to show the floating canvas (only after commit, not while suggesting)
	let showCanvas = $derived(
		!showSuggestions &&
		searchStore.results.length > 0 &&
		(searchStore.status === 'success' || searchStore.status === 'loading')
	);

	// ── URL → State sync ─────────────────────────────────────────────────────
	// Only overwrites inputQuery from URL during real navigations (back/forward,
	// direct links, filter clicks), NOT during live typing.
	$effect(() => {
		const q = $page.url.searchParams.get('q') || '';
		const appId = $page.url.searchParams.get('app');
		const platform = $page.url.searchParams.get('platform');
		const pricing = $page.url.searchParams.get('pricing');
		const category = $page.url.searchParams.get('category');

		const platforms = platform ? platform.split(',') : [];
		const pricings = pricing ? pricing.split(',') : [];
		const categories = category ? category.split(',') : [];

		let changed = false;

		// Only sync inputQuery ← URL when the URL's q actually changed
		// (back/forward navigation, direct link, filter click, commit).
		// This fixes the typing bug: liveSearch() updates searchStore.query but
		// NOT the URL, so q !== searchStore.query was triggering a false reset.
		if (q !== lastUrlQ) {
			lastUrlQ = q;
			inputQuery = q;
			searchStore.query = q;
			showSuggestions = false;
			changed = true;
		}

		if (JSON.stringify(searchStore.filters.platforms) !== JSON.stringify(platforms)) {
			searchStore.filters.platforms = platforms;
			changed = true;
		}
		if (JSON.stringify(searchStore.filters.pricing) !== JSON.stringify(pricings)) {
			searchStore.filters.pricing = pricings;
			changed = true;
		}
		if (JSON.stringify(searchStore.filters.categories) !== JSON.stringify(categories)) {
			searchStore.filters.categories = categories;
			changed = true;
		}

		if (!q && platforms.length === 0 && pricings.length === 0 && categories.length === 0) {
			searchStore.status = 'idle';
			searchStore.results = [];
		}

		if (changed) {
			const newOrder: {type: string, value: string}[] = [];
			const currentMap = new Map(searchStore.activeFiltersOrder.map(f => [f.type + ':' + f.value, f]));
			const addFilters = (type: string, vals: string[]) => {
				for (const v of vals) {
					const key = type + ':' + v;
					newOrder.push(currentMap.has(key) ? currentMap.get(key)! : { type, value: v });
				}
			};
			addFilters('platform', platforms);
			addFilters('pricing', pricings);
			addFilters('category', categories);
			searchStore.activeFiltersOrder = newOrder;

			if (q.trim().length > 0 || newOrder.length > 0) {
				executeSearch();
			} else {
				searchStore.status = 'idle';
				searchStore.results = [];
			}
		}

		if (appId !== detailsStore.activeAppId) {
			detailsStore.activeAppId = appId;
			if (appId) {
				fetchAppDetails(appId);
			} else {
				detailsStore.appDetails = null;
			}
		}
	});

	// Auto-close if active app is no longer in results (e.g. new search)
	$effect(() => {
		if (detailsStore.activeAppId && searchStore.status === 'success') {
			const found = searchStore.results.find(r => r.id === detailsStore.activeAppId);
			if (!found) closeDrawer();
		}
	});

	let activeApp = $derived(
		detailsStore.activeAppId
			? searchStore.results.find(r => r.id === detailsStore.activeAppId) || null
			: null
	);

	// ── Handlers ─────────────────────────────────────────────────────────────
	function handleInput(e: Event) {
		const target = e.target as HTMLInputElement;
		inputQuery = target.value;
		showSuggestions = true;
		liveSearch(inputQuery);
	}

	function handleFocus() {
		inputFocused = true;
		if (inputQuery.trim().length > 0) showSuggestions = true;
	}

	function handleBlur() {
		// Delay so a suggestion click registers before we hide the dropdown
		setTimeout(() => {
			inputFocused = false;
			showSuggestions = false;
		}, 180);
	}

	function commitSearch() {
		showSuggestions = false;
		executeSearch();
		const url = new URL($page.url);
		if (searchStore.query) {
			url.searchParams.set('q', searchStore.query);
		} else {
			url.searchParams.delete('q');
		}
		url.searchParams.delete('app');
		// Pre-set lastUrlQ so the URL effect doesn't re-execute search
		lastUrlQ = searchStore.query;
		goto(url, { keepFocus: false, noScroll: true, replaceState: true });
	}

	function handleSuggestionSelect(name: string) {
		inputQuery = name;
		showSuggestions = false;
		liveSearch(name);
		executeSearch();
		const url = new URL($page.url);
		url.searchParams.set('q', name);
		url.searchParams.delete('app');
		lastUrlQ = name;
		goto(url, { keepFocus: false, noScroll: true, replaceState: true });
	}

	function handleInputKeydown(e: KeyboardEvent) {
		if (e.key === 'Enter') {
			e.preventDefault();
			commitSearch();
		}
		if (e.key === 'Escape') {
			showSuggestions = false;
			inputEl?.blur();
		}
	}

	function handleChipClick(keyword: string) {
		inputQuery = keyword;
		showSuggestions = false;
		liveSearch(keyword);
		executeSearch();
		const url = new URL($page.url);
		url.searchParams.set('q', keyword);
		url.searchParams.delete('app');
		lastUrlQ = keyword;
		goto(url, { keepFocus: false, noScroll: true, replaceState: true });
	}

	function closeDrawer() {
		const q = $page.url.searchParams.get('q') || '';
		if (q) {
			goto(`/?q=${q}`, { keepFocus: true, noScroll: true });
		} else {
			goto(`/`, { keepFocus: true, noScroll: true });
		}
	}

	function handleGlobalKeydown(e: KeyboardEvent) {
		if (e.key === 'Escape') {
			if (showSuggestions) { showSuggestions = false; return; }
			if (detailsStore.activeAppId) { closeDrawer(); return; }
		}

		if (detailsStore.activeAppId && searchStore.status === 'success') {
			const idx = searchStore.results.findIndex(r => r.id === detailsStore.activeAppId);
			if (idx === -1) return;
			const q = $page.url.searchParams.get('q') || '';
			if ((e.key === 'ArrowRight' || e.key === 'ArrowDown') && idx < searchStore.results.length - 1) {
				e.preventDefault();
				goto(`/?q=${q}&app=${searchStore.results[idx + 1].id}`, { keepFocus: true, noScroll: true, replaceState: true });
			} else if ((e.key === 'ArrowLeft' || e.key === 'ArrowUp') && idx > 0) {
				e.preventDefault();
				goto(`/?q=${q}&app=${searchStore.results[idx - 1].id}`, { keepFocus: true, noScroll: true, replaceState: true });
			}
		}
	}
</script>

<svelte:window onkeydown={handleGlobalKeydown} />

<!-- ── NAVBAR ────────────────────────────────────────────────────── -->
<nav class="navbar">
	<span class="brand">APPENHEIMER</span>
	<a
		href="https://github.com/HarshalPatel1972/appenheimer"
		target="_blank"
		rel="noopener noreferrer"
		class="github-btn"
	>GitHub ↗</a>
</nav>

<main class="app-main">

	<!-- ── CENTRE STAGE ──────────────────────────────────────────────── -->
	<div class="center-stage" class:is-active={!!detailsStore.activeAppId}>

		<!-- Hero copy: visible only on idle, sits above the search bar -->
		{#if searchStore.status === 'idle'}
			<div class="hero-copy" transition:fade={{ duration: 220 }}>
				<h1>Find software.</h1>
				<p class="hero-sub">Not websites. Not articles. Just the right tool.</p>
			</div>
		{/if}

		<!-- Search bar wrapper.
		     When an app is active, padding-left makes room for the canvas icon
		     (which has animated to the left edge of this 600 px block). -->
		<div
			class="search-bar-wrapper"
			class:has-active-app={!!detailsStore.activeAppId}
		>
			<!-- Input row: input + arrow button share the same border box -->
			<div class="input-row" class:focused={inputFocused}>
				<input
					bind:this={inputEl}
					id="main-search"
					type="text"
					class="search-input"
					class:compact={!!detailsStore.activeAppId}
					placeholder={currentPlaceholder}
					value={inputQuery}
					oninput={handleInput}
					onkeydown={handleInputKeydown}
					onfocus={handleFocus}
					onblur={handleBlur}
					autocomplete="off"
					spellcheck="false"
					aria-label="Search for software"
					aria-haspopup="listbox"
				/>
				<!-- Arrow button: grayed when empty, primary colour when active -->
				<button
					class="search-arrow"
					class:has-query={inputQuery.trim().length > 0}
					onclick={commitSearch}
					aria-label="Search"
					tabindex={inputQuery.trim().length > 0 ? 0 : -1}
				>→</button>
			</div>

			<!-- Live suggestions dropdown -->
			{#if canShowSuggestions}
				<SuggestionsDropdown onSelect={handleSuggestionSelect} />
			{/if}

			<!-- Filter bar: hidden when app is active or suggestions showing -->
			{#if !detailsStore.activeAppId && !canShowSuggestions}
				<FilterBar />
			{/if}
		</div>
	</div>

	<!-- ── CANVAS LAYER ──────────────────────────────────────────────── -->
	<div class="content">
		{#if searchStore.status === 'idle'}
			<!-- Idle: just empty canvas — hero copy in center-stage covers this -->
		{:else if searchStore.status === 'error'}
			<ErrorState message={searchStore.error} />
		{:else if searchStore.status === 'loading' && searchStore.results.length === 0}
			<LoadingState />
		{:else if searchStore.status === 'success' && searchStore.results.length === 0 && !showSuggestions}
			<EmptyState onChipClick={handleChipClick} />
		{:else if showCanvas}
			<SearchCanvas />
		{/if}
	</div>

	<!-- ── 4-CORNER DETAIL PANELS ────────────────────────────────────── -->
	<!-- Fade away when user re-types while an app is active -->
	{#if detailsStore.activeAppId && activeApp && !isTypingWhileActive}
		<div class="details-layer" transition:fade={{ duration: 200 }}>
			<AppDetailsView app={activeApp} details={detailsStore.appDetails} />
		</div>
	{/if}

</main>

<style>
	/* ── Navbar ─────────────────────────────────────────────────────── */
	.navbar {
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		height: 48px;
		display: flex;
		align-items: center;
		justify-content: space-between;
		padding: 0 28px;
		z-index: 500;
		background: var(--bg-primary);
		border-bottom: 1.5px solid rgba(17,17,17,0.15);
		box-sizing: border-box;
	}

	.brand {
		font-family: var(--font-mono);
		font-weight: 800;
		font-size: 0.78rem;
		letter-spacing: 0.12em;
		color: var(--text-main);
		text-decoration: none;
		text-transform: uppercase;
	}

	.github-btn {
		font-family: var(--font-mono);
		font-size: 0.72rem;
		font-weight: 700;
		color: var(--text-muted);
		text-decoration: none;
		letter-spacing: 0.05em;
		text-transform: uppercase;
		border: 1.5px solid rgba(17,17,17,0.3);
		padding: 5px 11px;
		transition: all 0.18s;
	}

	.github-btn:hover {
		background: var(--text-main);
		color: var(--bg-surface);
		border-color: var(--text-main);
	}

	.app-main {
		display: flex;
		flex-direction: column;
		height: 100vh;
		width: 100vw;
		padding-top: 48px; /* clear the navbar */
		box-sizing: border-box;
		background: var(--bg-canvas);
		overflow: hidden;
		position: relative;
	}

	/* ── Centre stage ───────────────────────────────────────────────── */
	.center-stage {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		z-index: 200;
		display: flex;
		flex-direction: column;
		align-items: center;
		width: 100%;
		max-width: 600px;
		padding: 0 16px;
		box-sizing: border-box;
		gap: 0;
	}

	/* ── Hero copy ──────────────────────────────────────────────────── */
	.hero-copy {
		text-align: center;
		margin-bottom: 20px;
		width: 100%;
	}

	.hero-copy h1 {
		margin: 0 0 8px 0;
		font-size: clamp(2.2rem, 5vw, 3.2rem);
		font-weight: 800;
		letter-spacing: -0.03em;
		color: var(--text-main);
		line-height: 1;
	}

	.hero-sub {
		margin: 0;
		font-size: 0.88rem;
		color: var(--text-muted);
		font-family: var(--font-mono);
		font-weight: 500;
		letter-spacing: 0.01em;
	}

	/* ── Search bar wrapper ─────────────────────────────────────────── */
	.search-bar-wrapper {
		display: flex;
		flex-direction: column;
		width: 100%;
		position: relative;
		transition: padding-left 0.52s cubic-bezier(0.22, 1, 0.36, 1);
	}

	/* When an app is active the canvas icon sits in the left 96px,
	   so we indent the search bar to clear that space. */
	.search-bar-wrapper.has-active-app {
		padding-left: 96px; /* 80px icon + 16px gap */
	}

	/* ── Input row (input + arrow share one border) ──────────────────── */
	.input-row {
		display: flex;
		align-items: stretch;
		width: 100%;
		box-shadow: 6px 6px 0 rgba(0,0,0,1);
		transition: box-shadow 0.25s ease;
	}

	.input-row.focused {
		box-shadow: 6px 6px 0 var(--color-primary);
	}

	.search-input {
		flex: 1;
		min-width: 0;
		padding: 16px 20px;
		border: 3px solid var(--border-subtle);
		border-right: none;
		background: var(--bg-surface);
		color: var(--text-main);
		font-size: 1.15rem;
		font-weight: 700;
		font-family: var(--font-sans);
		outline: none;
		border-radius: 0;
		box-sizing: border-box;
		transition: border-color 0.2s, font-size 0.35s, padding 0.35s;
	}

	.search-input::placeholder {
		color: var(--text-muted);
		font-weight: 500;
		transition: opacity 0.3s;
	}

	.input-row.focused .search-input {
		border-color: var(--color-primary);
	}

	.search-input.compact {
		padding: 11px 16px;
		font-size: 0.95rem;
	}

	/* Arrow button */
	.search-arrow {
		flex-shrink: 0;
		padding: 0 20px;
		background: var(--bg-surface);
		border: 3px solid var(--border-subtle);
		border-left: 1px solid rgba(17,17,17,0.12);
		color: var(--text-muted);
		font-size: 1.3rem;
		font-weight: 700;
		cursor: default;
		transition: color 0.2s, background 0.2s, border-color 0.2s;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.search-arrow.has-query {
		cursor: pointer;
		color: var(--color-primary);
		border-color: var(--border-subtle);
	}

	.input-row.focused .search-arrow {
		border-color: var(--color-primary);
		border-left-color: rgba(217,56,30,0.2);
	}

	.search-arrow.has-query:hover {
		background: var(--color-primary);
		color: #fff;
	}

	/* ── Content layer ──────────────────────────────────────────────── */
	.content {
		flex: 1;
		width: 100%;
		height: 100%;
		position: relative;
	}

	/* ── Details overlay wrapper ─────────────────────────────────────── */
	.details-layer {
		position: absolute;
		inset: 0;
		pointer-events: none;
		z-index: 150;
	}

	/* ── Responsive ─────────────────────────────────────────────────── */
	@media (max-width: 640px) {
		.center-stage {
			max-width: 100%;
			padding: 0 12px;
		}

		.search-bar-wrapper.has-active-app {
			padding-left: 72px; /* smaller icon on mobile */
		}

		.hero-copy h1 {
			font-size: 1.6rem;
		}
	}
</style>
