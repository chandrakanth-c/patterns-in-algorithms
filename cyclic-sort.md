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

* **Counting Sort / Direct Indexing** (CLRS 8.2)
* **Missing Number** (LeetCode #268)
* **Find All Numbers Disappeared in an Array** (LeetCode #448)
* **Find the Duplicate Number** (LeetCode #287)
* **First Missing Positive** (LeetCode #41)
* **Set Mismatch** (LeetCode #645)

---

## 7.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 8: Sorting in Linear Time (pp. 205–221)
  * Section 8.1: Lower bounds for sorting (pp. 205–208)
* https://leetcode.com/problems/first-missing-positive/
* https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d43008d2bb2d978e2195
