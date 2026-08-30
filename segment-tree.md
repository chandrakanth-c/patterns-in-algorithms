# 22 - Segment Tree & Binary Indexed Tree (Fenwick)

## 22.1 - Overview & Theoretical Foundations (CLRS Chapter 13: Augmenting Data Structures)

* When range queries (e.g. Range Sum, Range Min/Max) and point/range updates are both required dynamically, static structures (Prefix Sums) fail because updates cost $\mathcal{O}(n)$.
* **Segment Tree:**
  * A full binary tree where each leaf corresponds to an array element and each internal node represents an aggregated query value over a segment $[L, R]$.
  * **Tree Invariant:** For node representing $[L, R]$, its left child covers $[L, \text{mid}]$ and right child covers $[\text{mid}+1, R]$, where $\text{mid} = \lfloor (L+R)/2 \rfloor$.
  * **Lazy Propagation:** Delays range updates to descendants by caching the update value at higher-level nodes until children are queried, achieving $\mathcal{O}(\log n)$ range updates.
* **Binary Indexed Tree (Fenwick Tree / BIT):**
  * An array-based tree structure exploiting the binary representation of indices.
  * Node $i$ stores the sum of range $(i - \text{LSB}(i), i]$, where $\text{LSB}(i) = i \ \& \ (-i)$.

---

## 22.2 - Comparison: Segment Tree vs Fenwick Tree vs Prefix Sum

| Structure | Build Time | Range Query | Point Update | Range Update | Memory Overhead |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Prefix Sum** | $\mathcal{O}(n)$ | $\mathcal{O}(1)$ | $\mathcal{O}(n)$ | $\mathcal{O}(n)$ | $\mathcal{O}(n)$ |
| **Fenwick Tree (BIT)** | $\mathcal{O}(n)$ | $\mathcal{O}(\log n)$ | $\mathcal{O}(\log n)$ | $\mathcal{O}(\log n)$ | $\mathcal{O}(n)$ |
| **Segment Tree** | $\mathcal{O}(n)$ | $\mathcal{O}(\log n)$ | $\mathcal{O}(\log n)$ | $\mathcal{O}(\log n)$ | $\mathcal{O}(4n)$ |

---

## 22.3 - Classic Example: Segment Tree for Dynamic Range Sum Query

### Java Implementation

```java
public class SegmentTree {
    private int[] tree;
    private int n;

    public SegmentTree(int[] nums) {
        if (nums.length > 0) {
            n = nums.length;
            tree = new int[4 * n];
            buildTree(nums, 0, 0, n - 1);
        }
    }

    private void buildTree(int[] nums, int node, int start, int end) {
        if (start == end) {
            tree[node] = nums[start];
            return;
        }
        int mid = start + (end - start) / 2;
        int leftChild = 2 * node + 1;
        int rightChild = 2 * node + 2;

        buildTree(nums, leftChild, start, mid);
        buildTree(nums, rightChild, mid + 1, end);
        tree[node] = tree[leftChild] + tree[rightChild];
    }

    public void update(int index, int val) {
        updateVal(0, 0, n - 1, index, val);
    }

    private void updateVal(int node, int start, int end, int idx, int val) {
        if (start == end) {
            tree[node] = val;
            return;
        }
        int mid = start + (end - start) / 2;
        int leftChild = 2 * node + 1;
        int rightChild = 2 * node + 2;

        if (idx <= mid) {
            updateVal(leftChild, start, mid, idx, val);
        } else {
            updateVal(rightChild, mid + 1, end, idx, val);
        }
        tree[node] = tree[leftChild] + tree[rightChild];
    }

    public int sumRange(int left, int right) {
        return querySum(0, 0, n - 1, left, right);
    }

    private int querySum(int node, int start, int end, int l, int r) {
        if (r < start || end < l) return 0; // Disjoint
        if (l <= start && end <= r) return tree[node]; // Fully covered

        int mid = start + (end - start) / 2;
        return querySum(2 * node + 1, start, mid, l, r)
             + querySum(2 * node + 2, mid + 1, end, l, r);
    }
}
```

---

### Go Implementation

```go
package main

type SegmentTree struct {
	tree []int
	n    int
}

func NewSegmentTree(nums []int) *SegmentTree {
	n := len(nums)
	if n == 0 {
		return &SegmentTree{}
	}
	st := &SegmentTree{
		tree: make([]int, 4*n),
		n:    n,
	}
	st.buildTree(nums, 0, 0, n-1)
	return st
}

func (st *SegmentTree) buildTree(nums []int, node, start, end int) {
	if start == end {
		st.tree[node] = nums[start]
		return
	}
	mid := start + (end-start)/2
	leftChild := 2*node + 1
	rightChild := 2*node + 2

	st.buildTree(nums, leftChild, start, mid)
	st.buildTree(nums, rightChild, mid+1, end)
	st.tree[node] = st.tree[leftChild] + st.tree[rightChild]
}

func (st *SegmentTree) Update(index, val int) {
	st.updateVal(0, 0, st.n-1, index, val)
}

func (st *SegmentTree) updateVal(node, start, end, idx, val int) {
	if start == end {
		st.tree[node] = val
		return
	}
	mid := start + (end-start)/2
	leftChild := 2*node + 1
	rightChild := 2*node + 2

	if idx <= mid {
		st.updateVal(leftChild, start, mid, idx, val)
	} else {
		st.updateVal(rightChild, mid+1, end, idx, val)
	}
	st.tree[node] = st.tree[leftChild] + st.tree[rightChild]
}

func (st *SegmentTree) SumRange(left, right int) int {
	return st.querySum(0, 0, st.n-1, left, right)
}

func (st *SegmentTree) querySum(node, start, end, l, r int) int {
	if r < start || end < l {
		return 0
	}
	if l <= start && end <= r {
		return st.tree[node]
	}
	mid := start + (end-start)/2
	return st.querySum(2*node+1, start, mid, l, r) +
		st.querySum(2*node+2, mid+1, end, l, r)
}
```

---

## 22.4 - Time & Space Complexity Analysis

* **Build Time:** $\mathcal{O}(n)$ — Exactly $2n - 1$ tree nodes initialized.
* **Range Query:** $\mathcal{O}(\log n)$ — Visits at most 4 nodes per tree depth level.
* **Point Update:** $\mathcal{O}(\log n)$ — Traverses the single path from root to leaf.
* **Space Complexity:** $\mathcal{O}(4n)$ auxiliary array buffer.

---

## 22.5 - Classic LeetCode & Benchmark Problems

* **Interval Trees / Augmenting Trees** (CLRS 13.3, 14.3)
* **Range Sum Query - Mutable** (LeetCode #307)
* **Count of Smaller Numbers After Self** (LeetCode #315)
* **The Skyline Problem** (LeetCode #218)
* **Falling Squares** (LeetCode #699)
* **Reverse Pairs** (LeetCode #493)

---

## 22.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 13: Red-Black Trees & Augmenting Data Structures (Interval trees pp. 367–381)
* https://cp-algorithms.com/data_structures/segment_tree.html
* https://cp-algorithms.com/data_structures/fenwick.html
