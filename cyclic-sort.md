# 7 - Cyclic Sort

## 7.1 - Overview

* The **Cyclic Sort** pattern iterates over an array of numbers where the values lie in a known, continuous range (typically $1$ to $n$ or $0$ to $n$).
* Since values are from $1$ to $n$, the element $x$ should ideally be placed at index $x - 1$.
* By swapping each number to its correct target index in-place, the entire array is sorted in $\mathcal{O}(n)$ time without extra space.

---

## 7.2 - Properties of a problem that suggests Cyclic Sort

* Problems involving an array of integers within a fixed range $[1, n]$ or $[0, n]$.
* Finding the **missing number**, **duplicate number**, **all missing numbers**, or the **first missing positive integer**.
* Must achieve $\mathcal{O}(n)$ time and $\mathcal{O}(1)$ auxiliary space.

---

## 7.3 - Classic Example: Find Missing and Duplicate Numbers

### Java Implementation

```java
public class CyclicSort {

    public static void cyclicSort(int[] nums) {
        int i = 0;
        while (i < nums.length) {
            int correctIndex = nums[i] - 1; // For range 1..n

            // If current element is within range and not at its correct position, swap it
            if (nums[i] > 0 && nums[i] <= nums.length && nums[i] != nums[correctIndex]) {
                swap(nums, i, correctIndex);
            } else {
                i++;
            }
        }
    }

    private static void swap(int[] arr, int i, int j) {
        int temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
    }
}
```

---

## 7.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(n)$ because each swap places at least one number in its correct final index position. The `while` loop executes at most $2n$ times.
* **Space Complexity:** $\mathcal{O}(1)$ auxiliary space since all swaps are executed in-place.

---

## 7.5 - Classic LeetCode Problems

* **Missing Number** (LeetCode #268)
* **Find All Numbers Disappeared in an Array** (LeetCode #448)
* **Find the Duplicate Number** (LeetCode #287)
* **Find All Duplicates in an Array** (LeetCode #442)
* **Set Mismatch** (LeetCode #645)
* **First Missing Positive** (LeetCode #41)

---

## 7.6 - Sources used for this file:
https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d43008d2bb2d978e2195 <br>
https://leetcode.com/problems/first-missing-positive/ <br>
https://emre.me/coding-patterns/cyclic-sort/ <br>
https://www.geeksforgeeks.org/cycle-sort/
