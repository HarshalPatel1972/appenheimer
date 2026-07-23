# Appenheimer 🪐

A deterministic, spatial web discovery engine. 

Appenheimer is built on a simple philosophy: **Help users discover software and websites instantly. Don't become a destination that delays them.** It uses a highly interactive, lightning-fast search experience to get users to the tools they need in under a minute.

No LLMs. No generative AI fluff. Just deterministic, reliable, and fast search.

---

## ✨ Core Features

*   **Universal Intent Search:** One search bar, zero modes. Type app names, tasks, categories, or features (e.g., "free markdown notes for students") and the engine figures it out.
*   **Spatial Discovery Canvas:** Search results don't render in boring lists. Apps spawn in a dynamic radial constellation, leveraging muscle memory and object permanence.
*   **Radial App Details:** Clicking an app doesn't open a traditional modal. The app's identity anchors to the center of the screen, while metadata (Platforms, Pricing, Categories, Links) radially surrounds it in consistent spatial quadrants.
*   **Live Context-Aware Filters:** Discover first, filter later. Filters dynamically generate based on search context (e.g., "Templates" for design apps) and apply instantly.
*   **Instant Search Chips:** Search queries automatically break down into removable UI chips. Removing a chip instantly updates the results canvas without reloading.

---

## 🛠️ Technology Stack

*   **Frontend:** [Svelte 5](https://svelte.dev/) utilizing runes (`$state`, `$derived`, `$effect`) for fine-grained reactivity.
*   **Styling:** Vanilla CSS focusing on glassmorphism, fluid animations, and a premium dark-mode aesthetic.
*   **Backend:** (In Progress) Go-based backend services for search indexing and telemetry.

---

## 🚀 Getting Started

The project is currently in active development. To run the frontend locally:

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the development server:
   ```bash
   npm run dev
   ```
4. Open `http://localhost:5173` in your browser.

---

## 🗺️ Roadmap & Features

The project is driven by a strict, community-focused feature roadmap.
You can view the detailed status of all Confirmed, Parked, and Killed features in the root `feature_index.html` file.

---

## 📁 Repository Structure

*   `frontend/` - SvelteKit frontend application.
*   `backend/` - Go-based backend services (WIP).
*   `docs/` - Core philosophy, architecture, and product documentation.
*   `assets/` - Static assets and media.
*   `infrastructure/` - Deployment and infrastructure configuration.
*   `scripts/` - Utility and automation scripts.
*   `specs/` - Detailed technical specifications for features.

---

*“Updates are earned, not pushed.”*
