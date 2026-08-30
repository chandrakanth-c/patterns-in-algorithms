package dp

// CutRod returns the maximum revenue obtainable by cutting up a rod of length n given prices p
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

// MatrixChainOrder finds the minimum number of scalar multiplications needed to multiply a chain of matrices
func MatrixChainOrder(p []int) int {
	n := len(p) - 1 // number of matrices
	m := make([][]int, n+1)
	for i := range m {
		m[i] = make([]int, n+1)
	}

	for l := 2; l <= n; l++ { // l is chain length
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

// LongestCommonSubsequence computes the length of the longest common subsequence of text1 and text2
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

// OptimalBST computes the expected search cost of an optimal binary search tree
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

// ClimbStairs returns number of distinct ways to climb n stairs taking 1 or 2 steps
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

// CoinChange finds minimum number of coins needed to make up amount (-1 if impossible)
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

// LengthOfLIS finds the length of the longest strictly increasing subsequence
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

// MinDistance finds minimum edit distance between word1 and word2
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
