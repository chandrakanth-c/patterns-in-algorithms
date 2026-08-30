# 3 - Two Pointers

## 3.1 - Overview

* The **Two Pointers** pattern uses two reference pointers that iterate through an iterable data structure (arrays, strings, or linked lists) simultaneously.
* Pointers can move:
  1. **Opposite Directions (Converging):** One pointer starts at the beginning (`left = 0`) and the other at the end (`right = n - 1`), moving inward.
  2. **Same Direction (Fast & Slow / Window):** Both pointers start at the beginning but advance at different speeds or conditions.
* This pattern drastically reduces brute-force $\mathcal{O}(n^2)$ pair searches to optimal $\mathcal{O}(n)$ linear scans by leveraging order/sorting.

---

## 3.2 - Properties of a problem that suggests Two Pointers

* The input array is **sorted** (or can be sorted in $\mathcal{O}(n \log n)$ without violating problem constraints).
* Looking for pairs, triplets, or subsegments meeting an exact condition (e.g., $A[i] + A[j] = \text{target}$).
* In-place element manipulation (e.g., deduplication, reversing, partitioning).

---

## 3.3 - Classic Example: Pair with Target Sum (Sorted Array)

### Java Implementation

```java
public class TwoPointers {

    public static int[] searchTwoSum(int[] arr, int targetSum) {
        int left = 0, right = arr.length - 1;

        while (left < right) {
            int currentSum = arr[left] + arr[right];

            if (currentSum == targetSum) {
                return new int[] { left, right }; // Found target pair
            }

            if (currentSum < targetSum) {
                left++; // Need a larger sum, move left pointer rightwards
            } else {
                right--; // Need a smaller sum, move right pointer leftwards
            }
        }

        return new int[] { -1, -1 }; // Pair not found
    }
}
```

---

## 3.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(n)$ because in each iteration at least one pointer moves closer to the other, inspecting each element at most once. (If unsorted initially, sorting dominates at $\mathcal{O}(n \log n)$).
* **Space Complexity:** $\mathcal{O}(1)$ auxiliary space as pointers require only a constant amount of memory.

---

## 3.5 - Classic LeetCode Problems

* **Two Sum II - Input Array Is Sorted** (LeetCode #167)
* **3Sum** (LeetCode #15)
* **Container With Most Water** (LeetCode #11)
* **Valid Palindrome** (LeetCode #125)
* **Remove Duplicates from Sorted Array** (LeetCode #26)
* **Trapping Rain Water (Two Pointer Approach)** (LeetCode #42)

---

## 3.6 - Sources used for this file:
https://leetcode.com/explore/interview/card/cheatsheets/720/resources/4723/ <br>
https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d38a08d2bb2d978e1cd7 <br>
https://techinterviewhandbook.org/algorithms/array/ <br>
https://www.geeksforgeeks.org/two-pointers-technique/
