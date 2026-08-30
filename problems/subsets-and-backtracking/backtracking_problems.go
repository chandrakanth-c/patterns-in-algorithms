package backtracking

import "sort"

// =============================================================
// Subsets — LeetCode #78
// Generate all 2^n subsets using iterative bit-masking approach.
// Time: O(n·2^n), Space: O(n·2^n).
// =============================================================

// Subsets returns all subsets of nums (no duplicates in input).
func Subsets(nums []int) [][]int {
	result := [][]int{}
	var dfs func(start int, current []int)
	dfs = func(start int, current []int) {
		tmp := make([]int, len(current))
		copy(tmp, current)
		result = append(result, tmp)
		for i := start; i < len(nums); i++ {
			dfs(i+1, append(current, nums[i]))
		}
	}
	dfs(0, []int{})
	return result
}

// =============================================================
// SubsetsWithDup — LeetCode #90
// Sort first, then skip duplicates at the same recursion depth.
// Time: O(n·2^n).
// =============================================================

// SubsetsWithDup returns all unique subsets (input may have duplicates).
func SubsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	var dfs func(start int, current []int)
	dfs = func(start int, current []int) {
		tmp := make([]int, len(current))
		copy(tmp, current)
		result = append(result, tmp)
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue // skip duplicate at same level
			}
			dfs(i+1, append(current, nums[i]))
		}
	}
	dfs(0, []int{})
	return result
}

// =============================================================
// Permute — LeetCode #46
// Classic backtracking swap-based permutation generation.
// Time: O(n·n!).
// =============================================================

// Permute returns all permutations of distinct nums.
func Permute(nums []int) [][]int {
	result := [][]int{}
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			result = append(result, tmp)
			return
		}
		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			backtrack(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	backtrack(0)
	return result
}

// =============================================================
// PermuteUnique — LeetCode #47
// Sort and use a "used" boolean array to skip duplicates.
// Time: O(n·n!).
// =============================================================

// PermuteUnique returns all unique permutations (input may have duplicates).
func PermuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	used := make([]bool, len(nums))
	var backtrack func(current []int)
	backtrack = func(current []int) {
		if len(current) == len(nums) {
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// Skip duplicate: same value as previous, previous not used
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			backtrack(append(current, nums[i]))
			used[i] = false
		}
	}
	backtrack([]int{})
	return result
}

// =============================================================
// CombinationSum — LeetCode #39
// Backtrack with the same element allowed multiple times.
// Time: O(n^(t/m)) where t=target, m=min candidate.
// =============================================================

// CombinationSum finds all unique combinations that sum to target.
func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	var backtrack func(start, remain int, current []int)
	backtrack = func(start, remain int, current []int) {
		if remain == 0 {
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > remain {
				break
			}
			backtrack(i, remain-candidates[i], append(current, candidates[i]))
		}
	}
	backtrack(0, target, []int{})
	return result
}

// =============================================================
// SolveNQueens — LeetCode #51
// Backtrack column by column; use three sets to track conflicts.
// Time: O(n!), Space: O(n).
// =============================================================

// SolveNQueens returns all distinct solutions to the n-queens puzzle.
func SolveNQueens(n int) [][]string {
	result := [][]string{}
	board := make([]int, n) // board[row] = column of queen in that row
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n)  // row-col+n
	diag2 := make([]bool, 2*n)  // row+col

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			// Build the string board.
			solution := make([]string, n)
			for r := 0; r < n; r++ {
				b := make([]byte, n)
				for c := 0; c < n; c++ {
					b[c] = '.'
				}
				b[board[r]] = 'Q'
				solution[r] = string(b)
			}
			result = append(result, solution)
			return
		}
		for col := 0; col < n; col++ {
			d1 := row - col + n
			d2 := row + col
			if cols[col] || diag1[d1] || diag2[d2] {
				continue
			}
			cols[col] = true
			diag1[d1] = true
			diag2[d2] = true
			board[row] = col
			backtrack(row + 1)
			cols[col] = false
			diag1[d1] = false
			diag2[d2] = false
		}
	}
	backtrack(0)
	return result
}
