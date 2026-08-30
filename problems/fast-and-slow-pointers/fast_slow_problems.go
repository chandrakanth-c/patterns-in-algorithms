package fast_slow

// ListNode is a singly-linked list node used by fast & slow pointer problems.
type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle detects whether a linked list contains a cycle (LeetCode #141).
// Floyd's tortoise-and-hare: slow moves 1 step, fast moves 2 steps.
// Time: O(n), Space: O(1)
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

// DetectCycle returns the node where the cycle begins, or nil if no cycle (LeetCode #142).
// After meeting point, reset one pointer to head; advance both 1 step until they meet.
// Time: O(n), Space: O(1)
func DetectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// Find entry point
			entry := head
			for entry != slow {
				entry = entry.Next
				slow = slow.Next
			}
			return entry
		}
	}
	return nil
}

// MiddleNode returns the middle node of a linked list (LeetCode #876).
// When two middle nodes exist, returns the second one.
// Time: O(n), Space: O(1)
func MiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// IsPalindrome checks whether the values in the linked list form a palindrome (LeetCode #234).
// Find middle, reverse second half, compare, then restore.
// Time: O(n), Space: O(1)
func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// Find middle
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// Reverse second half
	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	// Compare
	left, right := head, prev
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

// FindDuplicate finds the duplicate number in nums[1..n] using Floyd's cycle detection
// (LeetCode #287 — cyclic sort / fast-slow pointer version).
// Treat array values as next-pointers; the duplicate creates the cycle entrance.
// Time: O(n), Space: O(1)
func FindDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	// Find entrance
	entry := 0
	for entry != slow {
		entry = nums[entry]
		slow = nums[slow]
	}
	return entry
}
