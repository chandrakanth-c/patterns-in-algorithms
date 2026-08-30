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

## 1.5 - Classic Paradigm Example: Fibonacci Sequence & 0/1 Knapsack

### Mathematical Formulation
$$\text{Fib}(n) = \begin{cases} 0 & \text{if } n = 0 \\ 1 & \text{if } n = 1 \\ \text{Fib}(n-1) + \text{Fib}(n-2) & \text{if } n \ge 2 \end{cases}$$

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

> Standalone runnable code and unit tests are located in:
> * [Java Solutions & Tests](problems/dynamic-programming/DPProblems.java) / [Test Runner](problems/dynamic-programming/DPProblemsTest.java)
> * [Go Solutions](problems/dynamic-programming/dp_problems.go) / [Go Tests](problems/dynamic-programming/dp_problems_test.go)

---

### 1.7.1 - Rod Cutting Problem (CLRS 14.1)

#### 1. Problem Statement
Given a rod of length $n$ inches and a table of prices $p_i$ for $i = 1, 2, \dots, n$, determine the maximum revenue $r_n$ obtainable by cutting up the rod and selling the pieces. You may make any number of cuts (including zero cuts).

#### 2. Conceptual Link to Dynamic Programming
* **Optimal Substructure:** An optimal cut of length $n$ consists of a first piece of length $i$ ($1 \le i \le n$) sold at price $p_i$, plus an independently optimal cut of the remaining rod of length $n - i$.
* **Recurrence:**
  $$r_n = \max_{1 \le i \le n} (p_i + r_{n-i}), \quad \text{with } r_0 = 0$$

#### 3. Java Implementation
```java
public static int cutRod(int[] p, int n) {
    int[] r = new int[n + 1];
    r[0] = 0;
    for (int j = 1; j <= n; j++) {
        int q = Integer.MIN_VALUE;
        for (int i = 1; i <= j; i++) {
            q = Math.max(q, p[i - 1] + r[j - i]);
        }
        r[j] = q;
    }
    return r[n];
}
```

#### 4. Go Implementation
```go
func CutRod(p []int, n int) int {
	r := make([]int, n+1)
	r[0] = 0
	for j := 1; j <= n; j++ {
		q := -1
		for i := 1; i <= j; i++ {
			val := p[i-1] + r[j-i]
			if val > q {
				q = val
			}
		}
		r[j] = q
	}
	return r[n]
}
```

#### 5. Complexity Analysis
* **Time Complexity:** $\Theta(n^2)$ — Nested loops run $\sum_{j=1}^n j = \frac{n(n+1)}{2}$ steps.
* **Space Complexity:** $\Theta(n)$ — 1D table $r$ of size $n+1$.

---

### 1.7.2 - Matrix-Chain Multiplication (CLRS 14.2)

#### 1. Problem Statement
Given a sequence (chain) $\langle A_1, A_2, \dots, A_n \rangle$ of $n$ matrices, where matrix $A_i$ has dimension $p_{i-1} \times p_i$, fully parenthesize the product $A_1 A_2 \dots A_n$ in a way that minimizes the number of scalar multiplications.

#### 2. Conceptual Link to Dynamic Programming
* **Optimal Substructure:** Any optimal parenthesization of $A_i \dots A_j$ splits the product at some $k$ ($i \le k < j$) into $(A_i \dots A_k)(A_{k+1} \dots A_j)$, where each subchain must itself be optimally parenthesized.
* **Recurrence:**
  $$m[i, j] = \begin{cases} 0 & \text{if } i = j \\ \min_{i \le k < j} \{m[i, k] + m[k+1, j] + p_{i-1} p_k p_j\} & \text{if } i < j \end{cases}$$

#### 3. Java Implementation
```java
public static int matrixChainOrder(int[] p) {
    int n = p.length - 1;
    int[][] m = new int[n + 1][n + 1];
    for (int l = 2; l <= n; l++) { // chain length
        for (int i = 1; i <= n - l + 1; i++) {
            int j = i + l - 1;
            m[i][j] = Integer.MAX_VALUE;
            for (int k = i; k <= j - 1; k++) {
                int q = m[i][k] + m[k + 1][j] + p[i - 1] * p[k] * p[j];
                if (q < m[i][j]) m[i][j] = q;
            }
        }
    }
    return m[1][n];
}
```

