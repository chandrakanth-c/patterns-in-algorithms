# 9 - Monotonic Stack / Queue

## 9.1 - Overview & Theoretical Foundations (CLRS Chapter 10 & 17)

* A **Monotonic Stack** enforces a strict invariant: elements stored inside the stack remain monotonically increasing or monotonically decreasing from bottom to top.
* When evaluating a new element $x$, elements that violate the monotonic invariant are popped before pushing $x$.
* **Amortized Analysis (CLRS 17.1 - Aggregate & Accounting Method):**
  * Let each push operation be charged 2 credits (1 credit pays for the push, 1 credit is banked on the element to pay for its eventual pop).
  * Every element enters the stack at most once and leaves at most once.
  * For an input sequence of length $n$, the total number of operations across the entire algorithm is upper-bounded by $2n$.
  * Hence, the amortized cost per element is $\mathcal{O}(1)$, yielding an overall running time of $\Theta(n)$.

---

## 9.2 - Properties of a problem that suggests Monotonic Stack

* Finding the **Next Greater Element**, **Next Smaller Element**, **Previous Greater Element**, or **Previous Smaller Element** for every element in an array.
* Finding contiguous boundary limits (e.g. Largest Rectangle in Histogram, Trapping Rain Water).
* Monotonic Queue: Finding maximum/minimum in a sliding window of size $K$.

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
            while (!stack.isEmpty() && nums[i] > nums[stack.peek()]) {
                int poppedIndex = stack.pop();
                result[poppedIndex] = nums[i]; // nums[i] is the next greater element
            }
            stack.push(i);
        }

        return result;
    }
}
```

---

### Go Implementation

```go
package main

// NextGreaterElement finds the next greater element for each item in nums
func NextGreaterElement(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}

	// Monotonic stack storing indices
	stack := make([]int, 0, n)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			poppedIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Pop
			result[poppedIdx] = nums[i]
		}
		stack = append(stack, i) // Push
	}

	return result
}
```

---

## 9.4 - Time & Space Complexity Analysis

* **Time Complexity:** $\mathcal{O}(n)$ total amortized time.
* **Space Complexity:** $\mathcal{O}(n)$ to store indices on the stack.

---

## 9.5 - Classic LeetCode & Benchmark Problems

* **Daily Temperatures** (LeetCode #739)
* **Next Greater Element I & II** (LeetCode #496, #503)
* **Largest Rectangle in Histogram** (LeetCode #84)
* **Maximal Rectangle** (LeetCode #85)
* **Trapping Rain Water** (LeetCode #42)
* **Online Stock Span** (LeetCode #901)
* **Sliding Window Maximum** (Monotonic Deque) (LeetCode #239)

---

## 9.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 10: Elementary Data Structures (Stacks pp. 253–255)
  * Chapter 17: Amortized Analysis (Aggregate and Accounting method pp. 488–496)
* https://leetcode.com/discuss/study-guide/2347639/A-comprehensive-guide-and-template-for-monotonic-stack-based-problems
* https://techinterviewhandbook.org/algorithms/stack/
