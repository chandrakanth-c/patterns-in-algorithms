# 12 - Top 'K' Elements (Heaps & Quickselect)

## 12.1 - Overview

* The **Top 'K' Elements** pattern finds the top $K$ largest, smallest, or most frequent elements in an unsorted collection or continuous stream.
* Two primary approaches:
  1. **Min-Heap / Max-Heap:** Maintain a heap of fixed size $K$. To find the $K$ largest elements, maintain a **Min-Heap** of size $K$.
  2. **Quickselect:** Partitioning algorithm based on QuickSort achieving linear $\mathcal{O}(n)$ average time.

---

## 12.2 - Properties of a problem that suggests Top 'K' Elements

* Asked to find the **$K$-th largest/smallest** element.
* Asked to find the **$K$ most frequent** items or **$K$ closest points**.
* Unsorted input where sorting the entire collection $\mathcal{O}(n \log n)$ is suboptimal when $K \ll n$.

---

## 12.3 - Classic Example: Kth Largest Element in an Array (Min-Heap)

### Java Implementation

```java
import java.util.PriorityQueue;

public class TopKElements {

    public static int findKthLargest(int[] nums, int k) {
        // Min-heap to keep track of the k largest elements seen so far
        PriorityQueue<Integer> minHeap = new PriorityQueue<>(k);

        for (int num : nums) {
            minHeap.offer(num);

            // If heap size exceeds k, remove the smallest element
            if (minHeap.size() > k) {
                minHeap.poll();
            }
        }

        // The root of the min-heap contains the k-th largest element
        return minHeap.peek();
    }
}
```

---

## 12.4 - Time & Space Complexity

* **Heap Approach:**
  * **Time Complexity:** $\mathcal{O}(n \log K)$ because inserting into a heap of size $K$ takes $\mathcal{O}(\log K)$ time.
  * **Space Complexity:** $\mathcal{O}(K)$ to store $K$ elements in the heap.
* **Quickselect Approach:**
  * **Time Complexity:** $\mathcal{O}(n)$ average, $\mathcal{O}(n^2)$ worst case.
  * **Space Complexity:** $\mathcal{O}(1)$ iterative or $\mathcal{O}(\log n)$ recursive stack.

---

## 12.5 - Classic LeetCode Problems

* **Kth Largest Element in an Array** (LeetCode #215)
* **Top K Frequent Elements** (LeetCode #347)
* **K Closest Points to Origin** (LeetCode #973)
* **Sort Characters By Frequency** (LeetCode #451)
* **Reorganize String** (LeetCode #767)
* **Kth Largest Element in a Stream** (LeetCode #703)

---

## 12.6 - Sources used for this file:
https://leetcode.com/problems/kth-largest-element-in-an-array/ <br>
https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d4cb08d2bb2d978e27c1 <br>
https://www.geeksforgeeks.org/k-largestor-smallest-elements-in-an-array/ <br>
https://techinterviewhandbook.org/algorithms/heap/
