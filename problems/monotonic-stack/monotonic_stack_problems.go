package monotonic_stack

// --- DailyTemperatures (LeetCode #739) ---
// For each day, find how many days until a warmer temperature.
// Uses a monotonic decreasing stack of indices.
// Time: O(n), Space: O(n)
func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := []int{} // stack of indices
	for i, t := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < t {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = i - idx
		}
		stack = append(stack, i)
	}
	return result
}

// --- NextGreaterElement (LeetCode #496) ---
// For each element in nums1, find the next greater element in nums2.
// Uses a monotonic stack + hash map for O(n+m) lookup.
// Time: O(n+m), Space: O(n+m)
func NextGreaterElement(nums1, nums2 []int) []int {
	nextGreater := make(map[int]int)
	stack := []int{}
	for _, n := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < n {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextGreater[top] = n
		}
		stack = append(stack, n)
	}
	// Remaining elements in stack have no greater element.
	for _, n := range stack {
		nextGreater[n] = -1
	}
	result := make([]int, len(nums1))
	for i, n := range nums1 {
		result[i] = nextGreater[n]
	}
	return result
}

// --- LargestRectangleArea (LeetCode #84) ---
// Finds the largest rectangle that can be formed in a histogram.
// Uses a monotonic increasing stack; sentinel bar of height 0 at the end.
// Time: O(n), Space: O(n)
func LargestRectangleArea(heights []int) int {
	stack := []int{}   // indices, monotonic increasing by height
	maxArea := 0
	heights = append(heights, 0) // sentinel
	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			if area := height * width; area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}

// --- OnlineStockSpan (LeetCode #901) ---
// Returns a closure that computes the stock span for each price.
// The span of a stock's price on a given day is the max number of consecutive
// days (up to and including today) for which the price was <= today's price.
// Uses a monotonic decreasing stack of (price, span) pairs.
// Time: O(1) amortized per call, Space: O(n)
func OnlineStockSpan() func(price int) int {
	type pair struct{ price, span int }
	stack := []pair{}
	return func(price int) int {
		span := 1
		for len(stack) > 0 && stack[len(stack)-1].price <= price {
			span += stack[len(stack)-1].span
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{price, span})
		return span
	}
}
