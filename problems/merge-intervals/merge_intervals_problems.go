package merge_intervals

import (
	"container/heap"
	"sort"
)

// ActivitySelection returns the maximum number of non-overlapping activities (CLRS 16.1).
// Greedy: always pick the activity that finishes earliest.
// Precondition: finish[] is sorted ascending (caller must ensure this).
// Time: O(n log n) if sorting needed; O(n) if pre-sorted.
func ActivitySelection(start, finish []int) int {
	n := len(start)
	if n == 0 {
		return 0
	}
	// Sort by finish time
	type activity struct{ s, f int }
	acts := make([]activity, n)
	for i := range acts {
		acts[i] = activity{start[i], finish[i]}
	}
	sort.Slice(acts, func(i, j int) bool { return acts[i].f < acts[j].f })

	count := 1
	lastFinish := acts[0].f
	for i := 1; i < n; i++ {
		if acts[i].s >= lastFinish {
			count++
			lastFinish = acts[i].f
		}
	}
	return count
}

// Merge merges all overlapping intervals and returns the result (LeetCode #56).
// Sort by start time, then fold overlapping intervals.
// Time: O(n log n), Space: O(n)
func Merge(intervals [][2]int) [][2]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][2]int{intervals[0]}
	for _, iv := range intervals[1:] {
		last := &result[len(result)-1]
		if iv[0] <= last[1] {
			// Overlapping — extend end if needed
			if iv[1] > last[1] {
				last[1] = iv[1]
			}
		} else {
			result = append(result, iv)
		}
	}
	return result
}

// Insert inserts newInterval into a sorted, non-overlapping list of intervals and
// merges as necessary (LeetCode #57).
// Time: O(n), Space: O(n)
func Insert(intervals [][2]int, newInterval [2]int) [][2]int {
	result := [][2]int{}
	i, n := 0, len(intervals)
	// Add all intervals that end before newInterval starts
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}
	// Merge overlapping intervals with newInterval
	for i < n && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	result = append(result, newInterval)
	// Add remaining intervals
	for i < n {
		result = append(result, intervals[i])
		i++
	}
	return result
}

// EraseOverlapIntervals returns the minimum number of intervals to remove so that
// the rest are non-overlapping (LeetCode #435).
// Greedy: sort by end time, greedily keep intervals that don't overlap previous.
// Time: O(n log n), Space: O(1)
func EraseOverlapIntervals(intervals [][2]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	removed := 0
	prevEnd := intervals[0][1]
	for _, iv := range intervals[1:] {
		if iv[0] < prevEnd {
			// Overlap: remove current interval
			removed++
		} else {
			prevEnd = iv[1]
		}
	}
	return removed
}

// minHeap implements heap.Interface for int slices (used by MinMeetingRooms).
type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// MinMeetingRooms returns the minimum number of conference rooms required (LeetCode #253).
// Sort by start time, use a min-heap of end times to track room availability.
// Time: O(n log n), Space: O(n)
func MinMeetingRooms(intervals [][2]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	endTimes := &minHeap{}
	heap.Init(endTimes)
	for _, iv := range intervals {
		if endTimes.Len() > 0 && (*endTimes)[0] <= iv[0] {
			heap.Pop(endTimes) // Reuse the room
		}
		heap.Push(endTimes, iv[1])
	}
	return endTimes.Len()
}
