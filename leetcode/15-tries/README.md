# Tries (Prefix Trees)

## What it is

A tree where each edge is labeled with one character, and each root-to-node
path spells a prefix. Built for fast prefix operations on a set of strings
— something a hash set of strings can't do efficiently.

## Signals in a problem

- "prefix", "autocomplete", "word search on a board with a dictionary",
  "implement a search engine's typeahead", "longest common prefix of many
  words".

## Core techniques

- **Node = children map + end-of-word flag:**
  ```go
  type TrieNode struct {
      children map[byte]*TrieNode
      isEnd    bool
  }
  ```
- **Insert/search walk one character at a time**, creating child nodes on
  insert if missing, following existing ones on search — O(L) for a word of
  length L, independent of how many words are already stored.
- **`startsWith` (prefix check)** is the same walk as `search` but without
  requiring `isEnd` at the final node — this is the operation a trie gives
  you that a plain hash set can't do without scanning every entry.

## Complexity cheatsheet

| Operation | Trie | Hash Set of strings |
|-----------|------|----------------------|
| Insert / exact search | O(L) | O(L) average |
| Prefix search | O(L) | O(n · L) — must scan every word |
| Space | O(total characters), shared across common prefixes | O(total characters) |

## Pitfalls

- Forgetting `isEnd` — without it, `search("app")` returns true just
  because "apple" was inserted and shares that prefix path.
- `map[byte]*TrieNode` per node has overhead; for lowercase-only alphabets
  a fixed `[26]*TrieNode` array is faster and simpler in Go when the
  problem guarantees the character set.

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
