# RFC 005 - Hover Experience

## 1. Context
Following the successful implementation of the Search Canvas MVP, the next natural user interaction is exploration. The hover state acts as the first layer of progressive disclosure, allowing users to evaluate a search result without committing to a full page navigation or click. 

## 2. Goals
- **Contextual Clarity**: Display quick, relevant information immediately on hover.
- **Status Awareness**: Prevent unnecessary clicks by visually indicating if an application is currently experiencing downtime.
- **Frictionless Exploration**: Enhance discovery without cluttering the initial deterministic canvas.
- **Performance**: Maintain the 60fps organic feel of the Search Canvas during interactions.

## 3. Approved Decisions
- **Hover Delay**: `150ms` delay. Fast enough to feel responsive, slow enough to avoid accidental triggers when moving across the canvas.
- **Data Loading**: Include hover data (tags, status) in the initial search payload.
- **Presentation Strategy**: 
  - The hovered card scales slightly (`≈1.05x`).
  - Surrounding cards are blurred/dimmed.
  - A **floating information panel** spawns beside the card, positioned intelligently.
- **Keyboard Parity**: `focus` triggers the exact same hover logic as a mouse `hover`.
- **Uniqueness Constraint**: **Only one hover panel may exist at any time.** 
- **Strict Layout Stability Constraint**: **Hover must never trigger layout recalculation.** Only visual state changes are allowed. The canvas must remain perfectly stable.

## 4. Technical Implications
### Backend Changes
- The `App` domain model needs a new `Status` field (Enum: `Operational, Degraded, Maintenance, Outage, Unknown`).
- The in-memory dataset must be updated to include mocked statuses for testing.
- The `SearchResponse` schema must be updated to return this new hover data.

### Frontend Changes
- **Global Hover State**: A centralized state rune to guarantee the "only one hover panel" rule. It will store the full `hoveredApp` rather than just the ID.
- **Intelligent Positioning**: Logic to calculate the floating panel's coordinates relative to the hovered card's bounding box and the viewport boundaries.
