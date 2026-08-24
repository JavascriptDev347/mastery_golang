# Arrays & Strings

## What it is

The base data structure almost everything else builds on: contiguous memory,
O(1) index access, O(n) search/insert/delete in the middle. Strings are just
arrays of bytes/runes with the same complexity trade-offs, plus immutability
in Go (`string` concatenation allocates a new one every time).

## Signals in a problem

- "array of integers/strings", "in-place", "subarray", "substring"
- Anything asking for a running total, a count, or a rearrangement without
  extra structure.

## Core techniques

- **Prefix sums** — precompute cumulative sums so any range sum is O(1).
- **In-place overwrite** — use a slow-write pointer while a fast pointer
  reads, to compact/dedupe/filter without extra memory.
- **Kadane's algorithm** — running max/min for max-subarray-style problems.
- **`strings.Builder`** in Go for O(n) string building instead of repeated
  `+=` (which is O(n) per concat, O(n^2) total).

## Complexity cheatsheet

| Operation | Array/Slice | Notes |
|-----------|-------------|-------|
| Index access | O(1) | |
| Search (unsorted) | O(n) | |
| Insert/delete at end | O(1) amortized | Go slice `append` |
| Insert/delete at index | O(n) | shifts elements |

## Pitfalls

- Off-by-one on slice bounds (`nums[i:j]` is `[i, j)`).
- Mutating a slice you're iterating over shifts indices under you.
- Go strings are UTF-8 byte sequences — indexing by `[]byte` vs iterating
  runes (`for _, r := range s`) gives different units when non-ASCII is
  involved.

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| 0001 | [Two Sum](0001-two-sum) | Easy | Brute Force, Two Pointers |
