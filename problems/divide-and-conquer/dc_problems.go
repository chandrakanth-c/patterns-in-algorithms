package dc

import (
	"math"
	"sort"
)

// 1. MaxSubarray finds contiguous subarray with largest sum (CLRS 4.1 / LeetCode #53)
func MaxSubarray(nums []int) int {
	var helper func(low, high int) int
	helper = func(low, high int) int {
		if low == high {
			return nums[low]
		}
		mid := low + (high-low)/2
		leftMax := helper(low, mid)
		rightMax := helper(mid+1, high)
		crossMax := maxCrossingSum(nums, low, mid, high)

		maxVal := leftMax
		if rightMax > maxVal {
			maxVal = rightMax
		}
		if crossMax > maxVal {
			maxVal = crossMax
		}
		return maxVal
	}
	return helper(0, len(nums)-1)
}

func maxCrossingSum(nums []int, low, mid, high int) int {
	leftSum := -1 << 30
	sum := 0
	for i := mid; i >= low; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	rightSum := -1 << 30
	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	return leftSum + rightSum
}

// 2. StrassenMultiply multiplies two 2x2 matrices using Strassen's 7 multiplications (CLRS 4.2)
func StrassenMultiply2x2(A, B [][]int) [][]int {
	p1 := A[0][0] * (B[0][1] - B[1][1])
	p2 := (A[0][0] + A[0][1]) * B[1][1]
	p3 := (A[1][0] + A[1][1]) * B[0][0]
	p4 := A[1][1] * (B[1][0] - B[0][0])
	p5 := (A[0][0] + A[1][1]) * (B[0][0] + B[1][1])
	p6 := (A[0][1] - A[1][1]) * (B[1][0] + B[1][1])
	p7 := (A[0][0] - A[1][0]) * (B[0][0] + B[0][1])

	C := make([][]int, 2)
	for i := range C {
		C[i] = make([]int, 2)
	}

	C[0][0] = p5 + p4 - p2 + p6
	C[0][1] = p1 + p2
	C[1][0] = p3 + p4
	C[1][1] = p5 + p1 - p3 - p7
	return C
}

// 3. QuickSort (CLRS Chapter 7 / LeetCode #912)
func QuickSort(nums []int) {
	var sortHelper func(p, r int)
	sortHelper = func(p, r int) {
		if p < r {
			q := partition(nums, p, r)
			sortHelper(p, q-1)
			sortHelper(q+1, r)
		}
	}
	sortHelper(0, len(nums)-1)
}

func partition(nums []int, p, r int) int {
	x := nums[r] // Pivot
	i := p - 1
	for j := p; j < r; j++ {
		if nums[j] <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

// 4. RandomizedSelect finds the k-th smallest element (1-based) in expected linear time (CLRS 9.2)
func RandomizedSelect(nums []int, p, r, k int) int {
	if p == r {
		return nums[p]
	}
	q := partition(nums, p, r)
	rank := q - p + 1
	if k == rank {
		return nums[q]
	} else if k < rank {
		return RandomizedSelect(nums, p, q-1, k)
	} else {
		return RandomizedSelect(nums, q+1, r, k-rank)
	}
}

// 5. ClosestPairOfPoints finds distance between closest pair of 2D points in O(n log n) (CLRS Chapter 33)
type Point struct {
	X, Y float64
}

func ClosestPair(points []Point) float64 {
	pts := make([]Point, len(points))
	copy(pts, points)
	sort.Slice(pts, func(i, j int) bool {
		return pts[i].X < pts[j].X
	})

	var closestRec func(left, right int) float64
	closestRec = func(left, right int) float64 {
		if right-left <= 3 {
			minDist := math.MaxFloat64
			for i := left; i <= right; i++ {
				for j := i + 1; j <= right; j++ {
					d := dist(pts[i], pts[j])
					if d < minDist {
						minDist = d
					}
				}
			}
			return minDist
		}

		mid := left + (right-left)/2
		midX := pts[mid].X
		d1 := closestRec(left, mid)
		d2 := closestRec(mid+1, right)
		d := math.Min(d1, d2)

		// Strip within distance d of dividing line
		var strip []Point
		for i := left; i <= right; i++ {
			if math.Abs(pts[i].X-midX) < d {
				strip = append(strip, pts[i])
			}
		}

		sort.Slice(strip, func(i, j int) bool {
			return strip[i].Y < strip[j].Y
		})

		for i := 0; i < len(strip); i++ {
			for j := i + 1; j < len(strip) && (strip[j].Y-strip[i].Y) < d; j++ {
				d = math.Min(d, dist(strip[i], strip[j]))
			}
		}

		return d
	}

	return closestRec(0, len(pts)-1)
}

func dist(p1, p2 Point) float64 {
	return math.Sqrt((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))
}
