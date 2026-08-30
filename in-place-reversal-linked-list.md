# 8 - In-place Reversal of a Linked List

## 8.1 - Overview

* The **In-place Reversal of a Linked List** pattern reverses the links between nodes in a singly linked list without allocating new nodes in memory.
* It operates directly on the existing `ListNode` pointers by keeping track of three pointers:
  * `prev`: Reference to the previously reversed node.
  * `curr`: Reference to the current node being reversed.
  * `next`: Temporary reference to preserve the remaining unreversed list (`curr.next`).

---

## 8.2 - Properties of a problem that suggests In-place Reversal

* Need to reverse a linked list completely, between specific indices $[m, n]$, or in groups of size $k$.
* Strict constraint: $\mathcal{O}(1)$ auxiliary space (no modifying node values, no stack allocation).

---

## 8.3 - Classic Example: Reverse a Sub-list (Between positions left and right)

### Java Implementation

```java
public class InPlaceLinkedListReversal {

    static class ListNode {
        int val;
        ListNode next;
        ListNode(int val) { this.val = val; }
    }

    public static ListNode reverseBetween(ListNode head, int left, int right) {
        if (head == null || left == right) return head;

        ListNode dummy = new ListNode(0);
        dummy.next = head;
        ListNode prev = dummy;

        // 1. Skip the first (left - 1) nodes
        for (int i = 0; i < left - 1; i++) {
            prev = prev.next;
        }

        // 2. Reverse sublist from left to right
        ListNode curr = prev.next;
        ListNode sublistTail = curr;
        ListNode sublistPrev = null;

        for (int i = 0; i < right - left + 1; i++) {
            ListNode nextTemp = curr.next;
            curr.next = sublistPrev;
            sublistPrev = curr;
            curr = nextTemp;
        }

        // 3. Connect with the rest of the list
        prev.next = sublistPrev;
        sublistTail.next = curr;

        return dummy.next;
    }
}
```

---

## 8.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(n)$ because every node is visited and its pointer redirected at most once.
* **Space Complexity:** $\mathcal{O}(1)$ auxiliary space.

---

## 8.5 - Classic LeetCode Problems

* **Reverse Linked List** (LeetCode #206)
* **Reverse Linked List II** (LeetCode #92)
* **Reverse Nodes in k-Group** (LeetCode #25)
* **Rotate List** (LeetCode #61)
* **Reorder List** (LeetCode #143)
* **Swapping Nodes in a Linked List** (LeetCode #1721)

---

## 8.6 - Sources used for this file:
https://leetcode.com/problems/reverse-linked-list-ii/ <br>
https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d45d08d2bb2d978e22c9 <br>
https://www.geeksforgeeks.org/reverse-a-linked-list/ <br>
https://techinterviewhandbook.org/algorithms/linked-list/
