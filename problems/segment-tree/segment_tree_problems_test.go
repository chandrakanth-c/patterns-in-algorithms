package segment_tree

import (
	"reflect"
	"testing"
)

func TestNumArray(t *testing.T) {
	na := NewNumArray([]int{1, 3, 5})
	if got := na.SumRange(0, 2); got != 9 {
		t.Errorf("SumRange(0,2): got %d, want 9", got)
	}
	na.Update(1, 2) // change index 1 from 3 to 2 -> [1,2,5]
	if got := na.SumRange(0, 2); got != 8 {
		t.Errorf("SumRange(0,2) after update: got %d, want 8", got)
	}
	if got := na.SumRange(1, 2); got != 7 {
		t.Errorf("SumRange(1,2) after update: got %d, want 7", got)
	}
}

func TestCountSmaller(t *testing.T) {
	got := CountSmaller([]int{5, 2, 6, 1})
	expected := []int{2, 1, 1, 0}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("CountSmaller: got %v, want %v", got, expected)
	}
}

func TestGetSkyline(t *testing.T) {
	buildings := [][3]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	got := GetSkyline(buildings)
	expected := [][2]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("GetSkyline: got %v, want %v", got, expected)
	}
}
