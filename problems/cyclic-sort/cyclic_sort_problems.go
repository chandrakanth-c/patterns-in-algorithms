package cyclic_sort

// MissingNumber finds the missing number in nums containing n distinct numbers in [0, n]
// (LeetCode #268).
// Cyclic sort places each number at its correct index (value == index for 0-based).
// Time: O(n), Space: O(1)
func MissingNumber(nums []int) int {
	n := len(nums)
	i := 0
	for i < n {
		j := nums[i] // correct index for nums[i] is nums[i] itself (0-based)
		if nums[i] < n && nums[i] != i {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	for i, v := range nums {
		if v != i {
			return i
		}
	}
	return n
}

// FindDisappearedNumbers returns all numbers in [1,n] missing from nums (LeetCode #448).
// Cyclic sort places each nums[i] at index nums[i]-1.
// Time: O(n), Space: O(1) extra
func FindDisappearedNumbers(nums []int) []int {
	n := len(nums)
	i := 0
	for i < n {
		j := nums[i] - 1 // correct index for nums[i] (1-based)
		if nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	result := []int{}
	for i, v := range nums {
		if v != i+1 {
			result = append(result, i+1)
		}
	}
	return result
}

// FindDuplicate finds the duplicate in nums[1..n] using cyclic sort (LeetCode #287).
// Place each number at nums[i]-1; when nums[i]==nums[nums[i]-1] and i!=nums[i]-1, found it.
// Time: O(n), Space: O(1) — modifies input array
func FindDuplicate(nums []int) int {
	i := 0
	for i < len(nums) {
		j := nums[i] - 1
		if nums[i] != i+1 {
			if nums[i] != nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			} else {
				// Duplicate found
				return nums[i]
			}
		} else {
			i++
		}
	}
	return -1
}

// FirstMissingPositive returns the smallest missing positive integer (LeetCode #41).
// Cyclic sort: place each number x in [1,n] at index x-1; then scan for mismatch.
// Time: O(n), Space: O(1)
func FirstMissingPositive(nums []int) int {
	n := len(nums)
	i := 0
	for i < n {
		j := nums[i] - 1
		if nums[i] > 0 && nums[i] <= n && nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}
	return n + 1
}
