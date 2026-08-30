# 15 - Subsets, Permutations & Backtracking

## 15.1 - Overview & Theoretical Foundations (CLRS Appendix C & Chapter 34)

* **Backtracking** is a systematic method for iterating through all possible configurations of a search space (state-space tree).
* It builds candidates incrementally and abandons a candidate ("backtracks") as soon as it determines that the candidate cannot possibly be extended to a valid solution.
* **Combinatorial Structures:**
  1. **Subsets (Power Set):** Size is $2^n$. At each index $i$, binary decision: include $A[i]$ or exclude $A[i]$.
  2. **Permutations:** Size is $n!$. Order matters; maintain a visited state or swap-based selection across remaining elements.
  3. **Combinations:** Size is $\binom{n}{k} = \frac{n!}{k!(n-k)!}$. Fixed-length subset selections.

---

## 15.2 - Properties of a problem that suggests Backtracking

* Generating all valid arrangements, partitions, or assignments.
* Constraint satisfaction problems where pruning early cuts down exponential exploration (e.g. Sudoku, N-Queens).

---

## 15.3 - Classic Example: Generate All Subsets (Power Set)

### Java Implementation

```java
import java.util.ArrayList;
import java.util.List;

public class SubsetsBacktracking {

    public static List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        backtrack(result, new ArrayList<>(), nums, 0);
        return result;
    }

    private static void backtrack(List<List<Integer>> result, List<Integer> currentList, int[] nums, int start) {
        // Add snapshot of the current state to result
        result.add(new ArrayList<>(currentList));

        for (int i = start; i < nums.length; i++) {
            // 1. Choose candidate
            currentList.add(nums[i]);

            // 2. Explore deeper state
            backtrack(result, currentList, nums, i + 1);

            // 3. Unchoose / Backtrack to restore previous state
            currentList.remove(currentList.size() - 1);
        }
    }
}
```

---

### Go Implementation

```go
package main

// Subsets generates all possible subsets (the power set) of nums
func Subsets(nums []int) [][]int {
	var result [][]int
	var current []int

	var backtrack func(start int)
	backtrack = func(start int) {
		// Snapshot current subset
		subsetCopy := make([]int, len(current))
		copy(subsetCopy, current)
		result = append(result, subsetCopy)

		for i := start; i < len(nums); i++ {
			// Choose
			current = append(current, nums[i])

			// Explore
			backtrack(i + 1)

			// Backtrack
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return result
}
```

---

## 15.4 - Time & Space Complexity Analysis

* **Subsets:**
  * **Time Complexity:** $\mathcal{O}(n \cdot 2^n)$ because there are $2^n$ total subsets and copying each subset of average size $n/2$ takes $\mathcal{O}(n)$.
  * **Space Complexity:** $\mathcal{O}(n)$ maximum depth of the recursion tree.
* **Permutations:**
  * **Time Complexity:** $\mathcal{O}(n \cdot n!)$.
  * **Space Complexity:** $\mathcal{O}(n)$ call stack depth.

---

## 15.5 - Classic LeetCode & CLRS Benchmarks

* **Subsets I & II** (LeetCode #78, #90)
* **Permutations I & II** (LeetCode #46, #47)
* **Combination Sum I, II & III** (LeetCode #39, #40, #216)
* **N-Queens I & II** (LeetCode #51, #52)
* **Sudoku Solver** (LeetCode #37)
* **Word Search** (LeetCode #79)

---

## 15.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Appendix C: Counting and Probability (Permutations and combinations pp. 1198–1207)
* https://leetcode.com/problems/subsets/
* https://www.geeksforgeeks.org/backtracking-algorithms/
