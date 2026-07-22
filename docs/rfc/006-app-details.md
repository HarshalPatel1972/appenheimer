# RFC 006 - App Details Panel

## 1. Context
After a user discovers an app via the Search Canvas and previews it via the Hover Experience, the next step in the discovery funnel is the **App Details Panel**. This is the final decision-making surface before the user commits to visiting the official software website.

## 2. Goals
- Provide a rich, focused view containing deep metadata (screenshots, full descriptions, pricing models, platform support).
- Maintain the spatial context of the Search Canvas underneath to allow rapid exploration.
- Ensure the state is shareable (SEO friendly / direct linking).
- Keep the user experience native, fast, and un-intrusive.

## 3. Approved Decisions
- **Presentation**: On desktop, a right-side drawer. On mobile, a full-screen takeover.
- **Routing**: Utilize URL query parameters (`/?q=design&app=figma`) rather than strict routes (`/app/[id]`). This asserts that search is the primary context, allows seamless shareability, and makes "Back" functionality native.
- **Backend**: Implement `GET /api/v1/apps/{id}` to fetch heavy data on-demand, keeping the initial search payload incredibly light.
- **Media Loading**: Prioritize text rendering. Pre-fetch favicons and explicitly lazy-load screenshots to prevent blocking the drawer's initial render.
- **Drawer Persistence**: If the drawer is open and the user clicks a different app on the canvas, the drawer stays open and smoothly transitions its content. It does not close and re-open.
- **UX Controls**: Pressing `Esc` closes the drawer. Triggering the Browser Back button closes the drawer while preserving the search page.

## 4. Technical Implications
### Backend Changes
- **New Endpoint**: `GET /api/v1/apps/{id}`.
- **Domain Model**: A new `AppDetails` struct adding `Screenshots`, `Pricing`, `Platforms`, `WebsiteURL`, and `Developer` over the base model.

### Frontend Changes
- **SvelteKit Routing**: Deeply bind the `$page.url.searchParams` to the UI state.
- **Transitions**: Implement dual-layer transitions: a slide transition for the drawer itself, and a fade transition for the data inside it to support seamless "app switching" while the drawer is open.
