# Math & Geometry

## What it is

The catch-all for problems solved by a number-theory fact, a formula, or
simulating a geometric process directly — no "standard" data structure
pattern, the win is spotting the right identity.

## Signals in a problem

- "prime", "GCD/LCM", "modular arithmetic", "rotate a matrix", "spiral
  order", "points/rectangles overlap", anything phrased in terms of
  physical/geometric simulation ("robot moves", "game of life").

## Core techniques

- **Sieve of Eratosthenes** for "primes up to N" — O(n log log n), instead
  of trial-dividing each number (O(n·sqrt(n))).
- **GCD via Euclidean algorithm:** `gcd(a, b) = gcd(b, a%b)`, base case
  `gcd(a, 0) = a`; `lcm(a, b) = a*b/gcd(a,b)`.
- **Modular arithmetic:** when a problem says "answer mod 1e9+7", take the
  mod after every multiplication/addition, not just at the end, to avoid
  overflow — and remember Go's `%` can return negative for negative
  operands, unlike Python's.
- **Matrix rotation in place:** transpose then reverse rows (90° clockwise)
  — avoids O(n^2) extra space for a new matrix.
- **Simulate directly** when there's no clever shortcut — spiral traversal,
  Game of Life — just be careful to compute the *next* state fully before
  mutating (or use an out-of-bounds sentinel encoding to update in place).

## Complexity cheatsheet

Entirely problem-dependent — the whole category is "know the formula," so
the complexity win comes from picking O(log n) (GCD) or O(1) (a closed-form
formula) over a naive O(n) or O(n^2) simulation, when one exists.

## Pitfalls

- Integer overflow — Go's `int` is 64-bit on modern platforms so it's rare
  in practice here, but still worth checking against problem constraints
  (`n` up to 10^9, multiplied by another 10^9, etc.).
- Off-by-one in matrix boundary loops (rotating/traversing layer by layer).
- Floating point precision for geometry problems — prefer integer/rational
  arithmetic (cross-multiplication instead of division) when the inputs
  allow it.

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