#### 4. Go Implementation
```go
func MatrixChainOrder(p []int) int {
	n := len(p) - 1
	m := make([][]int, n+1)
	for i := range m {
		m[i] = make([]int, n+1)
	}
	for l := 2; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			m[i][j] = 1<<31 - 1
			for k := i; k <= j-1; k++ {
				q := m[i][k] + m[k+1][j] + p[i-1]*p[k]*p[j]
				if q < m[i][j] {
					m[i][j] = q
				}
			}
		}
	}
	return m[1][n]
}
```

#### 5. Complexity Analysis
* **Time Complexity:** $\mathcal{O}(n^3)$ — 3 nested loops (chain length $l$, starting index $i$, split position $k$).
* **Space Complexity:** $\Theta(n^2)$ — 2D matrix $m$ of size $(n+1) \times (n+1)$.

---

### 1.7.3 - Longest Common Subsequence (LCS) (CLRS 14.4 / LeetCode #1143)

#### 1. Problem Statement
Given two strings `text1` and `text2`, return the length of their longest common subsequence. A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

#### 2. Conceptual Link to Dynamic Programming
* **Optimal Substructure:** If $X[m] == Y[n]$, the last character must be in the LCS $\implies 1 + \text{LCS}(X_{m-1}, Y_{n-1})$. If $X[m] \ne Y[n]$, we solve two subproblems and take the max: $\max(\text{LCS}(X_{m-1}, Y_n), \text{LCS}(X_m, Y_{n-1}))$.

#### 3. Java Implementation
```java
public static int longestCommonSubsequence(String text1, String text2) {
    int m = text1.length(), n = text2.length();
    int[][] dp = new int[m + 1][n + 1];
    for (int i = 1; i <= m; i++) {
        for (int j = 1; j <= n; j++) {
            if (text1.charAt(i - 1) == text2.charAt(j - 1)) {
                dp[i][j] = dp[i - 1][j - 1] + 1;
            } else {
                dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1]);
            }
        }
    }
    return dp[m][n];
}
```

#### 4. Go Implementation
```go
func LongestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[m][n]
}
```

#### 5. Complexity Analysis
* **Time Complexity:** $\Theta(m \cdot n)$ — Tabulates an $(m+1) \times (n+1)$ matrix.
* **Space Complexity:** $\Theta(m \cdot n)$ table space, reducible to $\mathcal{O}(\min(m, n))$ using two rows.

---

### 1.7.4 - Optimal Binary Search Trees (CLRS 14.5)

#### 1. Problem Statement
Given a sequence $K = \langle k_1, k_2, \dots, k_n \rangle$ of $n$ distinct sorted keys with search probabilities $p_i$, and dummy keys $d_0, d_1, \dots, d_n$ representing values outside $K$ with search probabilities $q_i$, construct a binary search tree that minimizes expected search cost.

#### 2. Conceptual Link to Dynamic Programming
* **Optimal Substructure:** If an optimal BST has root $k_r$, its left subtree must be an optimal BST for keys $k_i \dots k_{r-1}$ and its right subtree an optimal BST for keys $k_{r+1} \dots k_j$.
* **Recurrence:**
  $$e[i, j] = \begin{cases} q_{i-1} & \text{if } j = i - 1 \\ \min_{i \le r \le j} \{e[i, r-1] + e[r+1, j] + w(i, j)\} & \text{if } i \le j \end{cases}$$

#### 3. Java Implementation
```java
public static double optimalBST(double[] p, double[] q, int n) {
    double[][] e = new double[n + 2][n + 2];
    double[][] w = new double[n + 2][n + 2];
    for (int i = 1; i <= n + 1; i++) {
        e[i][i - 1] = q[i - 1];
        w[i][i - 1] = q[i - 1];
    }
    for (int l = 1; l <= n; l++) {
        for (int i = 1; i <= n - l + 1; i++) {
            int j = i + l - 1;
            e[i][j] = Double.MAX_VALUE;
            w[i][j] = w[i][j - 1] + p[j - 1] + q[j];
            for (int r = i; r <= j; r++) {
                double t = e[i][r - 1] + e[r + 1][j] + w[i][j];
                if (t < e[i][j]) e[i][j] = t;
            }
        }
    }
    return e[1][n];
}
```

