# LeetCode + DSA Journal

Daily practice log: 1–2 problems a day, organized **by DSA topic first**, not by
LeetCode problem number. The goal isn't just "solved it" — for every problem I
write down every approach I can find (brute force → optimal), with complexity,
so the pattern actually sticks.

## Folder structure

```
leetcode/
  README.md                 <- this file (index + conventions)
  TEMPLATE.md                <- copy this into a new problem folder
  01-arrays-strings/
    README.md                <- topic notes: what/when/how, patterns, pitfalls
    0001-two-sum/
      README.md               <- problem statement, examples, approach notes
      solution.go              <- one function per approach
      solution_test.go         <- table-driven tests for every approach
  02-two-pointers/
  03-sliding-window/
  ...
```

Each topic folder is a numbered stage in the order I'm learning them (roughly
easiest/most-foundational first). Numbers are just for browsing order — a
problem can obviously use ideas from more than one topic; file it under the
technique I'm actually practicing.

## Conventions

- **Problem folder name:** `NNNN-kebab-case-title` where `NNNN` is the
  LeetCode problem number, zero-padded to 4 digits (e.g. `0001-two-sum`,
  `0053-maximum-subarray`).
- **Package per problem:** each `solution.go` is its own package
  (`package twosum`, `package maxsubarray`, ...) — no `main.go`, no `func
  main`. This keeps every problem independently testable and avoids
  `main`/`main` collisions across the module.
- **One function per approach**, named for the approach, not generically:
  `BruteForce`, `TwoPointers`, `HashMap`, `Dp`, `Greedy`, etc. Never overwrite
  a working approach with a "better" one — keep both, so the trade-off stays
  visible.
- **Tests are the runner.** No `main()` — verify with:
  ```
  go test ./leetcode/01-arrays-strings/0001-two-sum/...
  go test ./leetcode/...          # run everything
  ```
- **Every problem folder has a `README.md`** (start from `TEMPLATE.md`)
  documenting the problem, examples/constraints, and a table comparing every
  approach tried (time, space, trade-off, why you'd pick it).
- **Every topic folder has a `README.md`** with the DSA theory: what the
  technique is, the signals in a problem statement that suggest it, the core
  patterns, complexity cheatsheet, and common pitfalls — written *before* or
  *while* solving problems in that topic, not after.

## Progress log

| Date | Topic | Problem | Difficulty | Approaches |
|------|-------|---------|------------|------------|
| 2026-08-24 | Arrays & Strings | [0001. Two Sum](01-arrays-strings/0001-two-sum) | Easy | Brute Force, Two Pointers |

## Topics

| # | Topic | Focus |
|---|-------|-------|
| [01](01-arrays-strings) | Arrays & Strings | Iteration, prefix sums, in-place tricks |
| [02](02-two-pointers) | Two Pointers | Sorted-array / opposite-end scans |
| [03](03-sliding-window) | Sliding Window | Fixed & variable-size subarray/substring |
| [04](04-hashing) | Hashing | HashMap/HashSet for O(1) lookup tricks |
| [05](05-linked-list) | Linked List | Fast/slow pointers, reversal, merging |
| [06](06-stacks-queues) | Stacks & Queues | Monotonic stack, BFS-ready queues |
| [07](07-recursion-backtracking) | Recursion & Backtracking | State-space search, pruning |
| [08](08-binary-search) | Binary Search | On sorted arrays and on the answer space |
| [09](09-sorting) | Sorting | Custom comparators, sort-based tricks |
| [10](10-trees) | Trees | Traversals, BST invariants, recursion on trees |
| [11](11-heaps-priority-queue) | Heaps / Priority Queue | Top-K, streaming median, scheduling |
| [12](12-graphs) | Graphs | BFS/DFS, topological sort, shortest paths |
| [13](13-dynamic-programming) | Dynamic Programming | Memoization, tabulation, state design |
| [14](14-greedy) | Greedy | Local-optimal choices with a proof sketch |
| [15](15-tries) | Tries | Prefix trees for string sets |
| [16](16-union-find) | Union-Find | Disjoint sets, connectivity, cycles |
| [17](17-bit-manipulation) | Bit Manipulation | XOR tricks, bitmasks, counting bits |
| [18](18-math-geometry) | Math & Geometry | Number theory, simulation, geometry |
