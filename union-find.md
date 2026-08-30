# 19 - Union-Find (Disjoint Set Union - DSU)

## 19.1 - Overview

* A **Disjoint-Set data structure** (also called a **Union-Find data structure** or **DSU**) tracks a set of elements partitioned into a number of disjoint (non-overlapping) subsets.
* Provides two primary operations:
  1. `find(i)`: Determine which subset a particular element is in (finds the representative root of the set).
  2. `union(i, j)`: Join two subsets into a single subset.
* **Optimizations:**
  * **Path Compression:** Flattens the tree during `find()`, making nodes point directly to the root.
  * **Union by Rank / Size:** Attaches the smaller depth tree under the root of the deeper tree to prevent skew.

---

## 19.2 - Properties of a problem that suggests Union-Find

* Maintaining dynamic **connected components** in an undirected graph.
* Detecting cycles during incremental edge additions (e.g., **Kruskal's Minimum Spanning Tree Algorithm**).
* Grouping equivalent items together (e.g. Accounts Merge, sentence similarity).

---

## 19.3 - Classic Example: DSU Implementation with Path Compression & Union by Rank

### Java Implementation

```java
public class UnionFind {

    private int[] parent;
    private int[] rank;
    private int count; // number of disjoint components

    public UnionFind(int n) {
        this.count = n;
        this.parent = new int[n];
        this.rank = new int[n];
        for (int i = 0; i < n; i++) {
            parent[i] = i;
            rank[i] = 0;
        }
    }

    // Find with Path Compression
    public int find(int x) {
        if (parent[x] != x) {
            parent[x] = find(parent[x]); // Point directly to root
        }
        return parent[x];
    }

    // Union by Rank
    public boolean union(int x, int y) {
        int rootX = find(x);
        int rootY = find(y);

        if (rootX == rootY) {
            return false; // Already in the same set (cycle detected)
        }

        if (rank[rootX] < rank[rootY]) {
            parent[rootX] = rootY;
        } else if (rank[rootX] > rank[rootY]) {
            parent[rootY] = rootX;
        } else {
            parent[rootY] = rootX;
            rank[rootX]++;
        }

        count--;
        return true;
    }

    public int getCount() {
        return count;
    }
}
```

---

## 19.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(\alpha(n))$ per `find` and `union` operation, where $\alpha$ is the **Inverse Ackermann Function** (for all practical values of $n \le 10^{80}$, $\alpha(n) < 5$, rendering operations effectively $\mathcal{O}(1)$).
* **Space Complexity:** $\mathcal{O}(n)$ to store `parent` and `rank` arrays.

---

## 19.5 - Classic LeetCode Problems

* **Number of Connected Components in an Undirected Graph** (LeetCode #323)
* **Redundant Connection** (LeetCode #684)
* **Accounts Merge** (LeetCode #721)
* **Graph Valid Tree** (LeetCode #261)
* **Number of Operations to Make Network Connected** (LeetCode #1319)
* **Earliest Possible Day of Full Bloom / Kruskal's MST**

---

## 19.6 - Sources used for this file:
https://en.wikipedia.org/wiki/Disjoint-set_data_structure <br>
https://www.geeksforgeeks.org/disjoint-set-data-structures/ <br>
https://leetcode.com/explore/learn/card/graph/618/disjoint-set/ <br>
https://techinterviewhandbook.org/algorithms/graph/
