\# RFC-003: Floating Canvas



Status: Proposed



\---



\# Purpose



This RFC defines how search results are presented visually.



The floating canvas is one of Appenheimer's defining experiences. It should feel organic, memorable, and predictable while remaining easy to use.



This RFC defines behavior, not implementation.



\---



\# Goals



\- Create a memorable search experience.

\- Preserve readability.

\- Maintain deterministic layouts.

\- Encourage exploration.

\- Never sacrifice usability for aesthetics.



\---



\# Core Principles



\- Deterministic.

\- No overlap.

\- Equal visual importance.

\- Stable positions.

\- Smooth movement.

\- Responsive.



\---



\# Layout



Applications are not displayed in a traditional grid.



Instead, they occupy available space naturally.



The layout should appear random to users while being completely deterministic.



\---



\# Deterministic Placement



For the same search query:



\- Applications appear in the same locations.

\- Relative positions remain stable.

\- Users naturally develop spatial memory.



Layouts only change when the result set changes.



\---



\# Spacing



Every application should have sufficient space.



Requirements:



\- No overlapping.

\- No touching.

\- No crowded areas.

\- Balanced whitespace.



\---



\# Visual Hierarchy



No application should dominate the layout unless intentionally highlighted by future product decisions.



Every result begins with equal visual importance.



\---



\# Motion



Applications animate into position.



Animations should:



\- feel smooth

\- avoid distraction

\- preserve orientation

\- never cause layout confusion



Motion exists to explain where elements came from, not to entertain.



\---



\# Responsiveness



The canvas adapts to:



\- desktop

\- tablet

\- mobile



Without changing the deterministic nature of the layout.



\---



\# Interaction



Applications remain stationary after rendering.



Hovering or selecting an application should never rearrange unrelated applications.



The canvas should feel stable.



\---



\# Window Resize



Resizing the browser should preserve the relative structure of the layout whenever possible.



The user should not feel that the search has changed.



\---



\# Empty Space



Empty space is intentional.



It improves readability and gives every application room to breathe.



The goal is not to fill every pixel.



\---



\# Performance



Rendering should remain smooth.



Layout calculation should never noticeably delay search results.



\---



\# Out of Scope



This RFC does not define:



\- layout algorithm

\- collision detection

\- animation library

\- rendering implementation

\- hover behavior

\- application details



These belong in separate RFCs.



\---



\# Acceptance Criteria



The floating canvas is successful when:



\- it feels organic

\- it remains deterministic

\- users can easily explore applications

\- layouts are readable

\- motion enhances understanding

\- usability always takes priority over appearance

