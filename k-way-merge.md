# 14 - K-way Merge

## 14.1 - Overview & Theoretical Foundations (CLRS Chapter 6, Exercise 6.5-9)

* **K-way Merge** generalizes 2-way merging (as in MergeSort) to merge $K$ sorted sequences of total length $N$ into a single sorted list.
* Instead of merging pairwise ($\mathcal{O}(N \cdot K)$), a **Min-Heap of size $K$** is constructed containing the current minimum element of each of the $K$ sorted inputs.
* **Recurrence / Invariant:**
  * Extracting the minimum element from the Min-Heap gives the globally smallest unprocessed element in $\mathcal{O}(\log K)$ time.
  * We immediately replace it in the heap with the next element from the exact same list that provided the extracted element.
  * Repeating this $N$ times guarantees an overall sorted stream in $\mathcal{O}(N \log K)$ time.

---

## 14.2 - Properties of a problem that suggests K-way Merge

* Given **$K$ sorted arrays / linked lists**, or a 2D sorted matrix.
* Need to generate the full merged sorted order or find the $K$-th smallest element across all streams.

---

## 14.3 - Classic Example: Merge K Sorted Lists

### Java Implementation

```java
import java.util.PriorityQueue;

public class KWayMerge {

    public static class ListNode {
        public int val;
        public ListNode next;
        public ListNode(int val) { this.val = val; }
    }

    public static ListNode mergeKLists(ListNode[] lists) {
        if (lists == null || lists.length == 0) return null;

        // Min-heap of size K based on node value
        PriorityQueue<ListNode> minHeap = new PriorityQueue<>(
            (a, b) -> Integer.compare(a.val, b.val)
        );

        for (ListNode root : lists) {
            if (root != null) minHeap.offer(root);
        }

        ListNode dummy = new ListNode(0);
        ListNode tail = dummy;

        while (!minHeap.isEmpty()) {
            ListNode node = minHeap.poll();
            tail.next = node;
            tail = tail.next;

            if (node.next != null) {
                minHeap.offer(node.next);
            }
        }

        return dummy.next;
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

type ListNode struct {
	Val  int
	Next *ListNode
}

// NodeHeap implements heap.Interface for *ListNode
type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *NodeHeap) Push(x any)        { *h = append(*h, x.(*ListNode)) }
func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MergeKLists merges k sorted linked lists using a min-heap
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	h := &NodeHeap{}
	heap.Init(h)

	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}

	dummy := &ListNode{}
	tail := dummy

	for h.Len() > 0 {
		smallest := heap.Pop(h).(*ListNode)
		tail.Next = smallest
		tail = tail.Next

		if smallest.Next != nil {
			heap.Push(h, smallest.Next)
		}
	}

	return dummy.Next
}
```

---

## 14.4 - Time & Space Complexity Analysis

* **Time Complexity:** $\mathcal{O}(N \log K)$ where $N$ is total elements across all $K$ lists. Each element is inserted and extracted from the heap of size $K$ once ($\log K$ per operation).
* **Space Complexity:** $\mathcal{O}(K)$ auxiliary memory for the Min-Heap buffer.

---

## 14.5 - Classic LeetCode & CLRS Benchmarks

### 14.5.1 - Merge k Sorted Lists (LeetCode #23)

#### 1. Problem Statement
You are given an array of `k` linked-lists `lists`, each linked-list is sorted in ascending order. Merge all the linked-lists into one sorted linked-list and return it.

#### 2. Solution Link
* [Go Implementation](problems/k-way-merge/k_way_merge_problems.go) (Function: `MergeKLists`)
* [Java Implementation](k-way-merge.md) (Method: `mergeKLists`)

#### 3. Explanation
We use a Min-Heap to store the head nodes of all `k` lists. In each step, we extract the minimum node from the heap, attach it to our result list, and then push the next node from the same list into the heap. This ensures we always pick the globally smallest element among the current heads of the `k` lists.

#### 4. Conceptual Link to K-way Merge
This is the standard application of the pattern. It reduces a $K$-way comparison problem to a sequence of $\mathcal{O}(\log K)$ heap operations, optimizing the total time to $\mathcal{O}(N \log K)$.

### 14.5.2 - Kth Smallest Element in a Sorted Matrix (LeetCode #378)

#### 1. Problem Statement
Given an $n \times n$ matrix where each of the rows and columns is sorted in ascending order, return the $k$-th smallest element in the matrix.

#### 2. Solution Link
* [Go Implementation](problems/k-way-merge/k_way_merge_problems.go) (Function: `KthSmallest`)

#### 3. Explanation
We can treat each row of the matrix as a sorted list. We initialize a Min-Heap with the first element of each row. We perform $k$ extractions from the heap. Whenever we extract an element `matrix[i][j]`, we insert the next element from the same row `matrix[i][j+1]` into the heap. The $k$-th extracted element is our answer.

#### 4. Conceptual Link to K-way Merge
Treats a **Sorted Matrix as K-sorted lists**. It demonstrates that the pattern isn't limited to explicit lists but applies to any structure that can be viewed as multiple sorted streams.

### 14.5.3 - Smallest Range Covering Elements from K Lists (LeetCode #632)

#### 1. Problem Statement
You have `k` lists of sorted integers in non-decreasing order. Find the smallest range that includes at least one number from each of the `k` lists.

#### 2. Solution Link
* [Go Implementation](problems/k-way-merge/k_way_merge_problems.go) (Function: `SmallestRange`)

#### 3. Explanation
We maintain a Min-Heap containing one element from each of the `k` lists. We also track the maximum value currently in the heap. The current range is `[min_in_heap, max_in_heap]`. In each step, we pop the minimum element and push the next element from the same list. We update our "best range" if the new range is smaller. We stop when any list is exhausted.

#### 4. Conceptual Link to K-way Merge
This is a **Sliding Window over K-streams**. It uses the K-way merge mechanism to efficiently "slide" the window (by advancing the minimum pointer) while maintaining the property that one element from every list is present in the window.

---

## 14.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 6: Heapsort & Priority Queues (Exercise 6.5-9 pp. 179)
* https://leetcode.com/problems/merge-k-sorted-lists/
* https://en.wikipedia.org/wiki/K-way_merge_algorithm
