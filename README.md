> NOTE : I do not claim any rights to the content in [this repository](https://github.com/chandrakanth-c/patterns-in-algorithms/tree/main). 
>
> All sources are cited at the end of every topic file with due diligence.
> Theoretical foundations, recurrences, invariants, and proofs are grounded in **"Introduction to Algorithms (4th Edition)" by Cormen, Leiserson, Rivest, and Stein (CLRS)**.
> Every pattern includes full, idiomatic implementations in both **Java** and **Go**.

# Patterns in Algorithms & Data Structures

A comprehensive guide to popular problem-solving design patterns in Data Structures and Algorithms with rigorous theoretical foundations (CLRS 4th Ed.), dual-language code templates (**Java** & **Go**), time/space complexity analysis, and curated LeetCode benchmarks.

---

## 📚 Table of Contents

### 1. Algorithm Design Paradigms
* 1 - [Dynamic Programming (DP)](dynamic-programming.md) *(CLRS Chapter 14: Overlapping Subproblems, Optimal Substructure, Memoization, Tabulation)*
* 2 - [Divide and Conquer](divide-and-conquer.md) *(CLRS Chapter 4: Divide-Conquer-Combine, Recurrences, Master Theorem)*

### 2. Core Algorithmic Problem-Solving Patterns
* 3 - [Two Pointers](two-pointers.md) *(CLRS Chapter 2 & 10: Loop Invariants, Pair Search)*
* 4 - [Sliding Window](sliding-window.md) *(CLRS Chapter 17: Amortized Aggregate Analysis)*
* 5 - [Fast & Slow Pointers (Floyd's Tortoise and Hare)](fast-and-slow-pointers.md) *(Cycle Detection Proof & Cycle Start)*
* 6 - [Merge Intervals](merge-intervals.md) *(CLRS Chapter 15: Greedy Choice Property, Activity Selection)*
* 7 - [Cyclic Sort](cyclic-sort.md) *(CLRS Chapter 8: Sorting in Linear Time, Direct Indexing)*
* 8 - [In-place Reversal of a Linked List](in-place-reversal-linked-list.md) *(CLRS Chapter 10: Pointer Redirection & Sentinels)*
* 9 - [Monotonic Stack & Monotonic Queue](monotonic-stack.md) *(CLRS Chapter 17: Accounting Method, Next Greater Element)*
* 10 - [Prefix Sum & Difference Array](prefix-sum.md) *(CLRS Chapter 2: 1D & 2D Cumulative Sums, Range Updates)*
* 11 - [Modified Binary Search & Binary Search on Answer](modified-binary-search.md) *(CLRS Chapter 2 & 9: Invariants & Monotonic Predicates)*
* 12 - [Top 'K' Elements (Heaps & Quickselect)](top-k-elements.md) *(CLRS Chapter 6 & 9: Priority Queues, `Randomized-Select`)*
* 13 - [Two Heaps (Median of a Stream)](two-heaps.md) *(CLRS Chapter 6: Dual-Heap Partition Invariants)*
* 14 - [K-way Merge](k-way-merge.md) *(CLRS Chapter 6 Exercise 6.5-9: Min-Heap Multi-Way Merging)*
* 15 - [Subsets, Permutations & Backtracking](subsets-and-backtracking.md) *(CLRS Appendix C: Combinatorial State-Space Trees)*
* 16 - [Tree & Graph Traversals (BFS & DFS)](tree-and-graph-bfs-dfs.md) *(CLRS Chapter 20: Shortest Paths, Parenthesis Theorem)*
* 17 - [Matrix Traversal (Grid BFS/DFS & Flood Fill)](matrix-traversal.md) *(CLRS Chapter 20: Implicit Planar Graphs & Multi-Source BFS)*
* 18 - [Topological Sort (DAG Dependency Resolution)](topological-sort.md) *(CLRS Section 20.4: Kahn's Algorithm & DFS Finishing Times)*
* 19 - [Union-Find (Disjoint Set Union - DSU)](union-find.md) *(CLRS Chapter 19: Path Compression, Union by Rank, $\mathcal{O}(m \alpha(n))$)*
* 20 - [Bit Manipulation & Bitmasking](bit-manipulation.md) *(CLRS Chapter 31 & Appendix B: Bitwise Identities, Subset State DP)*

### 3. Data Structure Design Patterns
* 21 - [Trie (Prefix Tree)](trie.md) *(CLRS Chapter 31: Ordered Prefix Retrieval, Maximum XOR)*
* 22 - [Segment Tree & Binary Indexed Tree (Fenwick)](segment-tree.md) *(CLRS Chapter 13: Augmenting Data Structures, Range Queries)*
* 23 - [Cache Design Patterns (LRU & LFU)](lru-cache.md) *(CLRS Chapter 10 & 11: Hash Table + Doubly Linked List Sentinels)*

---

## 📖 Pattern Recognition Cheat Sheet

| "The Tell" in Problem Statement | Suggested Pattern | Primary Paradigm / Data Structure |
| :--- | :--- | :--- |
| Sorted array; find pair/triplet satisfying condition | **Two Pointers** | Converging Indices |
| Contiguous longest/shortest/optimal subarray or substring | **Sliding Window** | Rolling State (Amortized $\mathcal{O}(1)$) |
| Linked list cycle detection, middle node, index cycle | **Fast & Slow Pointers** | Floyd's Cycle Algorithm |
| Overlapping time intervals, meeting scheduling | **Merge Intervals** | Greedy Interval Sorting |
| Array elements in bounded range $[1, n]$ with missing/duplicates | **Cyclic Sort** | Linear-Time In-Place Permutation |
| Reversing subsegment of linked list in-place | **In-Place Linked List Reversal** | 3-Pointer Redirection |
| Next greater/smaller element, histogram rectangle | **Monotonic Stack** | Monotonic LIFO Invariant |
| Frequent range sum queries or range additions | **Prefix Sum / Difference Array** | Cumulative Precomputation |
| Monotonic search space; minimize maximum / search rotated array | **Modified Binary Search** | Halving Search Bounds |
| Find top / smallest / most frequent $K$ elements | **Top 'K' Elements** | Min-Heap / Quickselect |
| Dynamic median calculation from continuous data stream | **Two Heaps** | Dual-Heap Partitioning |
| Merging $K$ sorted arrays / matrices | **K-Way Merge** | Min-Heap Priority Queue |
| Combinatorial permutations, subsets, board search (Sudoku/N-Queens) | **Backtracking** | State-Space Tree Exploration |
| Shortest path on unweighted graph / Level-by-level processing | **BFS** | FIFO Queue Frontier |
| Connected components, dynamic graph connectivity, cycle detection | **Union-Find (DSU)** | Disjoint Set Forest ($\alpha(n)$) |
| Task dependencies, build ordering, prerequisite courses | **Topological Sort** | In-Degree BFS / DFS Finishing |
| String prefix matching, autocomplete dictionary | **Trie** | Prefix Tree |
| Dynamic range sum / min / max with updates | **Segment Tree / Fenwick Tree** | Tree / Binary Index Decomposition |
| Fast $\mathcal{O}(1)$ key lookup + eviction policy (LRU/LFU) | **LRU / LFU Cache** | Hash Table + Doubly Linked List |
