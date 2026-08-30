package binary_search

// --- SearchRotated (LeetCode #33) ---
// Searches a rotated sorted array for target. Returns index or -1.
// Key insight: one half of the array is always sorted.
// Time: O(log n), Space: O(1)
func SearchRotated(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return mid
		}
		// Left half is sorted.
		if nums[lo] <= nums[mid] {
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			// Right half is sorted.
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

// --- FindMinRotated (LeetCode #153) ---
// Finds the minimum element in a rotated sorted array (no duplicates).
// Time: O(log n), Space: O(1)
func FindMinRotated(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] > nums[hi] {
			lo = mid + 1 // min is in the right half
		} else {
			hi = mid // mid could be the min
		}
	}
	return nums[lo]
}

// --- MinEatingSpeed (LeetCode #875) ---
// Koko eats piles of bananas; finds minimum speed k (bananas/hour) to finish
// all piles in h hours. Binary search on the answer space [1, max(piles)].
// Time: O(n log(max)), Space: O(1)
func MinEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, p := range piles {
		if p > maxPile {
			maxPile = p
		}
	}
	lo, hi := 1, maxPile
	for lo < hi {
		mid := (lo + hi) / 2
		hours := 0
		for _, p := range piles {
			hours += (p + mid - 1) / mid // ceil(p/mid)
		}
		if hours <= h {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// --- ShipWithinDays (LeetCode #1011) ---
// Finds minimum ship capacity to ship all weights within days.
// Binary search on capacity in [max(weights), sum(weights)].
// Time: O(n log(sum)), Space: O(1)
func ShipWithinDays(weights []int, days int) int {
	lo, hi := 0, 0
	for _, w := range weights {
		if w > lo {
			lo = w
		}
		hi += w
	}
	for lo < hi {
		mid := (lo + hi) / 2
		need, cur := 1, 0
		for _, w := range weights {
			if cur+w > mid {
				need++
				cur = 0
			}
			cur += w
		}
		if need <= days {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
