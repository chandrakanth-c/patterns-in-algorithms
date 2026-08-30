# 12 - Top 'K' Elements (Heaps & Quickselect)

## 12.1 - Overview & Theoretical Foundations (CLRS Chapter 6 & 9)

* Finding the top $K$ largest, smallest, or most frequent elements can be solved using two fundamental paradigms:
  1. **Binary Min-Heap / Max-Heap (CLRS Chapter 6):**
     * A Min-Heap maintains the **heap property**: for every node $i$ other than root, $A[\text{PARENT}(i)] \le A[i]$.
     * Maintaining a Min-Heap of size $K$ allows keeping track of the $K$ largest elements seen so far. If a new element exceeds the root ($\text{min-element}$), we replace the root and call `Min-Heapify` in $\mathcal{O}(\log K)$ time.
  2. **Quickselect (`Randomized-Select`, CLRS 9.2):**
     * A divide-and-conquer selection algorithm based on Quicksort partitioning.
     * By only recursing into the partition containing the target $K$-th index, the recurrence becomes $T(n) = T(n/2) + \mathcal{O}(n) \implies \Theta(n)$ expected time.

---

## 12.2 - Properties of a problem that suggests Top 'K' Elements

* Asked to find the **$K$-th largest/smallest** element.
* Stream of incoming items where top $K$ candidates must be tracked online.
* Selecting $K$ elements when sorting the entire array ($\mathcal{O}(n \log n)$) is wasteful because $K \ll n$.

---

## 12.3 - Classic Example: Kth Largest Element in an Array

### Java Implementation

```java
import java.util.PriorityQueue;

public class TopKElements {

    public static int findKthLargest(int[] nums, int k) {
        // Min-heap storing the k largest elements
        PriorityQueue<Integer> minHeap = new PriorityQueue<>(k);

        for (int num : nums) {
            minHeap.offer(num);
            if (minHeap.size() > k) {
                minHeap.poll(); // Evict smallest element among the top candidates
            }
        }

        return minHeap.peek();
    }
}
```

---

### Go Implementation

```go
package main

import (
	"container/heap"
)

// IntMinHeap implements heap.Interface for a min-heap of ints
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// FindKthLargest finds the kth largest element using a min-heap of size k
func FindKthLargest(nums []int, k int) int {
	h := &IntMinHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}
```

---

## 12.4 - Time & Space Complexity Analysis

* **Min-Heap Approach:**
  * **Time Complexity:** $\mathcal{O}(n \log K)$ — Inserting $n$ elements into a heap of size $K$ costs $\log K$ each.
  * **Space Complexity:** $\mathcal{O}(K)$ auxiliary memory.
* **Quickselect (`Randomized-Select`) Approach:**
  * **Time Complexity:** $\Theta(n)$ expected average time, $\mathcal{O}(n^2)$ worst case.
  * **Space Complexity:** $\mathcal{O}(1)$ iterative auxiliary space.

---

## 12.5 - Classic LeetCode & CLRS Benchmarks

* **Randomized-Select / Deterministic Select in Linear Time** (CLRS 9.2, 9.3)
* **Kth Largest Element in an Array** (LeetCode #215)
* **Top K Frequent Elements** (LeetCode #347)
* **K Closest Points to Origin** (LeetCode #973)
* **Sort Characters By Frequency** (LeetCode #451)
* **Reorganize String** (LeetCode #767)

---

## 12.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 6: Heapsort & Priority Queues (pp. 161–181)
  * Chapter 9: Medians and Order Statistics (Section 9.2: Selection in expected linear time pp. 228–234)
* https://leetcode.com/problems/kth-largest-element-in-an-array/
* https://techinterviewhandbook.org/algorithms/heap/
