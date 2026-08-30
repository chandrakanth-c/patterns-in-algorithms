# 6 - Merge Intervals

## 6.1 - Overview

* The **Merge Intervals** pattern describes an efficient technique to deal with overlapping intervals.
* In a lot of interval problems, you either need to find overlapping intervals or merge intervals if they overlap.
* Given two intervals $A = [start_A, end_A]$ and $B = [start_B, end_B]$, they overlap if and only if:
  $$\max(start_A, start_B) \le \min(end_A, end_B)$$
  or when sorted by start times: $start_B \le end_A$.

---

## 6.2 - Properties of a problem that suggests Merge Intervals

* The problem involves **ranges**, **intervals**, **start/end times**, or **scheduling events**.
* You need to identify overlaps, combine overlapping time blocks, or find free time slots.

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

        // Step 1: Sort intervals by starting time
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));

        List<int[]> merged = new ArrayList<>();
        int[] currentInterval = intervals[0];
        merged.add(currentInterval);

        for (int[] interval : intervals) {
            int currentEnd = currentInterval[1];
            int nextStart = interval[0];
            int nextEnd = interval[1];

            if (nextStart <= currentEnd) {
                // Overlap exists: merge intervals by taking max end time
                currentInterval[1] = Math.max(currentEnd, nextEnd);
            } else {
                // No overlap: add new interval to list
                currentInterval = interval;
                merged.add(currentInterval);
            }
        }

        return merged.toArray(new int[merged.size()][]);
    }
}
```

---

## 6.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(n \log n)$ due to the initial sorting step. The subsequent single-pass linear scan takes $\mathcal{O}(n)$.
* **Space Complexity:** $\mathcal{O}(n)$ to store the output merged intervals (and sorting stack space).

---

## 6.5 - Classic LeetCode Problems

* **Merge Intervals** (LeetCode #56)
* **Insert Interval** (LeetCode #57)
* **Non-overlapping Intervals** (LeetCode #435)
* **Meeting Rooms I & II** (LeetCode #252, #253)
* **Interval List Intersections** (LeetCode #986)
* **Employee Free Time** (LeetCode #759)

---

## 6.6 - Sources used for this file:
https://leetcode.com/problems/merge-intervals/ <br>
https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d3fe08d2bb2d978e2025 <br>
https://www.geeksforgeeks.org/merging-intervals/ <br>
https://techinterviewhandbook.org/algorithms/interval/
