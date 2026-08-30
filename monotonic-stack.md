# 9 - Monotonic Stack / Queue

## 9.1 - Overview

* A **Monotonic Stack** (or Queue) is a stack whose elements are always sorted in a strictly monotonic order (either monotonically increasing or monotonically decreasing).
* When iterating through an array, whenever a new element violates the monotonic order, we continuously pop elements from the stack until the order is restored.
* This property allows finding the **Next Greater Element**, **Next Smaller Element**, **Previous Greater Element**, or **Previous Smaller Element** for every element in an array in $\mathcal{O}(n)$ time.

---

## 9.2 - Properties of a problem that suggests Monotonic Stack

* Finding nearest elements that are larger/smaller than current elements.
* Calculating bounded subsegment areas (e.g., maximum rectangle in histograms).
* Processing stock spans, temperature increases, or trapping rainwater boundaries.

---

## 9.3 - Classic Example: Next Greater Element

### Java Implementation

```java
import java.util.ArrayDeque;
import java.util.Arrays;
import java.util.Deque;

public class MonotonicStack {

    public static int[] nextGreaterElement(int[] nums) {
        int n = nums.length;
        int[] result = new int[n];
        Arrays.fill(result, -1);

        // Monotonic decreasing stack (stores indices)
        Deque<Integer> stack = new ArrayDeque<>();

        for (int i = 0; i < n; i++) {
            // While current number is greater than the number represented by the top index
            while (!stack.isEmpty() && nums[i] > nums[stack.peek()]) {
                int poppedIndex = stack.pop();
                result[poppedIndex] = nums[i]; // Found next greater element
            }
            stack.push(i);
        }

        return result;
    }
}
```

---

## 9.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(n)$ amortized. Although there is a nested `while` loop, every element is pushed onto the stack exactly once and popped at most once across the entire iteration ($2n$ operations).
* **Space Complexity:** $\mathcal{O}(n)$ to store indices on the stack.

---

## 9.5 - Classic LeetCode Problems

* **Daily Temperatures** (LeetCode #739)
* **Next Greater Element I & II** (LeetCode #496, #503)
* **Largest Rectangle in Histogram** (LeetCode #84)
* **Maximal Rectangle** (LeetCode #85)
* **Trapping Rain Water** (LeetCode #42)
* **Online Stock Span** (LeetCode #901)
* **Sliding Window Maximum** (Monotonic Deque) (LeetCode #239)

---

## 9.6 - Sources used for this file:
https://leetcode.com/discuss/study-guide/2347639/A-comprehensive-guide-and-template-for-monotonic-stack-based-problems <br>
https://www.geeksforgeeks.org/introduction-to-monotonic-stack-2/ <br>
https://techinterviewhandbook.org/algorithms/stack/ <br>
https://labuladong.gitbook.io/algo-en/ii.-data-structure/monotonicstack
