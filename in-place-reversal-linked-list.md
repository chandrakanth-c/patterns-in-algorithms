# 8 - In-place Reversal of a Linked List

## 8.1 - Overview & Theoretical Foundations (CLRS Chapter 10)

* In a singly linked list (CLRS 10.2), each node stores an element `key` and a pointer `next`.
* Reversing the list without allocating auxiliary nodes requires iterative pointer redirection using three tracking pointers:
  * `prev`: The head of the already reversed prefix.
  * `curr`: The node currently being redirected.
  * `next`: Temporary store of `curr.next` to maintain access to the unprocessed suffix.
* **Loop Invariant:** At the start of each iteration, the sublist from original head to `prev` is completely reversed, and `curr` points to the head of the remaining forward sublist.

---

## 8.2 - Properties of a problem that suggests In-place Reversal

* Need to reverse an entire linked list, a sublist between positions $L$ and $R$, or nodes in groups of $k$.
* Strict $\mathcal{O}(1)$ auxiliary space constraint.

---

## 8.3 - Classic Example: Reverse Sublist Between Left and Right

### Java Implementation

```java
public class InPlaceLinkedListReversal {

    public static class ListNode {
        public int val;
        public ListNode next;
        public ListNode(int val) { this.val = val; }
    }

    public static ListNode reverseBetween(ListNode head, int left, int right) {
        if (head == null || left == right) return head;

        ListNode dummy = new ListNode(0);
        dummy.next = head;
        ListNode prev = dummy;

        // Step 1: Reach the node just before 'left'
        for (int i = 0; i < left - 1; i++) {
            prev = prev.next;
        }

        // Step 2: Reverse sublist between left and right
        ListNode curr = prev.next;
        ListNode sublistTail = curr;
        ListNode sublistPrev = null;

        for (int i = 0; i < right - left + 1; i++) {
            ListNode nextTemp = curr.next;
            curr.next = sublistPrev;
            sublistPrev = curr;
            curr = nextTemp;
        }

        // Step 3: Reconnect reversed portion with boundary nodes
        prev.next = sublistPrev;
        sublistTail.next = curr;

        return dummy.next;
    }
}
```

---

### Go Implementation

```go
package main

// ListNode represents a singly linked list node
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseBetween reverses the nodes of the list from position left to position right
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy

	// Step 1: Advance prev to node at position left-1
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	// Step 2: In-place reversal of sublist
	curr := prev.Next
	sublistTail := curr
	var sublistPrev *ListNode = nil

	for i := 0; i < right-left+1; i++ {
		nextTemp := curr.Next
		curr.Next = sublistPrev
		sublistPrev = curr
		curr = nextTemp
	}

	// Step 3: Reconnect
	prev.Next = sublistPrev
	sublistTail.Next = curr

	return dummy.Next
}
```

---

## 8.4 - Time & Space Complexity Analysis

* **Time Complexity:** $\mathcal{O}(n)$ — Exactly $n$ pointer redirections performed in a single pass.
* **Space Complexity:** $\mathcal{O}(1)$ auxiliary space.

---

## 8.5 - Classic LeetCode Benchmarks

* **Reverse Linked List** (LeetCode #206)
* **Reverse Linked List II** (LeetCode #92)
* **Reverse Nodes in k-Group** (LeetCode #25)
* **Reorder List** (LeetCode #143)
* **Palindrome Linked List** (LeetCode #234)

---

## 8.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 10: Elementary Data Structures (Linked Lists pp. 256–264)
* https://leetcode.com/problems/reverse-linked-list-ii/
* https://techinterviewhandbook.org/algorithms/linked-list/
