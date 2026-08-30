# 14 - K-way Merge

## 14.1 - Overview

* The **K-way Merge** pattern is used to solve problems that involve merging $K$ sorted arrays, lists, or matrices into a single sorted sequence.
* Instead of merging lists pair-by-pair, a **Min-Heap of size $K$** is created containing the first element of each of the $K$ sorted inputs.
* The smallest element is popped from the heap into the output, and the next element from that same list is inserted into the heap.

---

## 14.2 - Properties of a problem that suggests K-way Merge

* The problem provides **$K$ sorted arrays / linked lists** or a row-wise/column-wise sorted matrix.
* You need to find the overall sorted order, the $K$-th smallest item among all lists, or the smallest range covering elements from every list.

---

## 14.3 - Classic Example: Merge K Sorted Lists

### Java Implementation

```java
import java.util.PriorityQueue;

public class KWayMerge {

    static class ListNode {
        int val;
        ListNode next;
        ListNode(int val) { this.val = val; }
    }

    public static ListNode mergeKLists(ListNode[] lists) {
        if (lists == null || lists.length == 0) return null;

        // Min-heap ordered by node value
        PriorityQueue<ListNode> minHeap = new PriorityQueue<>(
            (a, b) -> Integer.compare(a.val, b.val)
        );

        // Add head of each list to heap
        for (ListNode root : lists) {
            if (root != null) minHeap.offer(root);
        }

        ListNode dummy = new ListNode(0);
        ListNode tail = dummy;

        while (!minHeap.isEmpty()) {
            ListNode smallestNode = minHeap.poll();
            tail.next = smallestNode;
            tail = tail.next;

            // Push next node from the same list into heap
            if (smallestNode.next != null) {
                minHeap.offer(smallestNode.next);
            }
        }

        return dummy.next;
    }
}
```

---

## 14.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(N \log K)$ where $N$ is the total number of elements across all $K$ lists. Each insertion/deletion in the heap of size $K$ takes $\mathcal{O}(\log K)$.
* **Space Complexity:** $\mathcal{O}(K)$ auxiliary space to store elements in the Min-Heap.

---

## 14.5 - Classic LeetCode Problems

* **Merge k Sorted Lists** (LeetCode #23)
* **Kth Smallest Element in a Sorted Matrix** (LeetCode #378)
* **Find K Pairs with Smallest Sums** (LeetCode #373)
* **Smallest Range Covering Elements from K Lists** (LeetCode #632)

---

## 14.6 - Sources used for this file:
https://leetcode.com/problems/merge-k-sorted-lists/ <br>
https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d53c08d2bb2d978e2c02 <br>
https://www.geeksforgeeks.org/merge-k-sorted-arrays/ <br>
https://en.wikipedia.org/wiki/K-way_merge_algorithm
