# Graphs

## What it is

Nodes + edges, directed or undirected, weighted or not. Trees are a special
case (connected, no cycles); "real" graphs add cycles, disconnected
components, and direction to worry about.

## Signals in a problem

- "grid" (implicit graph, 4/8-directional neighbors), "islands", "course
  schedule" (dependency = directed edge), "network/connections",
  "shortest path", "clone a graph".

## Core techniques

- **BFS:** shortest path in an *unweighted* graph, level-by-level
  processing (see [06-stacks-queues](../06-stacks-queues) for the queue).
- **DFS:** connectivity, cycle detection, path existence, exhausting a
  search space — recursive or with an explicit stack.
- **Visited set:** mandatory for any graph with cycles, or you loop forever
  — a `map[node]bool` or a boolean grid for matrix-based graphs.
- **Topological sort** (Kahn's algorithm, BFS on in-degrees, or DFS
  postorder reversed): ordering under dependency constraints — only valid
  on a DAG (no cycles).
- **Union-Find** for connectivity/cycle questions without needing full
  traversal — see [16-union-find](../16-union-find).
- **Dijkstra's algorithm** (heap-based, see
  [11-heaps-priority-queue](../11-heaps-priority-queue)) for shortest path
  with non-negative weights.

## Complexity cheatsheet

| Algorithm | Time | Space |
|-----------|------|-------|
| BFS / DFS | O(V + E) | O(V) |
| Topological sort | O(V + E) | O(V) |
| Dijkstra (heap) | O((V + E) log V) | O(V) |

## Pitfalls

- Forgetting the visited set → infinite loop on any cyclic graph.
- Grid-as-graph: forgetting bounds checks on neighbor coordinates before
  indexing.
- BFS finds shortest path only when edges are unweighted (or all equal
  weight) — using BFS on a weighted graph silently gives a wrong answer.
- Directed vs undirected: an undirected edge `u-v` must be added to *both*
  adjacency lists.

## Problems solved

| # | Problem | Difficulty | Approaches |
|---|---------|------------|------------|
| — | — | — | — |
