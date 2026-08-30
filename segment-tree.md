# 22 - Segment Tree & Binary Indexed Tree (Fenwick)

## 22.1 - Overview

* A **Segment Tree** is a tree data structure used for storing information about intervals or segments. It allows answering range queries (e.g. range sum, range minimum/maximum) and updating elements in $\mathcal{O}(\log n)$ time.
* **Lazy Propagation** allows updating an entire range $[L, R]$ in $\mathcal{O}(\log n)$ by deferring updates to children nodes until they are visited.
* **Binary Indexed Tree (Fenwick Tree / BIT):** A more space-efficient, array-based alternative supporting prefix sums and point updates in $\mathcal{O}(\log n)$ using binary bit decomposition (`i & (-i)`).

---

## 22.2 - Comparison: Segment Tree vs Fenwick Tree vs Prefix Sum

| Structure | Build Time | Range Query | Point Update | Range Update |
| :--- | :--- | :--- | :--- | :--- |
| **Prefix Sum** | $\mathcal{O}(n)$ | $\mathcal{O}(1)$ | $\mathcal{O}(n)$ | $\mathcal{O}(n)$ |
| **Fenwick Tree (BIT)** | $\mathcal{O}(n)$ | $\mathcal{O}(\log n)$ | $\mathcal{O}(\log n)$ | $\mathcal{O}(\log n)$ (with diff BIT) |
| **Segment Tree** | $\mathcal{O}(n)$ | $\mathcal{O}(\log n)$ | $\mathcal{O}(\log n)$ | $\mathcal{O}(\log n)$ (with Lazy Prop) |

---

## 22.3 - Classic Example: Segment Tree for Range Sum Query with Point Updates

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
        if (r < start || end < l) return 0; // Out of range
        if (l <= start && end <= r) return tree[node]; // Completely inside range

        int mid = start + (end - start) / 2;
        return querySum(2 * node + 1, start, mid, l, r)
             + querySum(2 * node + 2, mid + 1, end, l, r);
    }
}
```

---

## 22.4 - Time & Space Complexity

* **Build Tree:** $\mathcal{O}(n)$ time.
* **Range Query:** $\mathcal{O}(\log n)$ time.
* **Point / Range Update:** $\mathcal{O}(\log n)$ time.
* **Space Complexity:** $\mathcal{O}(4n)$ auxiliary array space for the Segment Tree ($\mathcal{O}(n)$ for Fenwick Tree).

---

## 22.5 - Classic LeetCode Problems

* **Range Sum Query - Mutable** (LeetCode #307)
* **Count of Smaller Numbers After Self** (LeetCode #315)
* **The Skyline Problem** (LeetCode #218)
* **Falling Squares** (LeetCode #699)
* **Reverse Pairs** (LeetCode #493)

---

## 22.6 - Sources used for this file:
https://en.wikipedia.org/wiki/Segment_tree <br>
https://cp-algorithms.com/data_structures/segment_tree.html <br>
https://cp-algorithms.com/data_structures/fenwick.html <br>
https://www.geeksforgeeks.org/segment-tree-data-structure/ <br>
https://leetcode.com/problems/range-sum-query-mutable/
