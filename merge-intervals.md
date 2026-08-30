# 6 - Merge Intervals

## 6.1 - Overview & Theoretical Foundations (CLRS Chapter 15)

* The **Merge Intervals** pattern solves problems dealing with overlapping ranges $[start, end]$ and interval scheduling.
* Grounded in **Greedy Algorithms** (CLRS Chapter 15: An activity-selection problem), sorting intervals by start times allows a single greedy scan to determine whether consecutive intervals overlap and merge them immediately.
* **Greedy Choice Property:** When sorted by start time, interval $I_{i+1}$ can only overlap with the immediately preceding active merged interval $I_{\text{curr}}$. If $I_{i+1}.\text{start} \le I_{\text{curr}}.\text{end}$, merging them produces the locally and globally optimal merged interval $[I_{\text{curr}}.\text{start}, \max(I_{\text{curr}}.\text{end}, I_{i+1}.\text{end})]$.

---

## 6.2 - Properties of a problem that suggests Merge Intervals

* Problem deals with **intervals**, **ranges**, **scheduling time blocks**, or **resource allocation**.
* Identifying mutual exclusions, non-overlapping subsets, or finding maximum concurrent overlaps (meeting rooms).

---

## 6.3 - Classic Example: Merge Overlapping Intervals

### Java Implementation

```java
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class MergeIntervals {

    public static int[][] merge(int[][] intervals) {
        if (intervals.length <= 1) return intervals;

        // Step 1: Sort by start times (Greedy ordering)
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));

        List<int[]> result = new ArrayList<>();
        int[] currentInterval = intervals[0];
        result.add(currentInterval);

        for (int i = 1; i < intervals.length; i++) {
            int currentEnd = currentInterval[1];
            int nextStart = intervals[i][0];
            int nextEnd = intervals[i][1];

            if (nextStart <= currentEnd) {
                // Overlap: expand current interval end
                currentInterval[1] = Math.max(currentEnd, nextEnd);
            } else {
                // Disjoint: start a new interval
                currentInterval = intervals[i];
                result.add(currentInterval);
            }
        }

        return result.toArray(new int[result.size()][]);
    }
}
```

---

### Go Implementation

```go
package main

import (
	"sort"
)

// Merge takes a list of intervals and merges all overlapping intervals
func Merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// Step 1: Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		lastIdx := len(result) - 1
		currentEnd := result[lastIdx][1]
		nextStart := intervals[i][0]
		nextEnd := intervals[i][1]

		if nextStart <= currentEnd {
			// Overlap: merge intervals
			if nextEnd > currentEnd {
				result[lastIdx][1] = nextEnd
			}
		} else {
			// Disjoint: append new interval
			result = append(result, intervals[i])
		}
	}

	return result
}
```

---

## 6.4 - Time & Space Complexity Analysis

* **Time Complexity:** $\mathcal{O}(n \log n)$ dominated by sorting the $n$ intervals. The subsequent linear scan is $\Theta(n)$.
* **Space Complexity:** $\mathcal{O}(n)$ to hold the merged output and sorting stack space.

---

## 6.5 - Classic LeetCode & CLRS Benchmarks

### 1. Merge Intervals (LeetCode #56)
* **Problem Statement**: Given an array of `intervals` where `intervals[i] = [starti, endi]`, merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
* **Solution Link**: [problems/merge-intervals/merge_intervals_problems.go](problems/merge-intervals/merge_intervals_problems.go) (`Merge`)
* **Explanation**: First, sort the intervals by their start times. Then, iterate through the sorted intervals, merging the current interval with the last merged interval if they overlap.
* **Conceptual Link**: A foundational application of the **Greedy Choice Property**, where sorting ensures that we only need to compare the current interval with the most recently merged one.

### 2. Insert Interval (LeetCode #57)
* **Problem Statement**: You are given an array of non-overlapping intervals `intervals` sorted by their start times. Insert a `newInterval` into the list such that it remains sorted and non-overlapping, merging intervals if necessary.
* **Solution Link**: [problems/merge-intervals/merge_intervals_problems.go](problems/merge-intervals/merge_intervals_problems.go) (`Insert`)
* **Explanation**: The problem is solved in three parts: add all intervals that end before the new interval starts, merge all overlapping intervals with the new interval, and finally add all remaining intervals.
* **Conceptual Link**: Demonstrates how to maintain interval properties (sorted, non-overlapping) during a dynamic update operation in $\mathcal{O}(n)$ time.

### 3. Non-overlapping Intervals (LeetCode #435)
* **Problem Statement**: Given an array of intervals `intervals`, return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.
* **Solution Link**: [problems/merge-intervals/merge_intervals_problems.go](problems/merge-intervals/merge_intervals_problems.go) (`EraseOverlapIntervals`)
* **Explanation**: This is equivalent to finding the maximum number of non-overlapping intervals. By sorting intervals by their end times and greedily picking the ones that finish earliest, we maximize the space for remaining intervals.
* **Conceptual Link**: Directly maps to the **Activity-Selection Problem** from CLRS, illustrating the greedy strategy for maximizing interval counts.

### 4. Meeting Rooms II (LeetCode #253)
* **Problem Statement**: Given an array of meeting time intervals consisting of start and end times `[[s1,e1],[s2,e2],...]`, find the minimum number of conference rooms required.
* **Solution Link**: [problems/merge-intervals/merge_intervals_problems.go](problems/merge-intervals/merge_intervals_problems.go) (`MinMeetingRooms`)
* **Explanation**: Sort intervals by start time and use a min-heap to keep track of the end times of meetings currently in progress. If a new meeting starts after the earliest end time in the heap, we can reuse that room.
* **Conceptual Link**: Uses a **Min-Heap** to manage resources efficiently, a common technique for solving "maximum concurrent overlap" problems.

* **Activity-Selection Problem** (CLRS 15.1)
* **Meeting Rooms I** (LeetCode #252)
* **Interval List Intersections** (LeetCode #986)

---

## 6.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 15: Greedy Algorithms (An activity-selection problem pp. 428–436)
* https://leetcode.com/problems/merge-intervals/
* https://techinterviewhandbook.org/algorithms/interval/
