# 1 - Dynamic Programming (DP)

## 1.1 - Overview (CLRS Chapter 14)

* Dynamic Programming (DP), like the [Divide-and-Conquer](divide-and-conquer.md) method, solves problems by combining the solutions to subproblems.
* **Divide-and-Conquer** partitions a problem into *disjoint* subproblems, solves the subproblems recursively, and combines their solutions.
* In contrast, **Dynamic Programming** applies when the subproblems **overlap**—that is, when subproblems share subsubproblems. A divide-and-conquer algorithm would do more work than necessary, repeatedly solving the common subproblems. A DP algorithm solves every subsubproblem just once and saves its answer in a table (memory cache or table), thereby avoiding the work of recomputing the answer every time the subsubproblem is encountered.
* DP is typically applied to **optimization problems**. Such problems can have many possible solutions. Each solution has a value, and we wish to find a solution with an *optimal* (minimum or maximum) value.

---

## 1.2 - Four-Step Development Process (CLRS 14.1)

When developing a dynamic-programming algorithm, follow this formal four-step sequence:
1. **Characterize the structure of an optimal solution.**
2. **Recursively define the value of an optimal solution.**
3. **Compute the value of an optimal solution (typically in a bottom-up or top-down memoized fashion).**
4. **Construct an optimal solution from computed information (if the actual choices, not just the optimal value, are needed).**

---

## 1.3 - Two Essential Ingredients of Dynamic Programming (CLRS 14.3)

For dynamic programming to be applicable, an optimization problem must exhibit two key properties:

### 1. Optimal Substructure
A problem exhibits **optimal substructure** if an optimal solution to the problem contains within it optimal solutions to subproblems. 
* *Example (Shortest Path):* If vertex $w$ is on an optimal shortest path from $u$ to $v$, then the subpaths $u \to w$ and $w \to v$ must also be optimal shortest paths between their respective endpoints.

### 2. Overlapping Subproblems
A problem has **overlapping subproblems** if a recursive algorithm revisits the same subproblems repeatedly over a small state space, rather than generating new subproblems at every branch.
* In a problem of size $n$, the total number of distinct subproblems is typically a polynomial $\mathcal{O}(n)$ or $\mathcal{O}(n^2)$, whereas the recursive call tree without caching would be exponential $\mathcal{O}(2^n)$.

---

## 1.4 - Two Approaches to Implementation

### 1. Top-Down with Memoization
Write the recursive procedure naturally. Before calculating the solution to a subproblem, check if it is already stored in a cache/hash table. If so, return the cached result. If not, compute it recursively, store it in the table, and return it.

### 2. Bottom-Up with Tabulation
Order the subproblems by "size" such that smaller subproblems are solved before larger ones. Compute solutions iteratively from base cases up to the target state, storing answers in a 1D, 2D, or multi-dimensional array.

---

## 1.5 - Classic Example: Fibonacci Sequence & 0/1 Knapsack

### Mathematical Formulation
$$\text{Fib}(n) = \begin{cases} 0 & \text{if } n = 0 \\ 1 & \text{if } n = 1 \\ \text{Fib}(n-1) + \text{Fib}(n-2) & \text{if } n \ge 2 \end{cases}$$

---

### Java Implementation

```java
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class DynamicProgramming {

    // 1. Top-Down Memoization Approach (Fibonacci)
    public static long fibMemo(int n, Map<Integer, Long> memo) {
        if (n <= 1) return n;
        if (memo.containsKey(n)) return memo.get(n);

        long result = fibMemo(n - 1, memo) + fibMemo(n - 2, memo);
        memo.put(n, result);
        return result;
    }

    // 2. Bottom-Up Tabulation Approach with Space Optimization (Fibonacci)
    public static long fibTabulation(int n) {
        if (n <= 1) return n;

        long prev2 = 0; // Fib(i - 2)
        long prev1 = 1; // Fib(i - 1)
        long current = 0;

        for (int i = 2; i <= n; i++) {
            current = prev1 + prev2;
            prev2 = prev1;
            prev1 = current;
        }

        return current;
    }

    // 3. 0/1 Knapsack Problem (Bottom-Up Tabulation)
    public static int knapsack(int[] values, int[] weights, int capacity) {
        int n = values.length;
        // dp[i][w] represents max value using a subset of first i items with weight limit w
        int[][] dp = new int[n + 1][capacity + 1];

        for (int i = 1; i <= n; i++) {
            int val = values[i - 1];
            int wt = weights[i - 1];
            for (int w = 0; w <= capacity; w++) {
                if (wt <= w) {
                    // Max of excluding the item or including the item
                    dp[i][w] = Math.max(dp[i - 1][w], val + dp[i - 1][w - wt]);
                } else {
                    dp[i][w] = dp[i - 1][w]; // Cannot include item
                }
            }
        }
        return dp[n][capacity];
    }
}
```

