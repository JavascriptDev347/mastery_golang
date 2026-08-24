# Bit Manipulation

## What it is

Operating directly on an integer's binary representation — `& | ^ ~ << >>`
— for problems where the "trick" is a bitwise identity rather than an
algorithm.

## Signals in a problem

- "single number" (everything else appears twice/thrice), "count set
  bits", "power of two/four", "subsets" (bitmask enumeration), "without
  using + or -".

## Core techniques

- **XOR cancels pairs:** `a ^ a == 0` and `a ^ 0 == a`, and XOR is
  commutative/associative — XOR-ing a whole list cancels every value that
  appears an even number of times, leaving the odd-one-out. The core trick
  behind "Single Number".
- **`n & (n-1)` clears the lowest set bit** — repeat-count this to count set
  bits (Brian Kernighan's algorithm), or check `n & (n-1) == 0` for "is n a
  power of two".
- **`n & (-n)` isolates the lowest set bit** — useful when a problem needs
  it directly.
- **Bitmask for subsets:** a set of size n has 2^n subsets; iterate
  `mask` from `0` to `1<<n - 1`, and bit `i` of `mask` says whether element
  `i` is included — turns "all subsets" backtracking into a flat loop.
- **Shifts as multiply/divide by 2:** `x << 1` == `x*2`, `x >> 1` == `x/2`
  for non-negative integers — sometimes relevant, rarely the actual point.

## Complexity cheatsheet

Bit tricks are O(1) or O(number of bits) — always dramatically cheaper than
the naive loop-and-count they replace, on fixed-width integers.

## Pitfalls

- Go integers are signed by default — right-shifting a negative int (`>>`)
  is an *arithmetic* shift (sign-extends), not logical; know which one the
  problem needs, use `uint` if you specifically need logical shift
  behavior.
- Operator precedence: `&` and `^` bind looser than `==` in Go, so
  `if x & 1 == 1` needs the parens some other languages don't:
  `if (x & 1) == 1` reads safer even where Go doesn't strictly require it.

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
