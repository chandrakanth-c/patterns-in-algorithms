package bit_manipulation

import (
	"reflect"
	"sort"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	got := SingleNumber([]int{2, 2, 1})
	if got != 1 {
		t.Errorf("SingleNumber: got %d, want 1", got)
	}
	got = SingleNumber([]int{4, 1, 2, 1, 2})
	if got != 4 {
		t.Errorf("SingleNumber: got %d, want 4", got)
	}
}

func TestSingleNumberII(t *testing.T) {
	got := SingleNumberII([]int{2, 2, 3, 2})
	if got != 3 {
		t.Errorf("SingleNumberII: got %d, want 3", got)
	}
	got = SingleNumberII([]int{0, 1, 0, 1, 0, 1, 99})
	if got != 99 {
		t.Errorf("SingleNumberII: got %d, want 99", got)
	}
}

func TestSingleNumberIII(t *testing.T) {
	got := SingleNumberIII([]int{1, 2, 1, 3, 2, 5})
	expected := [2]int{3, 5}
	if got != expected {
		t.Errorf("SingleNumberIII: got %v, want %v", got, expected)
	}
}

func TestHammingWeight(t *testing.T) {
	got := HammingWeight(11) // 1011 -> 3 bits
	if got != 3 {
		t.Errorf("HammingWeight(11): got %d, want 3", got)
	}
	got = HammingWeight(128) // 10000000 -> 1 bit
	if got != 1 {
		t.Errorf("HammingWeight(128): got %d, want 1", got)
	}
}

func TestCountBits(t *testing.T) {
	got := CountBits(5)
	expected := []int{0, 1, 1, 2, 1, 2}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("CountBits(5): got %v, want %v", got, expected)
	}
}

func TestSubsetsBitmask(t *testing.T) {
	got := SubsetsBitmask([]int{1, 2, 3})
	if len(got) != 8 {
		t.Errorf("SubsetsBitmask: expected 8 subsets, got %d", len(got))
	}
	// Verify the empty subset is present
	hasEmpty := false
	for _, s := range got {
		if len(s) == 0 {
			hasEmpty = true
		}
	}
	if !hasEmpty {
		t.Error("SubsetsBitmask: missing empty subset")
	}
	// Verify the full set is present
	hasFull := false
	for _, s := range got {
		tmp := make([]int, len(s))
		copy(tmp, s)
		sort.Ints(tmp)
		if reflect.DeepEqual(tmp, []int{1, 2, 3}) {
			hasFull = true
		}
	}
	if !hasFull {
		t.Error("SubsetsBitmask: missing full subset {1,2,3}")
	}
}
