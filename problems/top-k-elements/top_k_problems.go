package top_k

import (
	"container/heap"
	"sort"
)

// ============================================================
// Min-heap of ints (used by FindKthLargest and KClosest)
// ============================================================

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// --- FindKthLargest (LeetCode #215) ---
// Maintains a min-heap of size k; the root is the kth largest element.
// Time: O(n log k), Space: O(k)
func FindKthLargest(nums []int, k int) int {
	h := &minHeap{}
	heap.Init(h)
	for _, n := range nums {
		heap.Push(h, n)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

// ============================================================
// Entry for TopKFrequent — (count, value) max-heap
// ============================================================

type freqEntry struct{ val, cnt int }
type freqHeap []freqEntry

func (h freqHeap) Len() int            { return len(h) }
func (h freqHeap) Less(i, j int) bool  { return h[i].cnt > h[j].cnt } // max-heap by count
func (h freqHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *freqHeap) Push(x interface{}) { *h = append(*h, x.(freqEntry)) }
func (h *freqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// --- TopKFrequent (LeetCode #347) ---
// Returns the k most frequent elements using a frequency max-heap.
// Time: O(n log k), Space: O(n)
func TopKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}
	h := &freqHeap{}
	heap.Init(h)
	for val, cnt := range freq {
		heap.Push(h, freqEntry{val, cnt})
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(h).(freqEntry).val
	}
	return result
}

// ============================================================
// KClosest — sort by squared Euclidean distance
// ============================================================

// --- KClosest (LeetCode #973) ---
// Returns the k closest points to the origin using sort.
// Time: O(n log n), Space: O(1) extra
func KClosest(points [][2]int, k int) [][2]int {
	dist := func(p [2]int) int { return p[0]*p[0] + p[1]*p[1] }
	sort.Slice(points, func(i, j int) bool {
		return dist(points[i]) < dist(points[j])
	})
	return points[:k]
}

// ============================================================
// Max-heap of (count, char) pairs for ReorganizeString
// ============================================================

type charEntry struct{ ch byte; cnt int }
type charHeap []charEntry

func (h charHeap) Len() int            { return len(h) }
func (h charHeap) Less(i, j int) bool  { return h[i].cnt > h[j].cnt }
func (h charHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *charHeap) Push(x interface{}) { *h = append(*h, x.(charEntry)) }
func (h *charHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// --- ReorganizeString (LeetCode #767) ---
// Rearranges s so no two adjacent characters are equal.
// Greedily places the two most-frequent characters at each step.
// Returns "" if impossible (any char's count > ceil(n/2)).
// Time: O(n log 26) = O(n), Space: O(26) = O(1)
func ReorganizeString(s string) string {
	freq := [26]int{}
	for _, ch := range s {
		freq[ch-'a']++
	}
	h := &charHeap{}
	heap.Init(h)
	for i, cnt := range freq {
		if cnt > 0 {
			heap.Push(h, charEntry{byte('a' + i), cnt})
		}
	}
	result := []byte{}
	for h.Len() >= 2 {
		first := heap.Pop(h).(charEntry)
		second := heap.Pop(h).(charEntry)
		result = append(result, first.ch, second.ch)
		first.cnt--
		second.cnt--
		if first.cnt > 0 {
			heap.Push(h, first)
		}
		if second.cnt > 0 {
			heap.Push(h, second)
		}
	}
	if h.Len() == 1 {
		last := heap.Pop(h).(charEntry)
		if last.cnt > 1 {
			return "" // impossible
		}
		result = append(result, last.ch)
	}
	return string(result)
}
