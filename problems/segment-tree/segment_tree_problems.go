package segment_tree

import "sort"

// --- Segment Tree for Range Sum Query with Point Update (LeetCode #307) ---

type NumArray struct {
	n    int
	tree []int
}

func NewNumArray(nums []int) *NumArray {
	n := len(nums)
	tree := make([]int, 4*n)
	na := &NumArray{n: n, tree: tree}
	na.build(nums, 0, 0, n-1)
	return na
}

func (na *NumArray) build(nums []int, node, start, end int) {
	if start == end {
		na.tree[node] = nums[start]
		return
	}
	mid := start + (end-start)/2
	na.build(nums, 2*node+1, start, mid)
	na.build(nums, 2*node+2, mid+1, end)
	na.tree[node] = na.tree[2*node+1] + na.tree[2*node+2]
}

func (na *NumArray) Update(index, val int) {
	na.updateHelper(0, 0, na.n-1, index, val)
}

func (na *NumArray) updateHelper(node, start, end, index, val int) {
	if start == end {
		na.tree[node] = val
		return
	}
	mid := start + (end-start)/2
	if index <= mid {
		na.updateHelper(2*node+1, start, mid, index, val)
	} else {
		na.updateHelper(2*node+2, mid+1, end, index, val)
	}
	na.tree[node] = na.tree[2*node+1] + na.tree[2*node+2]
}

// SumRange returns the sum of nums[left..right] inclusive.
func (na *NumArray) SumRange(left, right int) int {
	return na.queryHelper(0, 0, na.n-1, left, right)
}

func (na *NumArray) queryHelper(node, start, end, l, r int) int {
	if r < start || end < l {
		return 0
	}
	if l <= start && end <= r {
		return na.tree[node]
	}
	mid := start + (end-start)/2
	return na.queryHelper(2*node+1, start, mid, l, r) +
		na.queryHelper(2*node+2, mid+1, end, l, r)
}

// --- CountSmaller using Binary Indexed Tree / Merge Sort (LeetCode #315) ---
// Uses coordinate compression + BIT (Fenwick tree).

func CountSmaller(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	// Coordinate compress
	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Ints(sorted)
	// Deduplicate
	unique := []int{sorted[0]}
	for i := 1; i < len(sorted); i++ {
		if sorted[i] != sorted[i-1] {
			unique = append(unique, sorted[i])
		}
	}
	rank := func(v int) int {
		lo, hi := 0, len(unique)-1
		for lo <= hi {
			mid := (lo + hi) / 2
			if unique[mid] == v {
				return mid + 1 // 1-indexed
			} else if unique[mid] < v {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		return -1
	}
	// BIT
	bit := make([]int, len(unique)+1)
	bitUpdate := func(i int) {
		for ; i < len(bit); i += i & (-i) {
			bit[i]++
		}
	}
	bitQuery := func(i int) int {
		s := 0
		for ; i > 0; i -= i & (-i) {
			s += bit[i]
		}
		return s
	}
	// Traverse from right to left
	for i := n - 1; i >= 0; i-- {
		r := rank(nums[i])
		if r > 1 {
			result[i] = bitQuery(r - 1)
		}
		bitUpdate(r)
	}
	return result
}

// --- Skyline Problem (LeetCode #218) ---
// Uses a sorted event sweep with a max-heap via sorted slice.

func GetSkyline(buildings [][3]int) [][2]int {
	// Collect events: (x, height) — negative height = building start, positive = building end
	events := [][2]int{}
	for _, b := range buildings {
		events = append(events, [2]int{b[0], -b[2]}) // start
		events = append(events, [2]int{b[1], b[2]})  // end
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1]
	})

	// Use a sorted multiset simulation via a slice + sort (acceptable for test sizes)
	// heights is a sorted descending list of active building heights; 0 is always present
	heights := []int{0}
	prevMaxH := 0
	result := [][2]int{}

	addHeight := func(h int) {
		// Insert maintaining sorted order (ascending)
		pos := sort.SearchInts(heights, h)
		heights = append(heights, 0)
		copy(heights[pos+1:], heights[pos:])
		heights[pos] = h
	}
	removeHeight := func(h int) {
		pos := sort.SearchInts(heights, h)
		heights = append(heights[:pos], heights[pos+1:]...)
	}

	for _, e := range events {
		x, h := e[0], e[1]
		if h < 0 {
			addHeight(-h)
		} else {
			removeHeight(h)
		}
		curMax := heights[len(heights)-1]
		if curMax != prevMaxH {
			result = append(result, [2]int{x, curMax})
			prevMaxH = curMax
		}
	}
	return result
}
