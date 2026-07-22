\# RFC-001: Search Lifecycle



Status: Proposed



\---



\# Purpose



This RFC defines the complete user journey from entering Appenheimer to leaving it.



It intentionally avoids implementation details.



\---



\# Goals



\- Help users discover software in 30–60 seconds.

\- Keep interactions minimal.

\- Reveal information progressively.

\- Make every search deterministic.

\- Respect the user's attention.



\---



\# User Journey



\## 1. Landing



The user arrives at Appenheimer.



The interface is intentionally minimal.



Visible:



\- Search bar

\- Search suggestions area

\- Nothing else



No application cards are displayed.



\---



\## 2. Typing



As the user types:



\- The search engine interprets intent.

\- Related keywords appear around the interface using subtle animations.

\- Application cards do not appear.



This gives users confidence that their intent is being understood.



\---



\## 3. Suggestions



Suggestions may be:



\- Categories

\- Related searches

\- Similar phrases

\- Common searches



Users may:



\- Continue typing

\- Press Enter

\- Click a suggestion



\---



\## 4. Search



Search begins when:



\- Enter is pressed

\- A suggestion is clicked



Never before.



\---



\## 5. Search Processing



The backend performs deterministic search.



Results are ranked.



The layout seed is generated.



\---



\## 6. Results



Applications appear on the floating canvas.



Requirements:



\- No overlap

\- Equal spacing

\- Stable positions

\- Smooth entrance animation



The same query always produces the same layout.



\---



\## 7. Exploration



Users freely explore.



Hovering reveals quick information.



Clicking reveals complete details.



No navigation away from the results page occurs.



\---



\## 8. Official Website



Selecting "Visit Website":



\- Opens the official website.

\- Opens in a new tab.

\- Appenheimer remains open.



The interface indicates which application is currently being explored.



\---



\## 9. Return



Returning to Appenheimer restores the previous state.



Users continue where they left off.



No unnecessary reloads.



\---



\## 10. Session End



When the user leaves:



If session persistence is enabled:



\- Save the current state.



On the next visit:



Offer:



\- Continue previous session

\- Start new session



\---



\# Principles



Throughout the journey:



\- Never overwhelm users.

\- Never interrupt users.

\- Never require unnecessary clicks.

\- Never hide useful information.

\- Never surprise users.



\---



\# Out of Scope



This RFC does not define:



\- Ranking algorithm

\- Search indexing

\- Floating canvas algorithm

\- Hover UI

\- Application details

\- Filters

\- Backend architecture



Each has its own RFC.



\---



\# Acceptance Criteria



The lifecycle is complete when:



\- Users can discover software naturally.

\- Every interaction feels intentional.

\- Search behaves consistently.

\- Returning users can continue seamlessly.

\- Every subsequent RFC fits into this lifecycle without changing it.

