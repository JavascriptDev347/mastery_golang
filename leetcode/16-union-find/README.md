# Union-Find (Disjoint Set Union)

## What it is

A structure that tracks a partition of elements into disjoint sets, with two
near-O(1) operations: `Find(x)` (which set is x in) and `Union(x, y)` (merge
x's and y's sets). Built specifically for "are these connected" questions
without a full graph traversal each time.

## Signals in a problem

- "number of connected components", "redundant connection / cycle
  detection in an undirected graph", "accounts merge", "friend circles" —
  anything about grouping/connectivity that gets queried incrementally.

## Core techniques

- **Parent array:** `parent[i]` starts as `i` (its own root); `Find` walks
  parent pointers to the root.
- **Path compression:** during `Find`, point every visited node directly at
  the root — flattens the tree so future `Find` calls are near O(1).
  ```go
  func find(x int) int {
      if parent[x] != x {
          parent[x] = find(parent[x]) // path compression
      }
      return parent[x]
  }
  ```
- **Union by rank/size:** always attach the smaller tree under the bigger
  tree's root, instead of arbitrarily — keeps trees shallow. Combined with
  path compression, gives near-constant amortized time per operation
  (inverse Ackermann function — for any practical n, treat it as O(1)).

## Complexity cheatsheet

| Operation | Naive | With path compression + union by rank |
|-----------|-------|------------------------------------------|
| Find / Union | O(n) worst case | O(α(n)) amortized ≈ O(1) |

## Pitfalls

- Skipping path compression *or* union by rank still works correctly, just
  degrades toward O(n) per operation on adversarial input — both
  optimizations are cheap, always include them.
- Forgetting that `Union` should check `find(x) == find(y)` first (already
  same set → this edge is redundant/creates a cycle) *before* merging — the
  detection signal for "redundant connection" problems.

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
