# 19 - Union-Find (Disjoint Set Union - DSU)

## 19.1 - Overview & Theoretical Foundations (CLRS Chapter 19)

* A **Disjoint-Set Data Structure** maintains a collection $\mathcal{S} = \{S_1, S_2, \dots, S_k\}$ of disjoint dynamic sets.
* Each set is identified by a representative member $x \in S_i$.
* **Three Fundamental Operations (CLRS 19.1):**
  1. `MAKE-SET(x)`: Creates a new set whose only member is $x$.
  2. `UNION(x, y)`: Merges the dynamic sets containing $x$ and $y$ into a single new set.
  3. `FIND-SET(x)`: Returns a pointer to the representative of the unique set containing $x$.
* **Heuristics for Efficiency (CLRS 19.3):**
  1. **Union by Rank:** Keeps tree depth small by pointing the root with smaller rank to the root with larger rank.
  2. **Path Compression:** During `FIND-SET(x)`, makes each node on the find path point directly to the root.
* **CLRS Theorem 19.14:** A sequence of $m$ `MAKE-SET`, `UNION`, and `FIND-SET` operations on $n$ elements takes $\mathcal{O}(m \alpha(n))$ time in the worst case, where $\alpha(n)$ is the extremely slowly growing **Inverse Ackermann Function** ($\alpha(n) \le 4$ for all $n \le 10^{80}$).

---

## 19.2 - Properties of a problem that suggests Union-Find

* Dynamic connectivity queries on an undirected graph.
* Cycle detection during incremental edge additions (e.g. **Kruskal's Minimum Spanning Tree**, CLRS 21.2).
* Grouping equivalence relations (e.g. Accounts Merge).

---

## 19.3 - Classic Example: DSU with Path Compression & Union by Rank

### Java Implementation

```java
public class UnionFind {

    private int[] parent;
    private int[] rank;
    private int componentCount;

    public UnionFind(int n) {
        this.componentCount = n;
        this.parent = new int[n];
        this.rank = new int[n];
        for (int i = 0; i < n; i++) {
            parent[i] = i;
            rank[i] = 0;
        }
    }

    // FIND-SET with Path Compression
    public int find(int x) {
        if (parent[x] != x) {
            parent[x] = find(parent[x]); // Two-pass path compression
        }
        return parent[x];
    }

    // UNION with Union by Rank
    public boolean union(int x, int y) {
        int rootX = find(x);
        int rootY = find(y);

        if (rootX == rootY) {
            return false; // In the same set (cycle detected)
        }

        if (rank[rootX] < rank[rootY]) {
            parent[rootX] = rootY;
        } else if (rank[rootX] > rank[rootY]) {
            parent[rootY] = rootX;
        } else {
            parent[rootY] = rootX;
            rank[rootX]++;
        }

        componentCount--;
        return true;
    }

    public int getComponentCount() {
        return componentCount;
    }
}
```

---

### Go Implementation

```go
package main

// UnionFind represents the disjoint-set data structure with rank and path compression
type UnionFind struct {
	parent         []int
	rank           []int
	componentCount int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &UnionFind{
		parent:         parent,
		rank:           rank,
		componentCount: n,
	}
}

// Find finds the representative of the set containing x with path compression
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union joins two sets by rank; returns false if already in same set
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}

	uf.componentCount--
	return true
}

func (uf *UnionFind) GetComponentCount() int {
	return uf.componentCount
}
```

---

## 19.4 - Time & Space Complexity Analysis

* **Time Complexity:** $\Theta(m \alpha(n))$ for $m$ total operations — effectively $\mathcal{O}(1)$ amortized per operation.
* **Space Complexity:** $\Theta(n)$ auxiliary memory to hold the `parent` and `rank` slices.

---

## 19.5 - Classic LeetCode & CLRS Benchmarks

* **Kruskal's Minimum Spanning Tree Algorithm** (CLRS 21.2)
* **Number of Connected Components in an Undirected Graph** (LeetCode #323)
* **Redundant Connection** (LeetCode #684)
* **Accounts Merge** (LeetCode #721)
* **Graph Valid Tree** (LeetCode #261)
* **Number of Operations to Make Network Connected** (LeetCode #1319)

---

## 19.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 19: Data Structures for Disjoint Sets (pp. 556–585)
  * Section 19.1: Disjoint-set operations (pp. 556–562)
  * Section 19.3: Disjoint-set forests (pp. 566–572)
  * Section 19.4: Analysis of union by rank with path compression (pp. 573–583)
  * Section 21.2: Kruskal's algorithm (pp. 631–637)
* https://en.wikipedia.org/wiki/Disjoint-set_data_structure
