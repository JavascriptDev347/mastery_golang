# Binary Search

## What it is

Halve the search space every step by testing the midpoint, using a
monotonic property (sorted array, or a "yes/no" predicate that flips exactly
once). O(log n) instead of O(n).

## Signals in a problem

- "sorted array", "find the index/boundary of ...", "minimize/maximize X
  such that condition holds" (binary search *on the answer*, not on an
  array — e.g. "minimum capacity to ship in D days").

## Core techniques

- **Classic search:** find an exact target in a sorted array.
- **Boundary search (lower/upper bound):** find the first/last index where
  a predicate becomes true/false — the template that generalizes to almost
  every binary search problem:
  ```go
  lo, hi := 0, n
  for lo < hi {
      mid := lo + (hi-lo)/2
      if ok(mid) { hi = mid } else { lo = mid + 1 }
  }
  // lo is the boundary
  ```
- **Binary search on the answer space:** when the array isn't sorted but a
  *predicate over possible answers* is monotonic (true for all X >= some
  threshold), binary search the threshold directly.

## Complexity cheatsheet

O(log n) time, O(1) space (iterative). The whole point versus O(n) linear
scan — always worth asking "is there a monotonic property here?" before
reaching for a full scan.

## Pitfalls

- `mid := (lo + hi) / 2` can overflow for huge bounds — use
  `lo + (hi-lo)/2` (matters less in Go's 64-bit int range, but it's the
  habit that matters).
- Off-by-one on `lo <= hi` vs `lo < hi`, and `hi = mid` vs `hi = mid - 1` —
  pick one template (above) and stick to it instead of re-deriving each
  time.
- Infinite loop when `mid` rounds to `lo` and the update doesn't move `lo`
  forward (`lo = mid` instead of `lo = mid + 1` on the wrong branch).

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
