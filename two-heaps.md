# 13 - Two Heaps (Median of a Stream)

## 13.1 - Overview

* The **Two Heaps** pattern uses two priority queues to divide an unsorted dataset into two balanced halves:
  1. **Max-Heap:** Stores the smaller half of numbers (the root is the maximum of the small half).
  2. **Min-Heap:** Stores the larger half of numbers (the root is the minimum of the large half).
* By maintaining size balance ($\text{maxHeap.size} == \text{minHeap.size}$ or $\text{maxHeap.size} == \text{minHeap.size} + 1$), the median is always available in $\mathcal{O}(1)$ time.

---

## 13.2 - Properties of a problem that suggests Two Heaps

* You are given a continuous stream of data and need to calculate the **median** dynamically.
* You need to track the smallest and largest elements simultaneously in fluctuating datasets.
* Sliding window median or dynamic percentile calculations.

---

## 13.3 - Classic Example: Find Median from Data Stream

### Java Implementation

```java
import java.util.Collections;
import java.util.PriorityQueue;

public class MedianFinder {
    private PriorityQueue<Integer> maxHeap; // stores smaller half
    private PriorityQueue<Integer> minHeap; // stores larger half

    public MedianFinder() {
        maxHeap = new PriorityQueue<>(Collections.reverseOrder());
        minHeap = new PriorityQueue<>();
    }

    public void addNum(int num) {
        if (maxHeap.isEmpty() || num <= maxHeap.peek()) {
            maxHeap.offer(num);
        } else {
            minHeap.offer(num);
        }

        // Rebalance heaps: maxHeap can have at most 1 more element than minHeap
        if (maxHeap.size() > minHeap.size() + 1) {
            minHeap.offer(maxHeap.poll());
        } else if (minHeap.size() > maxHeap.size()) {
            maxHeap.offer(minHeap.poll());
        }
    }

    public double findMedian() {
        if (maxHeap.size() == minHeap.size()) {
            return (maxHeap.peek() + minHeap.peek()) / 2.0;
        }
        return maxHeap.peek();
    }
}
```

---

## 13.4 - Time & Space Complexity

* **`addNum()`:** $\mathcal{O}(\log n)$ because insertion and rebalancing perform heap push/pop operations.
* **`findMedian()`:** $\mathcal{O}(1)$ because the roots of the heaps are accessed directly via `.peek()`.
* **Space Complexity:** $\mathcal{O}(n)$ to store all incoming stream elements.

---

## 13.5 - Classic LeetCode Problems

* **Find Median from Data Stream** (LeetCode #295)
* **Sliding Window Median** (LeetCode #480)
* **IPO** (Maximize Capital with Projects) (LeetCode #502)

---

## 13.6 - Sources used for this file:
https://leetcode.com/problems/find-median-from-data-stream/ <br>
https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d4f308d2bb2d978e29a4 <br>
https://www.geeksforgeeks.org/median-of-stream-of-integers-running-integers/ <br>
https://techinterviewhandbook.org/algorithms/heap/
