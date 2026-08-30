package two_heaps

import "container/heap"

// =============================================================
// MedianFinder — LeetCode #295
// Maintain a max-heap for the lower half and a min-heap for the
// upper half so that the median is always at the tops.
// =============================================================

// maxHeap is a max-heap of ints.
type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// minHeap is a min-heap of ints.
type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MedianFinder keeps track of a data stream's running median.
type MedianFinder struct {
	lo *maxHeap // lower half
	hi *minHeap // upper half
}

// NewMedianFinder constructs a MedianFinder.
func NewMedianFinder() *MedianFinder {
	lo := &maxHeap{}
	hi := &minHeap{}
	heap.Init(lo)
	heap.Init(hi)
	return &MedianFinder{lo: lo, hi: hi}
}

// AddNum adds a number to the data stream. O(log n).
func (mf *MedianFinder) AddNum(num int) {
	heap.Push(mf.lo, num)
	// Balance: move max of lo to hi
	heap.Push(mf.hi, heap.Pop(mf.lo))
	// Ensure lo has >= elements as hi
	if mf.lo.Len() < mf.hi.Len() {
		heap.Push(mf.lo, heap.Pop(mf.hi))
	}
}

// FindMedian returns the running median. O(1).
func (mf *MedianFinder) FindMedian() float64 {
	if mf.lo.Len() > mf.hi.Len() {
		return float64((*mf.lo)[0])
	}
	return float64((*mf.lo)[0]+(*mf.hi)[0]) / 2.0
}

// =============================================================
// MedianSlidingWindow — LeetCode #480
// Slide a window of size k across nums, returning the median
// at each position.  Uses two heaps + lazy deletion.
// =============================================================

// MedianSlidingWindow returns the median of each window of size k.
func MedianSlidingWindow(nums []int, k int) []float64 {
	n := len(nums)
	result := make([]float64, 0, n-k+1)

	// We use two heaps with lazy deletion tracked by a map.
	lo := &maxHeap{} // lower half (max-heap)
	hi := &minHeap{} // upper half (min-heap)
	heap.Init(lo)
	heap.Init(hi)
	invalid := make(map[int]int) // count of elements to lazily remove

	// Helper: balance so lo.Len() == hi.Len() or lo.Len() == hi.Len()+1
	balance := func(diff int) {
		// diff > 0 means lo has too many; diff < 0 means hi has too many
		if diff > 0 {
			heap.Push(hi, heap.Pop(lo))
		} else if diff < 0 {
			heap.Push(lo, heap.Pop(hi))
		}
	}

	// Prune lazy-deleted elements from the top of a heap.
	pruneMax := func() {
		for lo.Len() > 0 && invalid[(*lo)[0]] > 0 {
			invalid[(*lo)[0]]--
			heap.Pop(lo)
		}
	}
	pruneMin := func() {
		for hi.Len() > 0 && invalid[(*hi)[0]] > 0 {
			invalid[(*hi)[0]]--
			heap.Pop(hi)
		}
	}

	// Seed the first window.
	for i := 0; i < k; i++ {
		heap.Push(lo, nums[i])
	}
	for i := 0; i < k/2; i++ {
		heap.Push(hi, heap.Pop(lo))
	}

	getMedian := func() float64 {
		pruneMax()
		pruneMin()
		if k%2 == 1 {
			return float64((*lo)[0])
		}
		return float64((*lo)[0]+(*hi)[0]) / 2.0
	}

	result = append(result, getMedian())

	for i := k; i < n; i++ {
		outNum := nums[i-k]
		inNum := nums[i]

		// Determine which side outNum lives on to maintain balance count.
		pruneMax()
		pruneMin()
		loSize := lo.Len()
		_ = loSize

		// Add incoming number.
		if lo.Len() == 0 || inNum <= (*lo)[0] {
			heap.Push(lo, inNum)
		} else {
			heap.Push(hi, inNum)
		}

		// Mark outgoing number as invalid.
		invalid[outNum]++

		// Re-balance: one side may have grown or shrunk.
		pruneMax()
		pruneMin()

		// Determine effective sizes after pruning.
		loEff := lo.Len()
		hiEff := hi.Len()
		diff := loEff - hiEff
		if diff > 1 {
			balance(1)
		} else if diff < 0 {
			balance(-1)
		}

		result = append(result, getMedian())
	}

	return result
}

// =============================================================
// FindMaximizedCapital — LeetCode #502 (IPO)
// Greedily pick the most profitable available project k times.
// Use a min-heap on capital to release projects as capital grows,
// then a max-heap on profit to always pick the best.
// =============================================================

// capProfit is an entry in the capital min-heap.
type capProfit struct{ cap, profit int }
type capHeap []capProfit

func (h capHeap) Len() int           { return len(h) }
func (h capHeap) Less(i, j int) bool { return h[i].cap < h[j].cap }
func (h capHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *capHeap) Push(x interface{}) { *h = append(*h, x.(capProfit)) }
func (h *capHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// FindMaximizedCapital selects at most k projects to maximise capital.
// Time: O(n log n). Ref: LeetCode #502.
func FindMaximizedCapital(k, w int, profits, capital []int) int {
	n := len(profits)
	capQ := &capHeap{}
	for i := 0; i < n; i++ {
		heap.Push(capQ, capProfit{capital[i], profits[i]})
	}

	profQ := &maxHeap{}
	heap.Init(profQ)

	for i := 0; i < k; i++ {
		// Unlock all projects we can afford.
		for capQ.Len() > 0 && (*capQ)[0].cap <= w {
			cp := heap.Pop(capQ).(capProfit)
			heap.Push(profQ, cp.profit)
		}
		if profQ.Len() == 0 {
			break // no affordable project
		}
		w += heap.Pop(profQ).(int)
	}
	return w
}
