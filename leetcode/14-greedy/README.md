# Greedy

## What it is

Make the locally-best choice at each step and never reconsider it, trusting
that the sequence of local choices adds up to a globally optimal answer.
Faster and simpler than DP — but only correct when the problem actually has
the "greedy-choice property," which needs a justification, not just a hunch.

## Signals in a problem

- "maximum number of ...", "minimum number of intervals/coins/jumps",
  "assign/schedule to minimize/maximize ...", often solvable after sorting
  by the right key (see [09-sorting](../09-sorting)).

## Core techniques

- **Sort, then greedily pick** — the most common shape: sort by end time
  (interval scheduling), by ratio (fractional knapsack), by start time
  (merge intervals), then sweep once making the obvious choice at each
  step.
- **Exchange argument** (the way to *prove* a greedy choice is safe, even
  informally): assume an optimal solution that doesn't make the greedy
  choice, show swapping in the greedy choice doesn't make it worse — if you
  can't construct this argument even loosely, be suspicious the problem
  actually needs DP.

## Complexity cheatsheet

Usually O(n log n) — dominated by the sort — with an O(n) sweep after.
Much cheaper than the DP alternative when it's actually valid.

## Pitfalls

- **The big one:** assuming greedy works without checking. Coin change with
  arbitrary denominations is the canonical counterexample where "always
  pick the biggest coin" fails — that's a DP problem
  ([13-dynamic-programming](../13-dynamic-programming)), not a greedy one.
  If you can't justify why the local choice never needs to be undone,
  don't reach for greedy.
- Picking the wrong sort key (e.g. sorting intervals by start when the
  problem needs sorting by end).

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
