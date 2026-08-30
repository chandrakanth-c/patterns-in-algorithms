package dp

import (
	"math"
	"testing"
)

func TestCutRod(t *testing.T) {
	prices := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	tests := []struct {
		n        int
		expected int
	}{
		{n: 1, expected: 1},
		{n: 4, expected: 10},
		{n: 8, expected: 22},
	}

	for _, tt := range tests {
		got := CutRod(prices, tt.n)
		if got != tt.expected {
			t.Errorf("CutRod(prices, %d) = %d; want %d", tt.n, got, tt.expected)
		}
	}
}

func TestMatrixChainOrder(t *testing.T) {
	p := []int{30, 35, 15, 5, 10, 20, 25} // Matrices A1(30x35), A2(35x15), A3(15x5), A4(5x10), A5(10x20), A6(20x25)
	expected := 15125
	got := MatrixChainOrder(p)
	if got != expected {
		t.Errorf("MatrixChainOrder(%v) = %d; want %d", p, got, expected)
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		text1, text2 string
		expected     int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}

	for _, tt := range tests {
		got := LongestCommonSubsequence(tt.text1, tt.text2)
		if got != tt.expected {
			t.Errorf("LongestCommonSubsequence(%q, %q) = %d; want %d", tt.text1, tt.text2, got, tt.expected)
		}
	}
}

func TestOptimalBST(t *testing.T) {
	p := []float64{0.15, 0.10, 0.05, 0.10, 0.20}
	q := []float64{0.05, 0.10, 0.05, 0.05, 0.05, 0.10}
	n := 5
	expected := 2.75
	got := OptimalBST(p, q, n)
	if math.Abs(got-expected) > 1e-4 {
		t.Errorf("OptimalBST() = %f; want %f", got, expected)
	}
}

func TestClimbStairs(t *testing.T) {
	if got := ClimbStairs(2); got != 2 {
		t.Errorf("ClimbStairs(2) = %d; want 2", got)
	}
	if got := ClimbStairs(3); got != 3 {
		t.Errorf("ClimbStairs(3) = %d; want 3", got)
	}
	if got := ClimbStairs(5); got != 8 {
		t.Errorf("ClimbStairs(5) = %d; want 8", got)
	}
}

func TestCoinChange(t *testing.T) {
	if got := CoinChange([]int{1, 2, 5}, 11); got != 3 {
		t.Errorf("CoinChange([1,2,5], 11) = %d; want 3", got)
	}
	if got := CoinChange([]int{2}, 3); got != -1 {
		t.Errorf("CoinChange([2], 3) = %d; want -1", got)
	}
	if got := CoinChange([]int{1}, 0); got != 0 {
		t.Errorf("CoinChange([1], 0) = %d; want 0", got)
	}
}

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7}, 1},
	}

	for _, tt := range tests {
		got := LengthOfLIS(tt.nums)
		if got != tt.expected {
			t.Errorf("LengthOfLIS(%v) = %d; want %d", tt.nums, got, tt.expected)
		}
	}
}

func TestMinDistance(t *testing.T) {
	tests := []struct {
		w1, w2   string
		expected int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	}

	for _, tt := range tests {
		got := MinDistance(tt.w1, tt.w2)
		if got != tt.expected {
			t.Errorf("MinDistance(%q, %q) = %d; want %d", tt.w1, tt.w2, got, tt.expected)
		}
	}
}
