# Sorting

## What it is

Reordering elements by some key — usually not the end goal itself, but the
enabler for a two-pointer, greedy, or interval-merging technique that only
works once the data has a known order.

## Signals in a problem

- "sort the array", "merge intervals", "meeting rooms", "largest number
  formed by concatenation" — anything where a smart comparator turns a hard
  problem trivial once order is fixed.

## Core techniques

- **`sort.Slice` with a custom `less` function** in Go — the practical tool
  for 95% of these problems instead of hand-writing a sort algorithm.
- **Sort, then sweep:** sort by start (or by whichever key resolves
  ambiguity), then a single linear pass handles the rest — the pattern
  behind interval merging, non-overlapping intervals, etc.
- Know the standard algorithms conceptually even if you don't hand-write
  them: merge sort (stable, O(n log n) guaranteed, O(n) space), quicksort
  (in-place, O(n log n) average but O(n^2) worst case), counting sort
  (O(n+k) when the key range k is small — bucket by value instead of
  comparing).

## Complexity cheatsheet

| Algorithm | Time | Space | Stable |
|-----------|------|-------|--------|
| Merge sort | O(n log n) | O(n) | Yes |
| Quicksort | O(n log n) avg, O(n^2) worst | O(log n) | No |
| Counting sort | O(n + k) | O(k) | Yes |
| Go `sort.Slice` | O(n log n) | — | No (use `sort.SliceStable` if needed) |

## Pitfalls

- Forgetting `sort.SliceStable` when relative order of equal elements
  matters (e.g. stable grouping after a first sort pass).
- A custom `less` function that isn't a valid strict weak ordering
  (inconsistent comparisons) — silently produces a wrong/undefined sort.

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
