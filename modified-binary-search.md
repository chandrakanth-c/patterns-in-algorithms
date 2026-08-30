# 11 - Modified Binary Search & Binary Search on Answer

## 11.1 - Overview & Theoretical Foundations (CLRS Chapter 2 & 9)

* **Binary Search** divides the remaining candidate search space in half at each iteration by testing the midpoint $q = \lfloor (p + r)/2 \rfloor$.
* **Loop Invariant:** At the start of each iteration, the target value (or optimal boundary) is guaranteed to lie within the closed interval $[left, right]$ if it exists.
* **Modified Binary Search Paradigms:**
  1. **Rotated Sorted Arrays:** Since a rotated array consists of two monotonically increasing sub-segments, at least one half ($[left, mid]$ or $[mid, right]$) is always strictly sorted. Checking boundary conditions allows determining which branch to prune in $\mathcal{O}(1)$.
  2. **Binary Search on Answer Space (Monotonic Predicates):** Given a monotonically non-decreasing or non-increasing boolean function $f(x) \in \{\text{false}, \dots, \text{false}, \text{true}, \dots, \text{true}\}$, binary search finds the exact boundary value $x^*$ where $f(x^*) = \text{true}$ in $\mathcal{O}(\text{cost}(f) \cdot \log(\text{Range}))$.

---

## 11.2 - Properties of a problem that suggests Binary Search

* Input is **sorted**, **rotated sorted**, or search domain is bounded in integer range $[\text{low}, \text{high}]$.
* Optimization queries asking to find the **minimum capacity**, **maximum speed**, or **smallest maximum** that makes a solution feasible.

---

## 11.3 - Classic Example: Search in Rotated Sorted Array

### Java Implementation

```java
public class ModifiedBinarySearch {

    public static int searchRotated(int[] nums, int target) {
        int left = 0, right = nums.length - 1;

        while (left <= right) {
            int mid = left + (right - left) / 2;

            if (nums[mid] == target) return mid;

            // Check if the left half is strictly sorted
            if (nums[left] <= nums[mid]) {
                if (target >= nums[left] && target < nums[mid]) {
                    right = mid - 1; // Target is in the sorted left half
                } else {
                    left = mid + 1;  // Target is in the right half
                }
            }
            // Otherwise, the right half must be strictly sorted
            else {
                if (target > nums[mid] && target <= nums[right]) {
                    left = mid + 1;  // Target is in the sorted right half
                } else {
                    right = mid - 1; // Target is in the left half
                }
            }
        }

        return -1; // Not found
    }
}
```

---

### Go Implementation

```go
package main

// SearchRotated searches for target in a rotated sorted array
func SearchRotated(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// Check if left half is sorted
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1 // Search left half
			} else {
				left = mid + 1 // Search right half
			}
		} else { // Right half is sorted
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1 // Search right half
			} else {
				right = mid - 1 // Search left half
			}
		}
	}

	return -1
}
```

---

## 11.4 - Time & Space Complexity Analysis

* **Time Complexity:**
  * Array Search: $\mathcal{O}(\log n)$ because the search space decreases by a factor of 2 per step ($T(n) = T(n/2) + \mathcal{O}(1) \implies \Theta(\log n)$ by Master Theorem).
  * Binary Search on Answer: $\mathcal{O}(n \log(\text{maxVal} - \text{minVal}))$ where feasibility function takes $\mathcal{O}(n)$.
* **Space Complexity:** $\mathcal{O}(1)$ iterative auxiliary space.

---

## 11.5 - Classic LeetCode & CLRS Benchmarks

* **Binary Search** (CLRS 2.3 / LeetCode #704)
* **Search in Rotated Sorted Array I & II** (LeetCode #33, #81)
* **Find Minimum in Rotated Sorted Array** (LeetCode #153)
* **Koko Eating Bananas** (LeetCode #875)
* **Capacity To Ship Packages Within D Days** (LeetCode #1011)
* **Split Array Largest Sum** (LeetCode #410)
* **Median of Two Sorted Arrays** (LeetCode #4)

---

## 11.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Section 2.3: Designing algorithms (Binary search exercise pp. 44)
  * Chapter 9: Medians and Order Statistics (pp. 227–241)
* https://leetcode.com/problems/search-in-rotated-sorted-array/
* https://techinterviewhandbook.org/algorithms/binary-search/
