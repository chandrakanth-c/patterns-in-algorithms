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

### 8.5.1 - Reverse Linked List (LeetCode #206)

#### 1. Problem Statement
Given the `head` of a singly linked list, reverse the list, and return the reversed list.

#### 2. Solution Link
* [Go Implementation](problems/in-place-reversal-linked-list/reversal_ll_problems.go) (Function: `ReverseList`)

#### 3. Explanation
We use three pointers: `prev`, `curr`, and `next`. We iterate through the list, and for each node, we save its `next` node, point its `next` to `prev`, and then move `prev` and `curr` one step forward.

#### 4. Conceptual Link to In-place Reversal
This is the **atomic operation** of the pattern. It demonstrates how to redirect pointers in a single pass while maintaining a reference to the remaining list, achieving $\mathcal{O}(n)$ time and $\mathcal{O}(1)$ space.

### 8.5.2 - Reverse Linked List II (LeetCode #92)

#### 1. Problem Statement
Given the `head` of a singly linked list and two integers `left` and `right` where `left <= right`, reverse the nodes of the list from position `left` to position `right`, and return the reversed list.

#### 2. Solution Link
* [Go Implementation](problems/in-place-reversal-linked-list/reversal_ll_problems.go) (Function: `ReverseBetween`)
* [Java Implementation](in-place-reversal-linked-list.md) (Method: `reverseBetween`)

#### 3. Explanation
1. Traverse to the $(left-1)$-th node.
2. Reverse the sublist from `left` to `right` using the iterative approach.
3. Reconnect the reversed sublist with the rest of the list.

#### 4. Conceptual Link to In-place Reversal
Demonstrates the **Sublist Reversal** technique. It highlights the importance of keeping track of the nodes just before and after the sublist to maintain list integrity after the in-place operation.

### 8.5.3 - Reverse Nodes in k-Group (LeetCode #25)

#### 1. Problem Statement
Given the `head` of a linked list, reverse the nodes of the list `k` at a time, and return the modified list. If the number of nodes is not a multiple of `k` left out, in the end, it should remain as it is.

#### 2. Solution Link
* [Go Implementation](problems/in-place-reversal-linked-list/reversal_ll_problems.go) (Function: `ReverseKGroup`)

#### 3. Explanation
This problem is solved by repeatedly applying sublist reversal. We first check if there are at least `k` nodes available. If so, we reverse them and then recursively call the function for the remaining list, connecting the result to the tail of the current reversed group.

#### 4. Conceptual Link to In-place Reversal
Extends the pattern to **Recursive/Iterative Grouping**. It shows how the base reversal logic can be composed to solve complex structural transformations on linked data.

### 8.5.4 - Reorder List (LeetCode #143)

#### 1. Problem Statement
You are given the head of a singly linked-list. The list can be represented as: $L_0 \rightarrow L_1 \rightarrow \dots \rightarrow L_{n-1} \rightarrow L_n$. Reorder the list to be on the following form: $L_0 \rightarrow L_n \rightarrow L_1 \rightarrow L_{n-1} \rightarrow L_2 \rightarrow L_{n-2} \rightarrow \dots$

#### 2. Solution Link
* [Go Implementation](problems/in-place-reversal-linked-list/reversal_ll_problems.go) (Function: `ReorderList`)

#### 3. Explanation
1. Find the middle of the list using Fast & Slow pointers.
2. Reverse the second half of the list in-place.
3. Merge the two halves by alternating nodes from each.

#### 4. Conceptual Link to In-place Reversal
Demonstrates how **In-place Reversal** is a critical building block for complex list reordering. By reversing the second half, we can access nodes from the "end" of the list in forward order, allowing for a linear-time interleaving merge.

---

## 8.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 10: Elementary Data Structures (Linked Lists pp. 256–264)
* https://leetcode.com/problems/reverse-linked-list-ii/
* https://techinterviewhandbook.org/algorithms/linked-list/