#### 4. Go Implementation
```go
func OptimalBST(p []float64, q []float64, n int) float64 {
	e := make([][]float64, n+2)
	w := make([][]float64, n+2)
	for i := range e {
		e[i] = make([]float64, n+2)
		w[i] = make([]float64, n+2)
	}
	for i := 1; i <= n+1; i++ {
		e[i][i-1] = q[i-1]
		w[i][i-1] = q[i-1]
	}
	for l := 1; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			e[i][j] = 1e9
			w[i][j] = w[i][j-1] + p[j-1] + q[j]
			for r := i; r <= j; r++ {
				t := e[i][r-1] + e[r+1][j] + w[i][j]
				if t < e[i][j] {
					e[i][j] = t
				}
			}
		}
	}
	return e[1][n]
}
```

#### 5. Complexity Analysis
* **Time Complexity:** $\Theta(n^3)$ — 3 nested loops over subproblem span $l$, starting index $i$, and root candidate $r$.
* **Space Complexity:** $\Theta(n^2)$ — 2D tables $e$ and $w$.

---

### 1.7.5 - Climbing Stairs (LeetCode #70)

#### 1. Problem Statement
You are climbing a staircase. It takes $n$ steps to reach the top. Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

#### 2. Conceptual Link to Dynamic Programming
* **Optimal Substructure:** To reach step $n$, you must come from either step $n-1$ or step $n-2$.
* **Recurrence:** $\text{dp}[n] = \text{dp}[n-1] + \text{dp}[n-2]$ with base cases $\text{dp}[1] = 1, \text{dp}[2] = 2$.

#### 3. Java Implementation
```java
public static int climbStairs(int n) {
    if (n <= 2) return n;
    int prev2 = 1, prev1 = 2;
    for (int i = 3; i <= n; i++) {
        int curr = prev1 + prev2;
        prev2 = prev1;
        prev1 = curr;
    }
    return prev1;
}
```

#### 4. Go Implementation
```go
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prev2, prev1 := 1, 2
	for i := 3; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}
	return prev1
}
```

#### 5. Complexity Analysis
* **Time Complexity:** $\mathcal{O}(n)$ — Single linear pass.
* **Space Complexity:** $\mathcal{O}(1)$ — Constant auxiliary variables.

---

### 1.7.6 - Coin Change (LeetCode #322)

#### 1. Problem Statement
Given an integer array `coins` representing coins of different denominations and an integer `amount`, return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return `-1`.

#### 2. Conceptual Link to Dynamic Programming
* **Optimal Substructure:** The minimum coins to make amount $a$ is $1 + \min_{c \in \text{coins}} \text{dp}[a - c]$.
* **Recurrence:**
  $$\text{dp}[a] = \min_{c \in \text{coins}, c \le a} (\text{dp}[a - c] + 1), \quad \text{with } \text{dp}[0] = 0$$

#### 3. Java Implementation
```java
public static int coinChange(int[] coins, int amount) {
    int maxVal = amount + 1;
    int[] dp = new int[amount + 1];
    java.util.Arrays.fill(dp, maxVal);
    dp[0] = 0;
    for (int i = 1; i <= amount; i++) {
        for (int coin : coins) {
            if (coin <= i) {
                dp[i] = Math.min(dp[i], dp[i - coin] + 1);
            }
        }
    }
    return dp[amount] > amount ? -1 : dp[amount];
}
```

#### 4. Go Implementation
```go
func CoinChange(coins []int, amount int) int {
	maxVal := amount + 1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = maxVal
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
```

#### 5. Complexity Analysis
* **Time Complexity:** $\mathcal{O}(\text{amount} \cdot |\text{coins}|)$ — For each amount up to target, iterates through all coin denominations.
* **Space Complexity:** $\mathcal{O}(\text{amount})$ — 1D DP table.

---

### 1.7.7 - Longest Increasing Subsequence (LIS) (LeetCode #300)

#### 1. Problem Statement
Given an integer array `nums`, return the length of the longest strictly increasing subsequence.

