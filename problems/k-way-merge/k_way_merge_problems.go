package k_way_merge

import "container/heap"

// =============================================================
// ListNode — singly linked-list node shared across problems.
// =============================================================

type ListNode struct {
	Val  int
	Next *ListNode
}

// =============================================================
// MergeKLists — LeetCode #23
// Use a min-heap of size at most k to merge k sorted linked lists.
// Time: O(N log k), Space: O(k).
// =============================================================

// listNodeHeap is a min-heap of ListNode pointers.
type listNodeHeap []*ListNode

func (h listNodeHeap) Len() int           { return len(h) }
func (h listNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h listNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *listNodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *listNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MergeKLists merges k sorted linked lists into one sorted list.
func MergeKLists(lists []*ListNode) *ListNode {
	h := &listNodeHeap{}
	heap.Init(h)
	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}
	dummy := &ListNode{}
	cur := dummy
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		cur.Next = node
		cur = cur.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return dummy.Next
}

// =============================================================
// KthSmallest — LeetCode #378
// Use a min-heap seeded with the first column; each pop expands
// to the next element in the same row.
// Time: O(k log k), Space: O(k).
// =============================================================

type cell struct{ val, row, col int }
type cellHeap []cell

func (h cellHeap) Len() int           { return len(h) }
func (h cellHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h cellHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *cellHeap) Push(x interface{}) { *h = append(*h, x.(cell)) }
func (h *cellHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// KthSmallest returns the k-th smallest element in an n×n sorted matrix.
func KthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	h := &cellHeap{}
	for r := 0; r < n; r++ {
		heap.Push(h, cell{matrix[r][0], r, 0})
	}
	var val int
	for i := 0; i < k; i++ {
		c := heap.Pop(h).(cell)
		val = c.val
		if c.col+1 < n {
			heap.Push(h, cell{matrix[c.row][c.col+1], c.row, c.col + 1})
		}
	}
	return val
}

// =============================================================
// SmallestRange — LeetCode #632
// Use a min-heap across all lists; track the global max to form
// the current range [min, maxSoFar].  Advance the min-pointer
// until any list is exhausted.
// Time: O(N log k), Space: O(k).
// =============================================================

type rangeCell struct{ val, row, col int }
type rangeCellHeap []rangeCell

func (h rangeCellHeap) Len() int           { return len(h) }
func (h rangeCellHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h rangeCellHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *rangeCellHeap) Push(x interface{}) { *h = append(*h, x.(rangeCell)) }
func (h *rangeCellHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// SmallestRange finds the smallest range that includes at least one element
// from each of the k sorted lists.
func SmallestRange(nums [][]int) [2]int {
	h := &rangeCellHeap{}
	maxVal := nums[0][0]
	for r, row := range nums {
		heap.Push(h, rangeCell{row[0], r, 0})
		if row[0] > maxVal {
			maxVal = row[0]
		}
	}

	best := [2]int{(*h)[0].val, maxVal}
	for {
		c := heap.Pop(h).(rangeCell)
		if c.col+1 >= len(nums[c.row]) {
			break // one list is exhausted
		}
		next := nums[c.row][c.col+1]
		if next > maxVal {
			maxVal = next
		}
		heap.Push(h, rangeCell{next, c.row, c.col + 1})
		lo := (*h)[0].val
		if maxVal-lo < best[1]-best[0] {
			best = [2]int{lo, maxVal}
		}
	}
	return best
}
