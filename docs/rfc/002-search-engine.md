\# RFC-002: Search Engine



Status: Proposed



\---



\# Purpose



This RFC defines how Appenheimer's search engine behaves.



It focuses on user experience and search behavior, not implementation.



\---



\# Goals



\- Understand natural language searches.

\- Produce deterministic results.

\- Return relevant software quickly.

\- Minimize user effort.

\- Build user trust through consistency.



\---



\# Search Inputs



Users may search using:



\- Application names

\- Categories

\- Features

\- Problems to solve

\- Alternatives

\- Platforms

\- Technologies

\- Abbreviations

\- Common aliases



Examples:



\- design app

\- canva

\- vscode alternative

\- remove background

\- windows video editor

\- markdown notes



\---



\# Search Understanding



The search engine should understand intent rather than exact wording.



Different queries expressing the same intent should produce nearly identical results.



Examples:



\- photo editor

\- image editor

\- picture editor



\---



\# Deterministic Behavior



The same query should always produce:



\- the same ranking

\- the same layout

\- the same ordering



Results change only when:



\- application data changes

\- ranking rules change



Never because of randomness.



\---



\# Ranking Philosophy



Ranking should consider:



\- relevance

\- category match

\- feature match

\- metadata quality

\- aliases

\- synonyms



Ranking must never consider:



\- advertisements

\- payments

\- popularity manipulation



\---



\# Search Suggestions



Suggestions exist to assist users.



Suggestions never replace search.



Suggestions may include:



\- related searches

\- categories

\- common searches

\- aliases



Suggestions never display application cards.



\---



\# Query Processing



Each search should follow the same lifecycle:



Normalize



↓



Resolve aliases



↓



Resolve synonyms



↓



Apply deterministic ranking



↓



Generate layout seed



↓



Return results



\---



\# Search Results



Each result contains enough information for exploration.



Additional details appear progressively through interaction.



\---



\# Zero Results



If no exact match exists:



\- show closest matches

\- explain that no exact result exists

\- allow users to suggest missing software



Users should never reach a dead end.



\---



\# Performance Goals



Target:



\- Search begins immediately after submission.

\- Results feel instant.

\- Interface remains responsive throughout.



Performance targets will be defined separately.



\---



\# Out of Scope



This RFC does not define:



\- database schema

\- indexing implementation

\- Meilisearch configuration

\- frontend rendering

\- floating canvas algorithm



\---



\# Acceptance Criteria



The search engine is complete when:



\- natural language searches work consistently

\- deterministic behavior is preserved

\- users trust the results

\- future ranking improvements do not break consistency

\- implementation remains independent of UI

