# Trees

## What it is

A hierarchical structure: each node points to child nodes, no cycles. A
Binary Search Tree (BST) additionally guarantees `left < node < right` at
every node, enabling O(log n) search when balanced.

## Signals in a problem

- "binary tree", "BST", "validate/balance/height/diameter of a tree",
  "lowest common ancestor", "serialize/deserialize".

## Core techniques

- **DFS traversals** (recursive, O(h) call-stack space where h = height):
  - Preorder (node, left, right) — good for copying/serializing a tree.
  - Inorder (left, node, right) — visits a BST in sorted order.
  - Postorder (left, right, node) — good when children must be processed
    before the parent (deletion, computing subtree aggregates like height).
- **BFS / level-order traversal** using a queue (see
  [06-stacks-queues](../06-stacks-queues)) — process one full depth level
  at a time.
- **Recursion returning info upward:** many tree problems (height, balance,
  diameter) are solved by a helper that returns *both* the answer for the
  subtree *and* extra state to the parent, computed bottom-up.

## Complexity cheatsheet

| Operation | Balanced BST | Unbalanced (worst case) |
|-----------|--------------|--------------------------|
| Search/insert/delete | O(log n) | O(n) — degenerates to a linked list |
| Traversal (any) | O(n) | O(n) |
| Space (recursive traversal) | O(h) = O(log n) | O(h) = O(n) |

## Pitfalls

- Confusing "balanced" with "sorted" — a BST is always searchable in order
  (inorder = sorted), but only O(log n) fast if it's also height-balanced.
- Recursive traversal stack overflow on pathologically deep/unbalanced
  trees (rare on LeetCode's constraints, but the reason iterative
  traversals with an explicit stack exist).
- Off-by-one confusing "height" (edges) vs "number of nodes" definitions —
  check which the problem means.

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
