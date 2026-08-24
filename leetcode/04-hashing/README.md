# Hashing

## What it is

Trading space for time: a `map[K]V` gives O(1) average lookup/insert, so any
"have I seen this before" or "what's the complement of X" question collapses
from O(n^2) to O(n).

## Signals in a problem

- "have you seen this element before", "find the pair/complement",
  "count frequency of ...", "duplicate detection", "group by ...".
- The moment a brute force does `for i { for j { ... } }` looking for a
  match — ask "could a map remember what I've already seen?"

## Core techniques

- **One-pass complement lookup** (Two Sum): for each element, check if
  `target - nums[i]` is already in the map *before* inserting `nums[i]` —
  avoids matching an element with itself.
- **Frequency map:** `map[T]int` for counting (anagrams, majority element).
- **Set via `map[T]struct{}`** in Go — `struct{}` costs zero bytes, use it
  instead of `map[T]bool` for pure membership tests.
- **Group by key:** `map[K][]V` to bucket elements (group anagrams).

## Complexity cheatsheet

| Operation | Go map | Notes |
|-----------|--------|-------|
| Lookup/insert/delete | O(1) average | O(n) worst case (hash collisions), not a practical concern here |
| Space | O(n) | the trade-off for O(1) time |

## Pitfalls

- Go map iteration order is randomized — never rely on it for
  deterministic output.
- Using a struct or slice as a map key requires it to be comparable (slices
  and maps aren't valid keys directly).
- Reading zero-value on a missing key silently (`m[k]` returns the zero
  value even if `k` isn't present) — use the two-value form `v, ok := m[k]`
  when "not present" is meaningfully different from "present with zero
  value".

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
