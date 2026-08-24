# Recursion & Backtracking

## What it is

Recursion: a function solving a problem by solving smaller instances of
itself. Backtracking: recursion that explores a decision tree, and
*undoes* a choice when it doesn't pan out — try, recurse, undo.

## Signals in a problem

- "all possible ...", "generate all permutations/combinations/subsets",
  "N-Queens", "Sudoku solver", "word search on a grid".
- Anything phrased as "how many ways" or "find all" over a combinatorial
  space.

## Core techniques

- **Base case first.** Every recursive function needs a condition that
  returns without recursing further, or it never terminates.
- **Choose → recurse → un-choose:** the backtracking skeleton —
  ```go
  func backtrack(path []int, ...) {
      if done(...) { record(path); return }
      for _, choice := range choices {
          path = append(path, choice)      // choose
          backtrack(path, ...)              // recurse
          path = path[:len(path)-1]         // un-choose
      }
  }
  ```
- **Pruning:** cut a branch early the moment it can't possibly lead to a
  valid answer (e.g. remaining sum already negative) — the difference
  between backtracking finishing instantly vs timing out.

## Complexity cheatsheet

Usually exponential (O(2^n) for subsets, O(n!) for permutations) — that's
inherent to enumerating the search space, not a sign something's wrong.
Pruning reduces the *constant*/effective branching factor, not the
worst-case class.

## Pitfalls

- Appending a slice to the result without copying it — Go slices share the
  underlying array, so `result = append(result, path)` after the loop
  finishes often stores the *same* backing array for every entry, corrupted
  by later un-choose steps. Copy: `append(result, append([]int{}, path...))`.
- Forgetting the un-choose step, silently turning backtracking into a plain
  (wrong) DFS that never explores siblings correctly.

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
