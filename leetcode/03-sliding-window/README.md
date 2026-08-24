# Sliding Window

## What it is

A contiguous window `[left, right]` over an array/string that expands and
contracts, maintaining some invariant, instead of recomputing from scratch
for every possible subarray/substring.

## Signals in a problem

- "contiguous subarray/substring", "longest/shortest ... with at most/exactly
  K ...", "maximum sum of a window of size K".

## Core techniques

- **Fixed-size window:** move `right` forward, and when the window exceeds
  size K, move `left` forward too — maintain a running aggregate (sum,
  count map) incrementally instead of recomputing.
- **Variable-size window:** grow `right` until the invariant breaks, then
  shrink `left` until it holds again, tracking the best window seen.
- Usually paired with a hash map/set to track window contents (char counts,
  distinct elements).

## Complexity cheatsheet

O(n) — each pointer moves forward at most n times total, even though it
looks like a nested loop. That amortized-O(n) argument is the thing to be
able to explain, not just recite.

## Pitfalls

- Shrinking the window with the wrong condition (off-by-one on whether the
  invariant is "still valid" vs "just broke").
- Forgetting to update the best-answer *inside* the loop at the right point
  (after expand, after shrink, or both — depends on whether you want max or
  min window).

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
