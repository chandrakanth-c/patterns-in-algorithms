# 11 - Modified Binary Search & Binary Search on Answer

## 11.1 - Overview

* Classical **Binary Search** finds an element in a strictly sorted array in $\mathcal{O}(\log n)$ time by dividing search bounds in half.
* **Modified Binary Search** expands this concept to:
  1. **Rotated / Shifted Sorted Arrays:** Identifying which half is strictly sorted to decide search branch.
  2. **Unknown Array Bounds:** Finding range via exponential step doubling.
  3. **Binary Search on Answer Space:** When optimizing a value $X$ where a feasibility check $f(X)$ is monotonically boolean (e.g. `[false, false, ..., true, true]`).

---

## 11.2 - Properties of a problem that suggests Binary Search

* Input is **sorted**, **nearly sorted**, or **rotated sorted**.
* Optimization questions asking to find the **minimum capacity**, **maximum speed**, or **smallest maximum** that satisfies a condition where the predicate is monotonic.

---

## 11.3 - Classic Example: Search in Rotated Sorted Array

### Java Implementation

```java
public class ModifiedBinarySearch {

    public static int searchInRotatedArray(int[] nums, int target) {
        int left = 0, right = nums.length - 1;

        while (left <= right) {
            int mid = left + (right - left) / 2;

            if (nums[mid] == target) return mid;

            // Check if left half is sorted
            if (nums[left] <= nums[mid]) {
                if (target >= nums[left] && target < nums[mid]) {
                    right = mid - 1; // Target lies in sorted left half
                } else {
                    left = mid + 1;  // Target lies in right half
                }
            }
            // Otherwise, right half is sorted
            else {
                if (target > nums[mid] && target <= nums[right]) {
                    left = mid + 1;  // Target lies in sorted right half
                } else {
                    right = mid - 1; // Target lies in left half
                }
            }
        }

        return -1; // Not found
    }
}
```

---

## 11.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(\log n)$ because the search space is halved at each iteration. (For Binary Search on Answer Space: $\mathcal{O}(\text{cost of } f(X) \times \log(\text{Range}))$.
* **Space Complexity:** $\mathcal{O}(1)$ iterative auxiliary space.

---

## 11.5 - Classic LeetCode Problems

* **Binary Search** (LeetCode #704)
* **Search in Rotated Sorted Array I & II** (LeetCode #33, #81)
* **Find Minimum in Rotated Sorted Array** (LeetCode #153)
* **Koko Eating Bananas** (Binary Search on Answer) (LeetCode #875)
* **Capacity To Ship Packages Within D Days** (LeetCode #1011)
* **Split Array Largest Sum** (LeetCode #410)
* **Median of Two Sorted Arrays** (LeetCode #4)

---

## 11.6 - Sources used for this file:
https://leetcode.com/problems/search-in-rotated-sorted-array/ <br>
https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d49108d2bb2d978e24fd <br>
https://techinterviewhandbook.org/algorithms/binary-search/ <br>
https://en.wikipedia.org/wiki/Binary_search_algorithm
