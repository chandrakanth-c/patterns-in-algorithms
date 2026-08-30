> NOTE : I do not claim any rights to the content in [this repository](https://github.com/chandrakanth-c/patterns-in-algorithms/tree/main). 
>
> All sources are listed at the end of every topic file with due diligence.
> For LLM sources - provider (gemini, chatgpt, claude etc) will be mentioned.

# Patterns in Algorithms & Data Structures

A comprehensive collection of popular problem-solving design patterns in Data Structures and Algorithms with explanations, Java templates, time/space complexity analysis, and curated LeetCode problems.

---

## 📚 Table of Contents

### 1. Algorithm Design Paradigms
* 1 - [Dynamic Programming (DP)](dynamic-programming.md)
* 2 - [Divide and Conquer](divide-and-conquer.md)

### 2. Core Algorithmic Problem-Solving Patterns
* 3 - [Two Pointers](two-pointers.md)
* 4 - [Sliding Window](sliding-window.md)
* 5 - [Fast & Slow Pointers (Hare & Tortoise)](fast-and-slow-pointers.md)
* 6 - [Merge Intervals](merge-intervals.md)
* 7 - [Cyclic Sort](cyclic-sort.md)
* 8 - [In-place Reversal of a Linked List](in-place-reversal-linked-list.md)
* 9 - [Monotonic Stack & Monotonic Queue](monotonic-stack.md)
* 10 - [Prefix Sum & Difference Array](prefix-sum.md)
* 11 - [Modified Binary Search & Binary Search on Answer](modified-binary-search.md)
* 12 - [Top 'K' Elements (Heaps & Quickselect)](top-k-elements.md)
* 13 - [Two Heaps (Median of a Stream)](two-heaps.md)
* 14 - [K-way Merge](k-way-merge.md)
* 15 - [Subsets, Permutations & Backtracking](subsets-and-backtracking.md)
* 16 - [Tree & Graph Traversals (BFS & DFS)](tree-and-graph-bfs-dfs.md)
* 17 - [Matrix Traversal (Grid BFS/DFS & Flood Fill)](matrix-traversal.md)
* 18 - [Topological Sort (DAG Dependency Resolution)](topological-sort.md)
* 19 - [Union-Find (Disjoint Set Union - DSU)](union-find.md)
* 20 - [Bit Manipulation & Bitmasking](bit-manipulation.md)

### 3. Data Structure Design Patterns
* 21 - [Trie (Prefix Tree)](trie.md)
* 22 - [Segment Tree & Binary Indexed Tree (Fenwick)](segment-tree.md)
* 23 - [Cache Design Patterns (LRU & LFU)](lru-cache.md)

---

## 📖 Pattern Recognition Cheat Sheet

| "The Tell" in Problem Statement | Suggested Pattern |
| :--- | :--- |
| Sorted array; find pair/triplet satisfying condition | **Two Pointers** |
| Contiguous longest/shortest/optimal subarray or substring | **Sliding Window** |
| Linked list cycle detection, middle node, index cycle | **Fast & Slow Pointers** |
| Overlapping time intervals, meeting scheduling | **Merge Intervals** |
| Array elements in bounded range $[1, n]$ with missing/duplicates | **Cyclic Sort** |
| Reversing subsegment of linked list in-place | **In-Place Linked List Reversal** |
| Next greater/smaller element, histogram rectangle | **Monotonic Stack** |
| Frequent range sum queries or range additions | **Prefix Sum / Difference Array** |
| Monotonic search space; minimize maximum / search rotated array | **Modified Binary Search** |
| Find top / smallest / most frequent $K$ elements | **Top 'K' Elements** |
| Dynamic median calculation from continuous data stream | **Two Heaps** |
| Merging $K$ sorted arrays / matrices | **K-Way Merge** |
| Combinatorial permutations, subsets, board search (Sudoku/N-Queens) | **Backtracking** |
| Shortest path on unweighted graph / Level-by-level processing | **BFS** |
| Connected components, dynamic graph connectivity, cycle detection | **Union-Find (DSU)** |
| Task dependencies, build ordering, prerequisite courses | **Topological Sort** |
| String prefix matching, autocomplete dictionary | **Trie** |
| Dynamic range sum / min / max with updates | **Segment Tree / Fenwick Tree** |
| Fast $\mathcal{O}(1)$ key lookup + eviction policy (LRU/LFU) | **LRU / LFU Cache** |
