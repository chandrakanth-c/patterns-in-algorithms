# 7 - Cyclic Sort

## 7.1 - Overview & Theoretical Foundations (CLRS Chapter 8)

* Standard comparison sorts (MergeSort, HeapSort) are lower-bounded by $\Omega(n \log n)$ (CLRS Theorem 8.1: Lower bounds for sorting).
* **Cyclic Sort** is an in-place, linear-time $\mathcal{O}(n)$ permutation sorting algorithm applicable when elements are bounded within a known discrete range (e.g. $[1, n]$ or $[0, n]$).
* Because the domain maps directly to array indices, the correct index for element $x$ in range $[1, n]$ is known *a priori* to be $x - 1$.
* By cyclically swapping each misplaced element $A[i]$ to its target index $A[i] - 1$, each swap places at least one element into its correct final position.

---

## 7.2 - Properties of a problem that suggests Cyclic Sort

* Input is an array containing numbers strictly in range $[1, n]$ or $[0, n]$.
* Finding the **missing number**, **duplicate number**, or **corrupted pair** in $\mathcal{O}(n)$ time and $\mathcal{O}(1)$ space.

---

## 7.3 - Classic Example: Find All Numbers Disappeared in an Array

### Java Implementation

```java
import java.util.ArrayList;
import java.util.List;

public class CyclicSort {

    public static List<Integer> findDisappearedNumbers(int[] nums) {
        int i = 0;
        while (i < nums.length) {
            int correctIndex = nums[i] - 1; // 1-based to 0-based mapping

            if (nums[i] > 0 && nums[i] <= nums.length && nums[i] != nums[correctIndex]) {
                swap(nums, i, correctIndex);
            } else {
                i++;
            }
        }

        List<Integer> missingNumbers = new ArrayList<>();
        for (i = 0; i < nums.length; i++) {
            if (nums[i] != i + 1) {
                missingNumbers.add(i + 1);
            }
        }

        return missingNumbers;
    }

    private static void swap(int[] arr, int i, int j) {
        int temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
    }
}
```

---

### Go Implementation

```go
package main

// FindDisappearedNumbers finds all numbers in range [1, n] that do not appear in the array
func FindDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		correctIndex := nums[i] - 1

		if nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[correctIndex] {
			nums[i], nums[correctIndex] = nums[correctIndex], nums[i]
		} else {
			i++
		}
	}

	var missing []int
	for idx, val := range nums {
		if val != idx+1 {
			missing = append(missing, idx+1)
		}
	}

	return missing
}
```

---

## 7.4 - Time & Space Complexity Analysis

* **Time Complexity:** $\mathcal{O}(n)$ — Each swap places at least one element into its final correct slot. Since an array of length $n$ can have at most $n$ elements placed into position, the `while` loop executes at most $2n$ total iterations.
* **Space Complexity:** $\mathcal{O}(1)$ auxiliary space (mutates input array in-place).

---

## 7.5 - Classic LeetCode & CLRS Benchmarks

### 7.5.1 - Missing Number (LeetCode #268)

#### 1. Problem Statement
Given an array `nums` containing `n` distinct numbers in the range `[0, n]`, return the only number in the range that is missing from the array.

#### 2. Solution Link
* [Go Implementation](problems/cyclic-sort/cyclic_sort_problems.go) (Function: `MissingNumber`)

#### 3. Explanation
Since the numbers are in the range `[0, n]`, each number `v` should ideally be at index `v`. We use cyclic sort to place each number `nums[i]` at index `nums[i]`, provided `nums[i] < n`. After sorting, we scan the array; the first index `i` where `nums[i] != i` is the missing number. If all indices match, then `n` is the missing number.

#### 4. Conceptual Link to Cyclic Sort
Demonstrates **Direct Indexing**. The value of the element itself tells us exactly where it belongs in the array, allowing for $\mathcal{O}(n)$ detection without extra space.

### 7.5.2 - Find All Numbers Disappeared in an Array (LeetCode #448)

#### 1. Problem Statement
Given an array `nums` of `n` integers where `nums[i]` is in the range `[1, n]`, return an array of all the integers in the range `[1, n]` that do not appear in `nums`.

#### 2. Solution Link
* [Go Implementation](problems/cyclic-sort/cyclic_sort_problems.go) (Function: `FindDisappearedNumbers`)
* [Java Implementation](problems/cyclic-sort/CyclicSort.java) (Method: `findDisappearedNumbers`)

#### 3. Explanation
For an array of size `n` with values in `[1, n]`, the correct index for value `x` is `x - 1`. We iterate through the array and swap `nums[i]` with `nums[nums[i]-1]` until each number is at its correct index or a duplicate is found. A second pass identifies indices `i` where `nums[i] != i + 1`.

#### 4. Conceptual Link to Cyclic Sort
Illustrates handling of **duplicate entries** in the input range. Cyclic sort effectively "ignores" duplicates by leaving them in slots that belong to missing numbers.

### 7.5.3 - Find the Duplicate Number (LeetCode #287)

#### 1. Problem Statement
Given an array of integers `nums` containing `n + 1` integers where each integer is in the range `[1, n]` inclusive. There is only one repeated number in `nums`, return this repeated number.

#### 2. Solution Link
* [Go Implementation](problems/cyclic-sort/cyclic_sort_problems.go) (Function: `FindDuplicate`)

#### 3. Explanation
We attempt to place each number `x` at index `x - 1`. If we encounter a number `nums[i]` that is already present at its correct target index `nums[nums[i]-1]`, and `i` is not that index, then `nums[i]` is the duplicate.

#### 4. Conceptual Link to Cyclic Sort
This problem highlights the **Pigeonhole Principle**. With `n+1` numbers in a range of `n`, at least one must be a duplicate. Cyclic sort finds it by identifying the "pigeon" that cannot find an empty "hole" because its correct hole is already occupied by an identical value.

### 7.5.4 - First Missing Positive (LeetCode #41)

#### 1. Problem Statement
Given an unsorted integer array `nums`, return the smallest missing positive integer. You must implement an algorithm that runs in $\mathcal{O}(n)$ time and uses constant extra space.

#### 2. Solution Link
* [Go Implementation](problems/cyclic-sort/cyclic_sort_problems.go) (Function: `FirstMissingPositive`)

#### 3. Explanation
We use cyclic sort to place every positive integer `x` at index `x - 1` if $1 \le x \le n$. Non-positive numbers and numbers larger than `n` are ignored. After the sorting pass, the first index `i` such that `nums[i] != i + 1` indicates that `i + 1` is the smallest missing positive integer.

#### 4. Conceptual Link to Cyclic Sort
Shows how to apply the pattern to **unbounded ranges** by filtering for the relevant sub-range `[1, n]`. It is a hard-level problem that reduces to a simple linear scan after applying the cyclic sort transformation.

---

## 7.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 8: Sorting in Linear Time (pp. 205–221)
  * Section 8.1: Lower bounds for sorting (pp. 205–208)
* https://leetcode.com/problems/first-missing-positive/
* https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d43008d2bb2d978e2195
