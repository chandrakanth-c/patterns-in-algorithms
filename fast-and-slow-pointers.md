# 5 - Fast & Slow Pointers (Hare & Tortoise)

## 5.1 - Overview

* The **Fast & Slow pointer** approach (also known as **Floyd's Tortoise and Hare Algorithm**) is a pointer algorithm that uses two pointers which move through the sequence (typically an array or linked list) at different speeds.
* Commonly, the `slow` pointer advances by **1 step** while the `fast` pointer advances by **2 steps**.
* If a cycle exists, the fast pointer will eventually catch up to the slow pointer from behind. If no cycle exists, the fast pointer reaches the end (`null`).

---

## 5.2 - Properties of a problem that suggests Fast & Slow Pointers

* Detecting **cycles / loops** in a Linked List or finite state sequence.
* Finding the **middle element** of a singly linked list in a single pass without knowing total length.
* Finding the **$k$-th element from the end** or start node of a cycle.
* Solving array problems where elements represent directed index jumps with bounded values $[1, n]$.

---

## 5.3 - Classic Example: Linked List Cycle Detection & Cycle Start

### Java Implementation

```java
public class FastAndSlowPointers {

    static class ListNode {
        int val;
        ListNode next;
        ListNode(int val) { this.val = val; }
    }

    // 1. Detect if cycle exists
    public static boolean hasCycle(ListNode head) {
        ListNode slow = head;
        ListNode fast = head;

        while (fast != null && fast.next != null) {
            slow = slow.next;         // 1 step
            fast = fast.next.next;    // 2 steps

            if (slow == fast) {
                return true; // Fast caught up to slow
            }
        }

        return false;
    }

    // 2. Find the node where cycle begins
    public static ListNode findCycleStart(ListNode head) {
        ListNode slow = head;
        ListNode fast = head;

        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
            if (slow == fast) {
                // Phase 2: Move slow to head; advance both by 1 step until they meet
                ListNode ptr1 = head;
                ListNode ptr2 = slow;
                while (ptr1 != ptr2) {
                    ptr1 = ptr1.next;
                    ptr2 = ptr2.next;
                }
                return ptr1; // Start of cycle
            }
        }
        return null;
    }
}
```

---

## 5.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(n)$
  * If no cycle: fast pointer traverses $n$ nodes in $n/2$ iterations $\implies \mathcal{O}(n)$.
  * If cycle: once slow enters the cycle of length $C$, fast catches up within at most $C$ steps $\implies \mathcal{O}(n)$.
* **Space Complexity:** $\mathcal{O}(1)$ auxiliary space as only two node pointers are allocated.

---

## 5.5 - Classic LeetCode Problems

* **Linked List Cycle** (LeetCode #141)
* **Linked List Cycle II** (LeetCode #142)
* **Middle of the Linked List** (LeetCode #876)
* **Palindrome Linked List** (LeetCode #234)
* **Happy Number** (LeetCode #202)
* **Find the Duplicate Number** (LeetCode #287)

---

## 5.6 - Sources used for this file:
https://en.wikipedia.org/wiki/Cycle_detection#Floyd's_tortoise_and_hare <br>
https://www.geeksforgeeks.org/how-does-floyds-slow-and-fast-pointers-approach-work/ <br>
https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d3c808d2bb2d978e1e79 <br>
https://leetcode.com/problems/linked-list-cycle/
