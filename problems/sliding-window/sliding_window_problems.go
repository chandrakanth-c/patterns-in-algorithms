package sliding_window

// LengthOfLongestSubstring returns the length of the longest substring without
// repeating characters (LeetCode #3).
// Sliding window with a map tracking the last seen index of each character.
// Time: O(n), Space: O(min(n, charset))
func LengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	maxLen := 0
	left := 0
	for right := 0; right < len(s); right++ {
		if idx, ok := lastSeen[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		lastSeen[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// MinWindow returns the minimum window substring of s that contains all characters
// of t (LeetCode #76).
// Sliding window expanding right; contract left when all chars are covered.
// Time: O(|s| + |t|), Space: O(|t|)
func MinWindow(s, t string) string {
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	window := make(map[byte]int)
	have, total := 0, len(need)
	left := 0
	bestLen := len(s) + 1
	bestLeft := 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++
		if need[c] > 0 && window[c] == need[c] {
			have++
		}
		for have == total {
			// Contract from left
			if right-left+1 < bestLen {
				bestLen = right - left + 1
				bestLeft = left
			}
			lc := s[left]
			window[lc]--
			if need[lc] > 0 && window[lc] < need[lc] {
				have--
			}
			left++
		}
	}
	if bestLen > len(s) {
		return ""
	}
	return s[bestLeft : bestLeft+bestLen]
}

// CharacterReplacement returns the length of the longest substring containing
// the same letter after at most k replacements (LeetCode #424).
// Sliding window: track max frequency in current window.
// Time: O(n), Space: O(1)
func CharacterReplacement(s string, k int) int {
	count := [26]int{}
	maxFreq := 0
	left := 0
	maxLen := 0
	for right := 0; right < len(s); right++ {
		count[s[right]-'A']++
		if count[s[right]-'A'] > maxFreq {
			maxFreq = count[s[right]-'A']
		}
		// Window size - maxFreq > k means we need to shrink
		if right-left+1-maxFreq > k {
			count[s[left]-'A']--
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// MaxConsecutiveOnes returns the maximum number of consecutive 1s after flipping
// at most k zeros (LeetCode #1004).
// Sliding window counting zeros in window.
// Time: O(n), Space: O(1)
func MaxConsecutiveOnes(nums []int, k int) int {
	left := 0
	zeros := 0
	maxLen := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// NumSubarrayProductLessThanK returns the number of subarrays where the product
// of all elements is strictly less than k (LeetCode #713).
// Sliding window: count subarrays ending at 'right'.
// Time: O(n), Space: O(1)
func NumSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	count := 0
	prod := 1
	left := 0
	for right := 0; right < len(nums); right++ {
		prod *= nums[right]
		for prod >= k {
			prod /= nums[left]
			left++
		}
		count += right - left + 1
	}
	return count
}

// MaxSlidingWindow returns the max value in each sliding window of size k (LeetCode #239).
// Monotonic deque: front always holds index of the current window maximum.
// Time: O(n), Space: O(k)
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	deque := []int{} // stores indices, decreasing by nums value
	result := []int{}
	for i := 0; i < len(nums); i++ {
		// Remove indices outside the window
		for len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}
		// Maintain decreasing order: remove smaller elements from back
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}
