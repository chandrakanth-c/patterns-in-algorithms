# 3 - Two Pointers

## 3.1 - Overview & Theoretical Foundations (CLRS Chapter 2 & 10)

* The **Two Pointers** pattern maintains two indices that traverse a linear data structure (such as an array, string, or linked list) in tandem.
* Pointers can move:
  1. **Converging (Opposite Directions):** Left pointer at index $0$ increasing, right pointer at index $n-1$ decreasing.
  2. **Co-directional (Fast & Slow / Equidirectional):** Both pointers start at the front and advance at varying rates based on predicates.
* **Loop Invariant (CLRS 2.1):**
  * *Initialization:* Prior to loop execution, the search space for the target solution contains all possible pairs $\{A[i], A[j] \mid 0 \le i < j \le n-1\}$.
  * *Maintenance:* In each iteration, if $A[i] + A[j] < \text{target}$, then for all $k \le j$, $A[i] + A[k] \le A[i] + A[j] < \text{target}$, so index $i$ cannot be part of any valid solution with any index $\le j$. Incrementing $i$ safely prunes the search space without missing any valid pair. Symmetrically, if $A[i] + A[j] > \text{target}$, decrementing $j$ safely prunes invalid pairings.
  * *Termination:* The loop terminates either when a matching pair is found or $i \ge j$, guaranteeing all pairs have been considered.

---

## 3.2 - Properties of a problem that suggests Two Pointers

* Input sequence is **sorted** (or sorting in $\mathcal{O}(n \log n)$ is acceptable).
* Finding pairs, triplets, or subsegments satisfying a target arithmetic condition.
* In-place element rearrangement, partitioning, or two-way partitioning (e.g. Hoare partition in Quicksort).

---

## 3.3 - Classic Example: Pair with Target Sum (Two Sum II)

### Java Implementation

```java
public class TwoPointers {

    public static int[] twoSumSorted(int[] numbers, int target) {
        int left = 0;
        int right = numbers.length - 1;

        while (left < right) {
            int currentSum = numbers[left] + numbers[right];

            if (currentSum == target) {
                return new int[] { left + 1, right + 1 }; // 1-indexed result
            } else if (currentSum < target) {
                left++; // Increase sum by moving left pointer rightward
            } else {
                right--; // Decrease sum by moving right pointer leftward
            }
        }

        return new int[] { -1, -1 };
    }
}
```

---

### Go Implementation

```go
package main

// TwoSumSorted finds two indices (1-based) in a sorted array that sum to target
func TwoSumSorted(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		currentSum := numbers[left] + numbers[right]

		if currentSum == target {
			return []int{left + 1, right + 1}
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}

	return []int{-1, -1}
}
```

---

## 3.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(n)$ — In each step, the distance `right - left` strictly decreases by at least 1, resulting in at most $n$ loop iterations.
* **Space Complexity:** $\mathcal{O}(1)$ auxiliary space.

---

## 3.5 - Classic LeetCode & CLRS Problems

* **Two Sum II - Input Array Is Sorted** (LeetCode #167)
* **3Sum / 4Sum** (LeetCode #15, #18)
* **Container With Most Water** (LeetCode #11)
* **Valid Palindrome** (LeetCode #125)
* **Remove Duplicates from Sorted Array** (LeetCode #26)
* **Hoare Partitioning in Quicksort** (CLRS 7.1)
* **Trapping Rain Water** (LeetCode #42)

---

## 3.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Section 2.1: Insertion sort & Loop Invariants (pp. 17–24)
  * Section 7.1: Description of quicksort & partition (pp. 182–190)
* https://leetcode.com/explore/interview/card/cheatsheets/720/resources/4723/
* https://techinterviewhandbook.org/algorithms/array/
