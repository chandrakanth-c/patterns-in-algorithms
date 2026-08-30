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

* **K-way Merge Problem** (CLRS Exercise 6.5-9)
* **Merge k Sorted Lists** (LeetCode #23)
* **Kth Smallest Element in a Sorted Matrix** (LeetCode #378)
* **Find K Pairs with Smallest Sums** (LeetCode #373)
* **Smallest Range Covering Elements from K Lists** (LeetCode #632)

---

## 14.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 6: Heapsort & Priority Queues (Exercise 6.5-9 pp. 179)
* https://leetcode.com/problems/merge-k-sorted-lists/
* https://en.wikipedia.org/wiki/K-way_merge_algorithm
