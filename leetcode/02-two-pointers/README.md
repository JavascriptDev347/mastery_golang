# Two Pointers

## What it is

Two indices walking through a (usually sorted) sequence — either from
opposite ends closing inward, or both moving forward at different speeds —
to avoid a nested loop.

## Signals in a problem

- "sorted array", "pair/triplet that sums to X", "palindrome check",
  "remove duplicates in-place", "container with most water".
- Anything a brute force solves in O(n^2) by checking all pairs, where the
  data is sorted or can be sorted.

## Core techniques

- **Opposite ends (converging):** `left, right := 0, len(nums)-1`; move
  whichever side helps based on a comparison, until they meet. Classic for
  sorted two-sum, container-with-most-water, valid palindrome.
- **Same direction (fast/slow):** one pointer scans, the other marks a
  write/boundary position. Classic for remove-duplicates, partitioning.
- **Fixed gap:** two pointers `k` apart, both advancing together.

## Complexity cheatsheet

Turns an O(n^2) nested-loop brute force into O(n) with O(1) extra space —
that trade is the entire point of this pattern. Requires the array be sorted
(or sortable, O(n log n) upfront) for the converging-ends variant to be
correct.

## Pitfalls

- Converging two pointers only works if the array is sorted *and* the
  problem doesn't need original indices preserved (sorting destroys index
  order — if you need the original indices, hash instead, see
  [04-hashing](../04-hashing)).
- Forgetting to skip duplicate values when the problem asks for unique
  pairs/triplets (e.g. 3Sum).

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
