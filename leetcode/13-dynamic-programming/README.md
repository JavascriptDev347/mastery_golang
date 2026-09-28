# Dynamic Programming

## What it is

Breaking a problem into overlapping subproblems, solving each exactly once,
and reusing the result — either top-down with memoization, or bottom-up by
filling a table in dependency order.

## Signals in a problem

- "count the number of ways", "minimum/maximum ... to reach/achieve X",
  "can you partition/reach ...", and a naive recursive solution that
  recomputes the same subproblem repeatedly (exponential blowup).

## Core techniques

- **Identify the state:** what varies between subproblems? (index, remaining
  capacity, remaining target...) — the state defines the memo key /
  table dimensions.
- **Recurrence relation:** how does `f(state)` relate to `f(smaller
  state)`? This is the actual "insight" step — get this right and the code
  is mechanical.
- **Top-down (memoization):** write the natural recursion first, add a
  `map[state]result` (or array) cache, return early on a cache hit. Easiest
  to derive correctly.
- **Bottom-up (tabulation):** fill a table (`dp[i]` or `dp[i][j]`) in an
  order where dependencies are already computed — usually faster in
  practice (no recursion overhead) and easier to then optimize space.
- **Space optimization:** if `dp[i]` only depends on `dp[i-1]` (or a fixed
  window back), collapse the table to O(1)/O(k) rolling variables instead
  of O(n) — a common follow-up.

## Complexity cheatsheet

Time is usually `(number of distinct states) × (work per state)`. Space is
`(number of distinct states)` for the memo/table, often reducible.

## Pitfalls

- Wrong base case → every answer built on top of it is wrong.
- Memoizing on a key that's incomplete (missing a piece of state that
  actually varies) — silently returns cached results for what are actually
  different subproblems.
- Confusing "optimal substructure exists" (DP applies) with "greedy works"
  (a *local* optimal choice is always safe) — see [14-greedy](../14-greedy)
  for when the simpler tool is enough.

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
