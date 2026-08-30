package dc

import (
	"math"
	"reflect"
	"testing"
)

func TestMaxSubarray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6
	got := MaxSubarray(nums)
	if got != expected {
		t.Errorf("MaxSubarray(%v) = %d; want %d", nums, got, expected)
	}
}

func TestStrassenMultiply2x2(t *testing.T) {
	A := [][]int{{1, 2}, {3, 4}}
	B := [][]int{{5, 6}, {7, 8}}
	expected := [][]int{{19, 22}, {43, 50}}
	got := StrassenMultiply2x2(A, B)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("StrassenMultiply2x2() = %v; want %v", got, expected)
	}
}

func TestQuickSort(t *testing.T) {
	nums := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	expected := []int{2, 3, 5, 6, 7, 9, 10, 11, 12, 14}
	QuickSort(nums)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("QuickSort() = %v; want %v", nums, expected)
	}
}

func TestRandomizedSelect(t *testing.T) {
	nums := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	// Sorted: 2, 3, 5, 6, 7, 9, 10, 11, 12, 14
	kthSmallest := RandomizedSelect(nums, 0, len(nums)-1, 5) // 5th smallest is 7
	if kthSmallest != 7 {
		t.Errorf("RandomizedSelect(k=5) = %d; want 7", kthSmallest)
	}
}

func TestClosestPair(t *testing.T) {
	points := []Point{
		{X: 2, Y: 3},
		{X: 12, Y: 30},
		{X: 40, Y: 50},
		{X: 5, Y: 1},
		{X: 12, Y: 10},
		{X: 3, Y: 4},
	}
	expected := math.Sqrt(2) // between (2,3) and (3,4) is ~1.4142
	got := ClosestPair(points)
	if math.Abs(got-expected) > 1e-4 {
		t.Errorf("ClosestPair() = %f; want %f", got, expected)
	}
}