#### 2. Conceptual Link to Dynamic Programming
* **Optimal Substructure:** Let $\text{dp}[i]$ be the length of the longest increasing subsequence ending at index $i$.
* **Recurrence:**
  $$\text{dp}[i] = 1 + \max_{0 \le j < i, \text{nums}[j] < \text{nums}[i]} \text{dp}[j]$$

#### 3. Java Implementation
```java
public static int lengthOfLIS(int[] nums) {
    if (nums == null || nums.length == 0) return 0;
    int[] dp = new int[nums.length];
    int maxLength = 1;
    for (int i = 0; i < nums.length; i++) {
        dp[i] = 1;
        for (int j = 0; j < i; j++) {
            if (nums[i] > nums[j]) {
                dp[i] = Math.max(dp[i], dp[j] + 1);
            }
        }
        maxLength = Math.max(maxLength, dp[i]);
    }
    return maxLength;
}
```

#### 4. Go Implementation
```go
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	maxLength := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxLength {
			maxLength = dp[i]
		}
	}
	return maxLength
}
```

#### 5. Complexity Analysis
* **Time Complexity:** $\mathcal{O}(n^2)$ standard dynamic programming (can be optimized to $\mathcal{O}(n \log n)$ via Patience Sorting + Binary Search).
* **Space Complexity:** $\mathcal{O}(n)$ to store subsequence lengths.

---

### 1.7.8 - Edit Distance (LeetCode #72)

#### 1. Problem Statement
Given two strings `word1` and `word2`, return the minimum number of operations required to convert `word1` to `word2`. You have the following three operations permitted on a word: insert a character, delete a character, or replace a character.

#### 2. Conceptual Link to Dynamic Programming
* **Optimal Substructure:** Let $\text{dp}[i][j]$ be the edit distance between `word1[0..i-1]` and `word2[0..j-1]`.
* **Recurrence:**
  $$\text{dp}[i][j] = \begin{cases} \text{dp}[i-1][j-1] & \text{if } \text{word1}[i-1] == \text{word2}[j-1] \\ 1 + \min(\text{dp}[i][j-1], \text{dp}[i-1][j], \text{dp}[i-1][j-1]) & \text{otherwise (insert, delete, replace)} \end{cases}$$

#### 3. Java Implementation
```java
public static int minDistance(String word1, String word2) {
    int m = word1.length(), n = word2.length();
    int[][] dp = new int[m + 1][n + 1];
    for (int i = 0; i <= m; i++) dp[i][0] = i;
    for (int j = 0; j <= n; j++) dp[0][j] = j;
    for (int i = 1; i <= m; i++) {
        for (int j = 1; j <= n; j++) {
            if (word1.charAt(i - 1) == word2.charAt(j - 1)) {
                dp[i][j] = dp[i - 1][j - 1];
            } else {
                dp[i][j] = 1 + Math.min(dp[i - 1][j - 1], Math.min(dp[i - 1][j], dp[i][j - 1]));
            }
        }
    }
    return dp[m][n];
}
```

#### 4. Go Implementation
```go
func MinDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				insertOp := dp[i][j-1]
				deleteOp := dp[i-1][j]
				replaceOp := dp[i-1][j-1]
				minOp := insertOp
				if deleteOp < minOp {
					minOp = deleteOp
				}
				if replaceOp < minOp {
					minOp = replaceOp
				}
				dp[i][j] = 1 + minOp
			}
		}
	}
	return dp[m][n]
}
```

#### 5. Complexity Analysis
* **Time Complexity:** $\Theta(m \cdot n)$ — Standard 2D dynamic programming grid.
* **Space Complexity:** $\Theta(m \cdot n)$ matrix space.

---

## 1.8 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 14: Dynamic Programming (pp. 382–427)
  * Section 14.1: Rod cutting (pp. 383–394)
  * Section 14.2: Matrix-chain multiplication (pp. 395–404)
  * Section 14.3: Elements of dynamic programming (pp. 405–415)
  * Section 14.4: Longest common subsequence (pp. 415–422)
  * Section 14.5: Optimal binary search trees (pp. 422–427)
* https://en.wikipedia.org/wiki/Dynamic_programming
* https://leetcode.com/problems/longest-common-subsequence/
* https://leetcode.com/problems/coin-change/
* https://leetcode.com/problems/edit-distance/
