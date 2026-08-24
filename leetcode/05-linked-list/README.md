# Linked List

## What it is

A chain of nodes (`Val` + `Next` pointer). No random access (no O(1) index),
but O(1) insert/delete once you're at the right node — the opposite
trade-off from arrays.

## Signals in a problem

- "linked list", "reverse", "detect a cycle", "merge two sorted lists",
  "find the middle node", "Nth node from the end".

## Core techniques

- **Dummy head node:** `dummy := &ListNode{Next: head}` sidesteps special-casing
  "what if the head itself needs to change" (removal, merging).
- **Fast/slow pointers (Floyd's):** slow moves 1 step, fast moves 2 — fast
  reaches the end when slow is at the middle; they meet inside a cycle if
  one exists.
- **Iterative reversal:** walk with `prev, curr, next` and flip `curr.Next`
  one node at a time — O(1) space, vs O(n) for a recursive reversal's call
  stack.

## Complexity cheatsheet

| Operation | Singly Linked List |
|-----------|--------------------|
| Access by index | O(n) |
| Insert/delete at known node | O(1) |
| Search | O(n) |

## Pitfalls

- Losing the reference to the rest of the list before rewiring `Next`
  (always save `next := curr.Next` before mutating `curr.Next`).
- Off-by-one when finding "the Nth node from the end" or the middle with
  even-length lists — decide up front whether "middle" means the first or
  second of the two center nodes.
- Nil-pointer panics from not checking `head == nil` / `head.Next == nil`
  base cases.

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
