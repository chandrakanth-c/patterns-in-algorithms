# 5 - Fast & Slow Pointers (Floyd's Tortoise and Hare)

## 5.1 - Overview & Mathematical Proof

* The **Fast & Slow Pointers** algorithm uses two pointers that traverse a sequence or linked structure at different speeds (typically slow moves 1 step and fast moves 2 steps).
* **Mathematical Proof of Cycle Detection:**
  * Let the distance from head to cycle entrance be $F$.
  * Let the perimeter/length of the cycle be $C$.
  * Let the meeting point inside the cycle be $a$ steps from the cycle start ($0 \le a < C$).
  * When they meet:
    $$\text{Distance}(\text{slow}) = F + a$$
    $$\text{Distance}(\text{fast}) = F + a + k \cdot C \quad (\text{for some integer } k \ge 1)$$
  * Since fast travels twice as fast:
    $$2(F + a) = F + a + k \cdot C \implies F + a = k \cdot C \implies F = k \cdot C - a = (k - 1)C + (C - a)$$
  * **Conclusion:** The distance from the list head to the cycle start ($F$) is equal to the distance from the meeting point to the cycle start moving forward ($C - a$). Moving one pointer to `head` and advancing both pointers 1 step at a time guarantees they meet precisely at the cycle entrance!

---

## 5.2 - Properties of a problem that suggests Fast & Slow Pointers

* Detecting **cycles** in linked lists, finite state transitions, or mathematical recurrences.
* Finding the **middle element** in a single pass with $\mathcal{O}(1)$ extra memory.
* Finding cycle length or the entry node to a cycle.

---

## 5.3 - Classic Example: Cycle Detection & Cycle Start Finding

### Java Implementation

```java
public class FastAndSlowPointers {

    public static class ListNode {
        public int val;
        public ListNode next;
        public ListNode(int val) { this.val = val; }
    }

    // 1. Detect if cycle exists
    public static boolean hasCycle(ListNode head) {
        ListNode slow = head, fast = head;
        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
            if (slow == fast) return true;
        }
        return false;
    }

    // 2. Find start node of cycle
    public static ListNode detectCycle(ListNode head) {
        ListNode slow = head, fast = head;
        boolean hasLoop = false;

        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
            if (slow == fast) {
                hasLoop = true;
                break;
            }
        }

        if (!hasLoop) return null;

        // Advance both 1 step from head and meeting point
        ListNode ptr1 = head;
        ListNode ptr2 = slow;
        while (ptr1 != ptr2) {
            ptr1 = ptr1.next;
            ptr2 = ptr2.next;
        }

        return ptr1; // Cycle entrance node
    }
}
```

---

### Go Implementation

```go
package main

// ListNode represents a node in a singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle returns true if the linked list contains a cycle
func HasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// DetectCycle returns the node where the cycle begins, or nil if no cycle exists
func DetectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	hasLoop := false

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			hasLoop = true
			break
		}
	}

	if !hasLoop {
		return nil
	}

	ptr1 := head
	ptr2 := slow

	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	return ptr1
}
```

---

## 5.4 - Time & Space Complexity Analysis

* **Time Complexity:** $\mathcal{O}(n)$ — If no cycle exists, fast pointer reaches the end in $n/2$ steps. If a cycle exists, once slow reaches the cycle, fast catches up in at most $C$ steps, where $C \le n$.
* **Space Complexity:** $\mathcal{O}(1)$ auxiliary space.

---

## 5.5 - Classic LeetCode Benchmarks

### 5.5.1 - Linked List Cycle (LeetCode #141)

#### 1. Problem Statement
Given `head`, the head of a linked list, determine if the linked list has a cycle in it.

#### 2. Solution Link
* [Go Implementation](problems/fast-and-slow-pointers/fast_slow_problems.go) (Function: `HasCycle`)
* [Java Implementation](fast-and-slow-pointers.md) (Method: `hasCycle`)

#### 3. Explanation
We initialize two pointers, `slow` and `fast`, both at the head. In each iteration, `slow` moves one step while `fast` moves two. If there is a cycle, the `fast` pointer will eventually "lap" the `slow` pointer, and they will meet. If `fast` reaches the end (`null`), no cycle exists.

#### 4. Conceptual Link to Fast & Slow Pointers
This is the foundational application of **Floyd's Tortoise and Hare**. It proves that two runners at different speeds on a circular track must eventually meet, allowing us to detect cycles in $\mathcal{O}(n)$ time without storing visited nodes.

### 5.5.2 - Linked List Cycle II (LeetCode #142)

#### 1. Problem Statement
Given the `head` of a linked list, return the node where the cycle begins. If there is no cycle, return `null`.

#### 2. Solution Link
* [Go Implementation](problems/fast-and-slow-pointers/fast_slow_problems.go) (Function: `DetectCycle`)
* [Java Implementation](fast-and-slow-pointers.md) (Method: `detectCycle`)

#### 3. Explanation
First, detect if a cycle exists using the standard Floyd's algorithm. Once the pointers meet, reset one pointer to the `head` of the list. Move both pointers one step at a time; the node where they meet again is the entrance to the cycle.

#### 4. Conceptual Link to Fast & Slow Pointers
Utilizes the **Mathematical Proof** that the distance from the head to the cycle start is equal to the distance from the meeting point to the cycle start. It demonstrates how different speeds can reveal structural properties of the list.

### 5.5.3 - Middle of the Linked List (LeetCode #876)

#### 1. Problem Statement
Given the `head` of a singly linked list, return the middle node of the linked list. If there are two middle nodes, return the second middle node.

#### 2. Solution Link
* [Go Implementation](problems/fast-and-slow-pointers/fast_slow_problems.go) (Function: `MiddleNode`)

#### 3. Explanation
Advance the `slow` pointer by one step and the `fast` pointer by two steps. When the `fast` pointer reaches the end of the list, the `slow` pointer will be exactly at the midpoint.

#### 4. Conceptual Link to Fast & Slow Pointers
Demonstrates the **Relative Speed** principle. By moving `fast` at twice the speed of `slow`, `slow` naturally covers half the distance by the time `fast` finishes, finding the middle in a single pass.

### 5.5.4 - Palindrome Linked List (LeetCode #234)

#### 1. Problem Statement
Given the `head` of a singly linked list, return `true` if it is a palindrome or `false` otherwise.

#### 2. Solution Link
* [Go Implementation](problems/fast-and-slow-pointers/fast_slow_problems.go) (Function: `IsPalindrome`)

#### 3. Explanation
1. Use fast/slow pointers to find the middle of the list.
2. Reverse the second half of the list in-place.
3. Compare the first half and the reversed second half node by node.
4. (Optional) Restore the list by reversing the second half again.

#### 4. Conceptual Link to Fast & Slow Pointers
Combines **Fast & Slow Pointers** (to find the midpoint) with **In-place Reversal**. It shows how this pattern can be a sub-component of more complex algorithms to achieve $\mathcal{O}(1)$ space complexity.

---

## 5.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 10: Elementary Data Structures (Linked lists pp. 256–264)
  * Exercises on Cycle Detection in Linked Data Structures
* https://en.wikipedia.org/wiki/Cycle_detection#Floyd's_tortoise_and_hare
* https://leetcode.com/problems/linked-list-cycle-ii/
