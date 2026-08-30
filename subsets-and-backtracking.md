# 15 - Subsets, Permutations & Backtracking

## 15.1 - Overview

* **Backtracking** is a general algorithmic technique for finding all (or some) solutions to computational problems incrementally, one piece at a time, removing solutions that fail to satisfy constraints at any point in time ("backtracking").
* Useful for generating the combinatorial search space:
  * **Subsets (Power Set):** $2^n$ combinations (choose or don't choose).
  * **Permutations:** $n!$ orderings of elements.
  * **Combinations:** Subsets of a fixed length $k$.

---

## 15.2 - Properties of a problem that suggests Backtracking

* Asked to find **all combinations**, **all permutations**, or **all valid paths/configurations**.
* The problem asks to solve a constraint satisfaction puzzle (e.g. Sudoku, N-Queens, Word Search).
* The brute-force requires exploring a state-space tree.

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
        // Add a copy of the current subset to result
        result.add(new ArrayList<>(currentList));

        for (int i = start; i < nums.length; i++) {
            // 1. Choose
            currentList.add(nums[i]);

            // 2. Explore
            backtrack(result, currentList, nums, i + 1);

            // 3. Un-choose (Backtrack)
            currentList.remove(currentList.size() - 1);
        }
    }
}
```

---

## 15.4 - Time & Space Complexity

* **Subsets:**
  * **Time Complexity:** $\mathcal{O}(n \cdot 2^n)$ because there are $2^n$ subsets and copying each subset takes $\mathcal{O}(n)$ time.
  * **Space Complexity:** $\mathcal{O}(n)$ recursion call stack space.
* **Permutations:**
  * **Time Complexity:** $\mathcal{O}(n \cdot n!)$
  * **Space Complexity:** $\mathcal{O}(n)$ recursion depth.

---

## 15.5 - Classic LeetCode Problems

* **Subsets I & II** (LeetCode #78, #90)
* **Permutations I & II** (LeetCode #46, #47)
* **Combination Sum I, II & III** (LeetCode #39, #40, #216)
* **Generate Parentheses** (LeetCode #22)
* **N-Queens I & II** (LeetCode #51, #52)
* **Sudoku Solver** (LeetCode #37)
* **Word Search** (LeetCode #79)

---

## 15.6 - Sources used for this file:
https://leetcode.com/problems/subsets/ <br>
https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d47508d2bb2d978e23ca <br>
https://www.geeksforgeeks.org/backtracking-algorithms/ <br>
https://techinterviewhandbook.org/algorithms/recursion/
