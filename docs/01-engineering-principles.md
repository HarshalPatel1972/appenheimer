# Engineering Principles

> Build software that is easy to understand, easy to maintain, and easy to evolve.

---

# Purpose

These principles define how Appenheimer is engineered.

Every implementation, code review, refactor, and architectural decision should follow these principles.

If an implementation violates these principles, it should be reconsidered regardless of whether it works.

---

# Core Principles

## Design the data first.

Before writing code, design the data model.

Code changes frequently.

Data lasts for years.

---

## Keep functions small.

A function should have one clear responsibility.

If it becomes difficult to explain, it is probably doing too much.

---

## Avoid deep nesting.

Prefer early exits over nested conditionals.

The happy path should be obvious.

---

## Use early returns.

Handle invalid cases immediately.

Keep the main logic readable.

---

## Remove unnecessary abstractions.

Do not introduce interfaces, layers, or patterns until they solve a real problem.

Abstractions should reduce complexity, not create it.

---

## Write obvious code.

Prioritize readability over cleverness.

Future contributors should understand the code without needing explanations.

---

## Delete code whenever possible.

Every line of code becomes a maintenance responsibility.

If something is unnecessary, remove it.

---

## Optimize only after measuring.

Never optimize based on assumptions.

Measure first.

Optimize the actual bottleneck.

---

## Name things clearly.

Names should describe purpose.

Avoid abbreviations, clever wording, and unnecessary jargon.

---

## Make maintenance easy.

Every decision should reduce future maintenance.

Choose the solution that will still make sense years from now.

---

## Consistency over preference.

Follow established patterns throughout the project.

Consistency is more valuable than personal style.

---

## Build only what is needed.

Do not build for hypothetical future requirements.

Solve today's problems well.

Future requirements can be implemented when they actually exist.

---

## Every dependency must justify itself.

Before adding a dependency, ask:

- Does it solve a real problem?
- Is it well maintained?
- Does it reduce maintenance?
- Can the standard library or existing code solve this instead?

If the answer is no, do not add it.

---

## One source of truth.

Every piece of data should have exactly one owner.

Caches, indexes, and derived data must never become the primary source.

---

## Replaceability.

Components should be loosely coupled.

Replacing one technology should not require rewriting the entire system.

---

# Code Review Checklist

Every implementation should answer "Yes" to the following:

- Is it correct?
- Is it simple?
- Is it easy to understand?
- Does it follow the architecture?
- Is the data model correct?
- Is the naming clear?
- Is the complexity justified?
- Can it be tested?
- Will it still make sense two years from now?

---

# Guiding Principle

> Every piece of complexity must earn its place.

Features, abstractions, dependencies, and architecture all have a maintenance cost.

If the value does not clearly outweigh that cost, choose the simpler solution.