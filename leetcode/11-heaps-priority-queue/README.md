# Heaps / Priority Queue

## What it is

A binary tree (usually array-backed) where every parent is smaller (min-heap)
or larger (max-heap) than its children — gives O(1) peek and O(log n)
insert/pop of the min/max, without full sorting.

## Signals in a problem

- "Kth largest/smallest", "top K frequent elements", "merge K sorted
  lists", "median of a data stream", "task scheduler".

## Core techniques

- **`container/heap`** in Go: implement `Len`, `Less`, `Swap`, `Push`, `Pop`
  to satisfy `heap.Interface`, then use `heap.Init/Push/Pop`. It's
  boilerplate-heavy but the standard tool — worth building a small reusable
  int-heap and generic-ish wrapper once.
- **Fixed-size heap for "top K":** keep a heap of size K; for each new
  element, push then pop the worst if size exceeds K — O(n log k) instead
  of sorting everything (O(n log n)).
- **Two heaps for streaming median:** a max-heap for the lower half, a
  min-heap for the upper half, kept balanced in size — median is O(1) peek
  from whichever heap (or both) after each O(log n) insert.

## Complexity cheatsheet

| Operation | Time |
|-----------|------|
| Peek min/max | O(1) |
| Insert | O(log n) |
| Pop min/max | O(log n) |
| Build heap from n elements | O(n) (not O(n log n) — heapify is linear) |

## Pitfalls

- Go's `container/heap` needs `Fix`/`Push`/`Pop` calls to keep the heap
  invariant — direct slice manipulation bypasses it and corrupts the
  structure silently.
- Using a max-heap for "Kth largest" naively (O(n log n) to build + K pops)
  when a size-K min-heap gets it in O(n log k) — pick the heap whose root
  is the element you'd want to *evict*, not the one you want.

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
