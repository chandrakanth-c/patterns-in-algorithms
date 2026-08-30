package trie

import (
	"sort"
	"testing"
)

func TestTrie(t *testing.T) {
	tr := NewTrie()
	tr.Insert("apple")
	if !tr.Search("apple") {
		t.Error("Trie.Search: expected true for 'apple'")
	}
	if tr.Search("app") {
		t.Error("Trie.Search: expected false for 'app' (not inserted)")
	}
	if !tr.StartsWith("app") {
		t.Error("Trie.StartsWith: expected true for 'app'")
	}
	tr.Insert("app")
	if !tr.Search("app") {
		t.Error("Trie.Search: expected true for 'app' after insert")
	}
}

func TestWordDictionary(t *testing.T) {
	d := NewWordDictionary()
	d.AddWord("bad")
	d.AddWord("dad")
	d.AddWord("mad")
	if d.Search("pad") {
		t.Error("WordDictionary.Search: expected false for 'pad'")
	}
	if !d.Search("bad") {
		t.Error("WordDictionary.Search: expected true for 'bad'")
	}
	if !d.Search(".ad") {
		t.Error("WordDictionary.Search: expected true for '.ad'")
	}
	if !d.Search("b..") {
		t.Error("WordDictionary.Search: expected true for 'b..'")
	}
}

func TestFindWords(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	got := FindWords(board, words)
	sort.Strings(got)
	expected := []string{"eat", "oath"}
	if len(got) != len(expected) {
		t.Errorf("FindWords: got %v, want %v", got, expected)
		return
	}
	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("FindWords: got %v, want %v", got, expected)
		}
	}
}

func TestFindMaximumXOR(t *testing.T) {
	got := FindMaximumXOR([]int{3, 10, 5, 25, 2, 8})
	if got != 28 {
		t.Errorf("FindMaximumXOR: got %d, want 28", got)
	}
	got2 := FindMaximumXOR([]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70})
	if got2 != 127 {
		t.Errorf("FindMaximumXOR: got %d, want 127", got2)
	}
}
