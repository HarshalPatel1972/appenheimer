# Appenheimer Project Rules

## Definition of Done (Completeness)
From this point onward, never report a phase as "complete" unless it satisfies all four conditions:
1. **Implemented** – the code exists and is wired into the application.
2. **Verified** – automated tests or runtime evidence demonstrate it works.
3. **Observable** – health checks, logs, metrics, or traces expose its behavior.
4. **Used** – the running application actually exercises the feature.

If any of those conditions are missing, classify the feature as **Scaffolded**, **Partial**, or **Not Implemented**. Never use "complete" for architecture alone.
