package prefix_sum

// --- NumArray (LeetCode #303) ---
// Immutable prefix-sum structure for range sum queries.
// Time: O(n) build, O(1) query.

type NumArray struct {
	prefix []int
}

// NewNumArray builds the prefix sum array.
func NewNumArray(nums []int) *NumArray {
	prefix := make([]int, len(nums)+1)
	for i, v := range nums {
		prefix[i+1] = prefix[i] + v
	}
	return &NumArray{prefix: prefix}
}

// SumRange returns the sum of nums[left..right] (inclusive, 0-indexed).
func (na *NumArray) SumRange(left, right int) int {
	return na.prefix[right+1] - na.prefix[left]
}

// --- SubarraySum (LeetCode #560) ---
// Counts the number of contiguous subarrays that sum to k.
// Uses prefix sum + hash map to achieve O(n) time.
// Time: O(n), Space: O(n)
func SubarraySum(nums []int, k int) int {
	count := 0
	prefixCount := map[int]int{0: 1}
	sum := 0
	for _, v := range nums {
		sum += v
		count += prefixCount[sum-k]
		prefixCount[sum]++
	}
	return count
}

// --- CheckSubarraySum (LeetCode #523) ---
// Returns true if there exists a subarray of length >= 2 whose sum is a multiple of k.
// Stores the earliest index where each (prefix % k) value was seen.
// Time: O(n), Space: O(k)
func CheckSubarraySum(nums []int, k int) bool {
	// Map from remainder -> earliest index
	remainderIndex := map[int]int{0: -1}
	sum := 0
	for i, v := range nums {
		sum += v
		rem := sum % k
		if idx, ok := remainderIndex[rem]; ok {
			if i-idx >= 2 {
				return true
			}
		} else {
			remainderIndex[rem] = i
		}
	}
	return false
}

// --- CarPooling (LeetCode #1094) ---
// Returns true if all passengers can be transported without exceeding capacity.
// trips[i] = [numPassengers, from, to].
// Uses a difference array (event sweep) on stops 0..1000.
// Time: O(n + S) where S=max stop, Space: O(S)
func CarPooling(trips [][3]int, capacity int) bool {
	diff := [1001]int{}
	for _, trip := range trips {
		num, from, to := trip[0], trip[1], trip[2]
		diff[from] += num
		diff[to] -= num
	}
	cur := 0
	for _, d := range diff {
		cur += d
		if cur > capacity {
			return false
		}
	}
	return true
}
