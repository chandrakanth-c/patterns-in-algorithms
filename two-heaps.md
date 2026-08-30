# 13 - Two Heaps (Median of a Stream)

## 13.1 - Overview & Theoretical Foundations (CLRS Chapter 6)

* The **Two Heaps** pattern partitions a dynamic, incoming stream of real-time numbers into two balanced halves:
  1. **Max-Heap ($H_{\text{low}}$):** Stores the lower half of numbers (the root is $\max(H_{\text{low}})$).
  2. **Min-Heap ($H_{\text{high}}$):** Stores the upper half of numbers (the root is $\min(H_{\text{high}})$).
* **Heap Invariants:**
  * **Order Invariant:** Every element in $H_{\text{low}}$ is less than or equal to every element in $H_{\text{high}}$: $\max(H_{\text{low}}) \le \min(H_{\text{high}})$.
  * **Size Invariant:** Either $|H_{\text{low}}| = |H_{\text{high}}|$ (when total elements $N$ is even), or $|H_{\text{low}}| = |H_{\text{high}}| + 1$ (when $N$ is odd).
* When invariants hold, the median is calculated in $\mathcal{O}(1)$ time:
  $$\text{Median} = \begin{cases} \max(H_{\text{low}}) & \text{if } |H_{\text{low}}| > |H_{\text{high}}| \\ \frac{\max(H_{\text{low}}) + \min(H_{\text{high}})}{2.0} & \text{if } |H_{\text{low}}| = |H_{\text{high}}| \end{cases}$$

---

## 13.2 - Properties of a problem that suggests Two Heaps

* Finding the **median** or split-percentiles dynamically in a stream where numbers are constantly inserted.
* Real-time tracking of two extreme boundaries of a partition.

---

## 13.3 - Classic Example: Find Median from Data Stream

### Java Implementation

```java
import java.util.Collections;
import java.util.PriorityQueue;

public class MedianFinder {
    private PriorityQueue<Integer> maxHeap; // Lower half
    private PriorityQueue<Integer> minHeap; // Upper half

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

        // Rebalance size invariant
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

### Go Implementation

```go
package main

import (
	"container/heap"
)

// IntMaxHeap implements max-heap of ints
type IntMaxHeap []int
func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Reverse for max-heap
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// IntMinHeap implements min-heap of ints
type IntMinHeap []int
func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	maxHeap *IntMaxHeap
	minHeap *IntMinHeap
}

func Constructor() MedianFinder {
	maxH := &IntMaxHeap{}
	minH := &IntMinHeap{}
	heap.Init(maxH)
	heap.Init(minH)
	return MedianFinder{maxHeap: maxH, minHeap: minH}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.maxHeap.Len() == 0 || num <= (*mf.maxHeap)[0] {
		heap.Push(mf.maxHeap, num)
	} else {
		heap.Push(mf.minHeap, num)
	}

	// Rebalance
	if mf.maxHeap.Len() > mf.minHeap.Len()+1 {
		val := heap.Pop(mf.maxHeap).(int)
		heap.Push(mf.minHeap, val)
	} else if mf.minHeap.Len() > mf.maxHeap.Len() {
		val := heap.Pop(mf.minHeap).(int)
		heap.Push(mf.maxHeap, val)
	}
}

func (mf *MedianFinder) FindMedian() double {
	if mf.maxHeap.Len() == mf.minHeap.Len() {
		return float64((*mf.maxHeap)[0]+(*mf.minHeap)[0]) / 2.0
	}
	return float64((*mf.maxHeap)[0])
}
```

---

## 13.4 - Time & Space Complexity Analysis

* **`addNum(num)`:** $\mathcal{O}(\log n)$ time due to heap insertion and balance operations.
* **`findMedian()`:** $\mathcal{O}(1)$ direct lookup of heap roots.
* **Space Complexity:** $\mathcal{O}(n)$ to hold the stream values.

---

## 13.5 - Classic LeetCode & CLRS Benchmarks

* **Find Median from Data Stream** (LeetCode #295)
* **Sliding Window Median** (LeetCode #480)
* **IPO** (Maximize Capital with Projects) (LeetCode #502)

---

## 13.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 6: Heapsort & Priority Queues (pp. 161–181)
* https://leetcode.com/problems/find-median-from-data-stream/
* https://www.geeksforgeeks.org/median-of-stream-of-integers-running-integers/
