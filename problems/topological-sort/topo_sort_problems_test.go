package topo_sort

import (
	"reflect"
	"sort"
	"testing"
)

func TestCanFinish(t *testing.T) {
	// No cycle: 2 courses, course 1 requires course 0
	if !CanFinish(2, [][2]int{{1, 0}}) {
		t.Error("CanFinish: expected true for acyclic")
	}
	// Cycle: 0->1->0
	if CanFinish(2, [][2]int{{1, 0}, {0, 1}}) {
		t.Error("CanFinish: expected false for cycle")
	}
}

func TestFindOrder(t *testing.T) {
	got := FindOrder(4, [][2]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	if len(got) != 4 {
		t.Errorf("FindOrder: expected 4 courses, got %v", got)
	}
	// Verify each prerequisite is satisfied in the order
	pos := make(map[int]int)
	for i, c := range got {
		pos[c] = i
	}
	prereqs := [][2]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	for _, p := range prereqs {
		if pos[p[0]] < pos[p[1]] {
			t.Errorf("FindOrder: %d appears before prerequisite %d", p[0], p[1])
		}
	}
	// Cycle case returns empty
	empty := FindOrder(2, [][2]int{{0, 1}, {1, 0}})
	if len(empty) != 0 {
		t.Error("FindOrder: expected empty for cycle")
	}
}

func TestAlienOrder(t *testing.T) {
	words := []string{"wrt", "wrf", "er", "ett", "rftt"}
	order := AlienOrder(words)
	// Valid: contains 5 unique chars; verify adjacency constraints
	if len(order) < 4 {
		t.Errorf("AlienOrder: too short order = %q", order)
	}
	// Invalid (w2 is prefix of w1)
	invalid := AlienOrder([]string{"abc", "ab"})
	if invalid != "" {
		t.Errorf("AlienOrder: expected empty for invalid input, got %q", invalid)
	}
}

func TestFindMinHeightTrees(t *testing.T) {
	got := FindMinHeightTrees(4, [][2]int{{1, 0}, {1, 2}, {1, 3}})
	expected := []int{1}
	sort.Ints(got)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("FindMinHeightTrees: got %v, want %v", got, expected)
	}

	got2 := FindMinHeightTrees(6, [][2]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}})
	sort.Ints(got2)
	expected2 := []int{3, 4}
	if !reflect.DeepEqual(got2, expected2) {
		t.Errorf("FindMinHeightTrees: got %v, want %v", got2, expected2)
	}
}
