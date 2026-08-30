package two_pointers

import (
	"sort"
	"unicode"
)

// TwoSumII returns 1-indexed positions of two numbers that add to target (LeetCode #167).
// Array is already sorted; shrink window from both ends.
// Time: O(n), Space: O(1)
func TwoSumII(numbers []int, target int) [2]int {
	lo, hi := 0, len(numbers)-1
	for lo < hi {
		sum := numbers[lo] + numbers[hi]
		switch {
		case sum == target:
			return [2]int{lo + 1, hi + 1}
		case sum < target:
			lo++
		default:
			hi--
		}
	}
	return [2]int{-1, -1}
}

// ThreeSum finds all unique triplets in nums that sum to zero (LeetCode #15).
// Sort, then for each element use two-pointer on the remainder.
// Time: O(n²), Space: O(1) extra
func ThreeSum(nums []int) [][3]int {
	sort.Ints(nums)
	result := [][3]int{}
	n := len(nums)
	for i := 0; i < n-2; i++ {
		// Skip duplicate values for the first element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		lo, hi := i+1, n-1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]
			switch {
			case sum == 0:
				result = append(result, [3]int{nums[i], nums[lo], nums[hi]})
				for lo < hi && nums[lo] == nums[lo+1] {
					lo++
				}
				for lo < hi && nums[hi] == nums[hi-1] {
					hi--
				}
				lo++
				hi--
			case sum < 0:
				lo++
			default:
				hi--
			}
		}
	}
	return result
}

// ContainerWithMostWater finds two lines that together with the x-axis forms a container
// holding the most water (LeetCode #11).
// Time: O(n), Space: O(1)
func ContainerWithMostWater(height []int) int {
	lo, hi := 0, len(height)-1
	maxWater := 0
	for lo < hi {
		h := height[lo]
		if height[hi] < h {
			h = height[hi]
		}
		water := h * (hi - lo)
		if water > maxWater {
			maxWater = water
		}
		if height[lo] < height[hi] {
			lo++
		} else {
			hi--
		}
	}
	return maxWater
}

// ValidPalindrome checks whether s is a palindrome, considering only alphanumerics
// and ignoring case (LeetCode #125).
// Time: O(n), Space: O(1)
func ValidPalindrome(s string) bool {
	runes := []rune(s)
	lo, hi := 0, len(runes)-1
	for lo < hi {
		for lo < hi && !unicode.IsLetter(runes[lo]) && !unicode.IsDigit(runes[lo]) {
			lo++
		}
		for lo < hi && !unicode.IsLetter(runes[hi]) && !unicode.IsDigit(runes[hi]) {
			hi--
		}
		if unicode.ToLower(runes[lo]) != unicode.ToLower(runes[hi]) {
			return false
		}
		lo++
		hi--
	}
	return true
}

// TrappingRainWater computes total water trapped after raining (LeetCode #42).
// Two-pointer approach: track maxLeft, maxRight and accumulate trapped water.
// Time: O(n), Space: O(1)
func TrappingRainWater(height []int) int {
	lo, hi := 0, len(height)-1
	maxLeft, maxRight := 0, 0
	water := 0
	for lo < hi {
		if height[lo] < height[hi] {
			if height[lo] >= maxLeft {
				maxLeft = height[lo]
			} else {
				water += maxLeft - height[lo]
			}
			lo++
		} else {
			if height[hi] >= maxRight {
				maxRight = height[hi]
			} else {
				water += maxRight - height[hi]
			}
			hi--
		}
	}
	return water
}
