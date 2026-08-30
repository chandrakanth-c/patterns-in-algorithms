package reversal_ll

// ListNode is a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// --- ReverseList (LeetCode #206) ---
// Iteratively reverses a linked list in place.
// Time: O(n), Space: O(1)
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// --- ReverseBetween (LeetCode #92) ---
// Reverses the sub-list from position left to right (1-indexed).
// Uses a dummy head to simplify edge cases.
// Time: O(n), Space: O(1)
func ReverseBetween(head *ListNode, left, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	// Advance pre to the node just before position left.
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	curr := pre.Next
	// Perform (right-left) in-place reversals using the "insert after pre" technique.
	for i := 0; i < right-left; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}

// --- ReverseKGroup (LeetCode #25) ---
// Reverses nodes in k-sized groups; leaves the last group as-is if smaller than k.
// Time: O(n), Space: O(1)
func ReverseKGroup(head *ListNode, k int) *ListNode {
	// Check if there are at least k nodes remaining.
	count := 0
	node := head
	for node != nil && count < k {
		node = node.Next
		count++
	}
	if count < k {
		return head // fewer than k nodes left — don't reverse
	}
	// Reverse k nodes.
	var prev *ListNode
	curr := head
	for i := 0; i < k; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	// head is now the tail of the reversed group; recurse for remainder.
	head.Next = ReverseKGroup(curr, k)
	return prev
}

// --- ReorderList (LeetCode #143) ---
// Reorders list in-place: L0→L1→…→Ln-1→Ln  becomes  L0→Ln→L1→Ln-1→…
// Strategy: find mid, reverse second half, merge two halves.
// Time: O(n), Space: O(1)
func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// Step 1: Find the middle using slow/fast pointers.
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// Step 2: Reverse the second half.
	secondHalf := ReverseList(slow.Next)
	slow.Next = nil // cut the list

	// Step 3: Merge the two halves.
	first, second := head, secondHalf
	for second != nil {
		tmp1 := first.Next
		tmp2 := second.Next
		first.Next = second
		second.Next = tmp1
		first = tmp1
		second = tmp2
	}
}
