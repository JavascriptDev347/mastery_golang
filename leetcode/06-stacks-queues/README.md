# Stacks & Queues

## What it is

Stack: LIFO — push/pop from one end, O(1) each. Queue: FIFO — push at one
end, pop from the other, O(1) each. In Go, both are usually just a
`[]T` slice (`append` for push, slice re-slicing for pop) or a
`container/list` for a queue where you don't want amortized-cost slice
shifting.

## Signals in a problem

- Stack: "matching parentheses/brackets", "next greater/smaller element",
  "evaluate expression", "undo operation", anything needing "the most
  recent unmatched thing".
- Queue: "process in order", "level by level" (paired with BFS, see
  [12-graphs](../12-graphs) and [10-trees](../10-trees)).

## Core techniques

- **Matching/validity via stack:** push opening brackets, pop-and-compare on
  closing ones; valid iff the stack empties exactly at the end.
- **Monotonic stack:** keep the stack increasing or decreasing; popping
  elements that violate the order as you go finds "next greater/smaller
  element" in O(n) total instead of O(n^2).
- **Queue via two stacks / stack via one queue** — a classic "implement X
  using Y" exercise for understanding the two structures' duality.

## Complexity cheatsheet

Both structures: O(1) push and pop. The whole value of a monotonic stack is
that even though it looks like nested work, each element is pushed and
popped at most once → O(n) total, not O(n^2).

## Pitfalls

- Go slices as a stack: popping is `s = s[:len(s)-1]` — forgetting this
  reuses/aliases the underlying array in subtle ways if you kept another
  reference to it.
- Checking `len(stack) == 0` before popping/peeking to avoid a panic.
- For "queue via slice", `s = s[1:]` for dequeue is O(1) amortized but leaks
  the underlying array until GC'd — fine for LeetCode, not for a long-lived
  production queue.

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
