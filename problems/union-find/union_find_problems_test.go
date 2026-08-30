package union_find

import (
	"sort"
	"testing"
)

func TestKruskalMST(t *testing.T) {
	// Simple triangle: 3 nodes, edges with weights
	edges := [][3]int{{0, 1, 1}, {1, 2, 2}, {0, 2, 3}}
	got := KruskalMST(3, edges)
	if got != 3 { // MST picks edges weight 1 + 2
		t.Errorf("KruskalMST: got %d, want 3", got)
	}
}

func TestCountComponents(t *testing.T) {
	got := CountComponents(5, [][2]int{{0, 1}, {1, 2}, {3, 4}})
	if got != 2 {
		t.Errorf("CountComponents: got %d, want 2", got)
	}
	got2 := CountComponents(5, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}})
	if got2 != 1 {
		t.Errorf("CountComponents: got %d, want 1", got2)
	}
}

func TestFindRedundantConnection(t *testing.T) {
	edges := [][2]int{{1, 2}, {1, 3}, {2, 3}}
	got := FindRedundantConnection(edges)
	if got != [2]int{2, 3} {
		t.Errorf("FindRedundantConnection: got %v, want [2 3]", got)
	}
}

func TestAccountsMerge(t *testing.T) {
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"},
	}
	result := AccountsMerge(accounts)
	// Should produce 3 merged accounts
	if len(result) != 3 {
		t.Errorf("AccountsMerge: expected 3 merged accounts, got %d", len(result))
	}
	// Find the John account with 3 emails
	found := false
	for _, acc := range result {
		if acc[0] == "John" && len(acc) == 4 { // name + 3 emails
			emails := acc[1:]
			sort.Strings(emails)
			expected := []string{"john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"}
			match := true
			for i := range expected {
				if emails[i] != expected[i] {
					match = false
				}
			}
			if match {
				found = true
			}
		}
	}
	if !found {
		t.Error("AccountsMerge: merged John account not found")
	}
}
