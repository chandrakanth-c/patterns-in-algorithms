package bit_manipulation

// 1. SingleNumber returns the element that appears only once; others appear twice (LeetCode #136)
// XOR all numbers: duplicates cancel out.
func SingleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result ^= n
	}
	return result
}

// 2. SingleNumberII returns the element appearing once; others appear three times (LeetCode #137)
// Count bits modulo 3 using two bitmasks.
func SingleNumberII(nums []int) int {
	ones, twos := 0, 0
	for _, n := range nums {
		ones = (ones ^ n) & ^twos
		twos = (twos ^ n) & ^ones
	}
	return ones
}

// 3. SingleNumberIII returns two elements each appearing once; others appear twice (LeetCode #260)
// XOR all to get xor of the two uniques, then split by a differing bit.
func SingleNumberIII(nums []int) [2]int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}
	// Isolate rightmost set bit
	diff := xor & (-xor)
	a := 0
	for _, n := range nums {
		if n&diff != 0 {
			a ^= n
		}
	}
	b := xor ^ a
	if a < b {
		return [2]int{a, b}
	}
	return [2]int{b, a}
}

// 4. HammingWeight counts the number of set bits (LeetCode #191)
// Brian Kernighan's algorithm: repeatedly clear the lowest set bit.
func HammingWeight(n uint32) int {
	count := 0
	for n != 0 {
		n &= n - 1
		count++
	}
	return count
}

// 5. CountBits returns arr where arr[i] = number of 1-bits in i (LeetCode #338)
// DP: dp[i] = dp[i >> 1] + (i & 1)
func CountBits(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i>>1] + (i & 1)
	}
	return dp
}

// 6. SubsetsBitmask generates all subsets of nums using bitmask enumeration.
// For n elements, iterate mask from 0 to 2^n-1.
func SubsetsBitmask(nums []int) [][]int {
	n := len(nums)
	total := 1 << n
	result := make([][]int, 0, total)
	for mask := 0; mask < total; mask++ {
		subset := []int{}
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				subset = append(subset, nums[i])
			}
		}
		result = append(result, subset)
	}
	return result
}