---

### Go Implementation

```go
package main

import (
	"fmt"
)

// 1. Top-Down Memoization Approach (Fibonacci)
func FibMemo(n int, memo map[int]int64) int64 {
	if n <= 1 {
		return int64(n)
	}
	if val, found := memo[n]; found {
		return val
	}

	result := FibMemo(n-1, memo) + FibMemo(n-2, memo)
	memo[n] = result
	return result
}

// 2. Bottom-Up Tabulation with Space Optimization (Fibonacci)
func FibTabulation(n int) int64 {
	if n <= 1 {
		return int64(n)
	}

	var prev2 int64 = 0 // Fib(i - 2)
	var prev1 int64 = 1 // Fib(i - 1)
	var current int64

	for i := 2; i <= n; i++ {
		current = prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return current
}

// 3. 0/1 Knapsack Problem (Bottom-Up Tabulation)
func Knapsack(values []int, weights []int, capacity int) int {
	n := len(values)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		val := values[i-1]
		wt := weights[i-1]
		for w := 0; w <= capacity; w++ {
			if wt <= w {
				// Max of (excluding item, including item)
				includeVal := val + dp[i-1][w-wt]
				excludeVal := dp[i-1][w]
				if includeVal > excludeVal {
					dp[i][w] = includeVal
				} else {
					dp[i][w] = excludeVal
				}
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[n][capacity]
}
```

---

## 1.6 - Asymptotic Complexity Analysis

* **Fibonacci Memoization / Tabulation:**
  * **Time Complexity:** $\mathcal{O}(n)$ because each of the $n$ distinct subproblems is evaluated exactly once in $\mathcal{O}(1)$ arithmetic time.
  * **Space Complexity:** $\mathcal{O}(n)$ for memo table/recursion stack, reducible to $\mathcal{O}(1)$ auxiliary space using two variables for the rolling window.
* **0/1 Knapsack:**
  * **Time Complexity:** $\mathcal{O}(n \cdot W)$ pseudo-polynomial time where $n$ is the number of items and $W$ is knapsack capacity.
  * **Space Complexity:** $\mathcal{O}(n \cdot W)$ matrix space, reducible to $\mathcal{O}(W)$ using a 1D array traversed in reverse.

---

## 1.7 - Classic LeetCode & CLRS Benchmark Problems

* **Rod Cutting Problem** (CLRS 14.1)
* **Matrix-Chain Multiplication** (CLRS 14.2)
* **Longest Common Subsequence (LCS)** (CLRS 14.4 / LeetCode #1143)
* **Optimal Binary Search Trees** (CLRS 14.5)
* **Climbing Stairs / Fibonacci Number** (LeetCode #70, #509)
* **Coin Change I & II** (LeetCode #322, #518)
* **Longest Increasing Subsequence (LIS)** (LeetCode #300)
* **Edit Distance** (LeetCode #72)

---

## 1.8 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 14: Dynamic Programming (pp. 382–427)
  * Section 14.1: Rod cutting (pp. 383–394)
  * Section 14.2: Matrix-chain multiplication (pp. 395–404)
  * Section 14.3: Elements of dynamic programming (pp. 405–415)
  * Section 14.4: Longest common subsequence (pp. 415–422)
* https://en.wikipedia.org/wiki/Dynamic_programming
* https://www.geeksforgeeks.org/dynamic-programming/
* https://leetcode.com/discuss/general-discussion/458695/dynamic-programming-patterns
