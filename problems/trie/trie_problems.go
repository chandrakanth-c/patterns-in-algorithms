package trie

// --- Trie (LeetCode #208) ---

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{}}
}

// Insert inserts a word into the trie.
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

// Search returns true if word is in the trie.
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return node.isEnd
}

// StartsWith returns true if any word in the trie starts with prefix.
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return true
}

// --- WordDictionary with '.' wildcard support (LeetCode #211) ---

type WordDictNode struct {
	children [26]*WordDictNode
	isEnd    bool
}

type WordDictionary struct {
	root *WordDictNode
}

func NewWordDictionary() *WordDictionary {
	return &WordDictionary{root: &WordDictNode{}}
}

func (d *WordDictionary) AddWord(word string) {
	node := d.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &WordDictNode{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

// Search supports '.' as a wildcard matching any single character.
func (d *WordDictionary) Search(word string) bool {
	return d.searchHelper(d.root, word)
}

func (d *WordDictionary) searchHelper(node *WordDictNode, word string) bool {
	if len(word) == 0 {
		return node.isEnd
	}
	ch := word[0]
	rest := word[1:]
	if ch == '.' {
		for _, child := range node.children {
			if child != nil && d.searchHelper(child, rest) {
				return true
			}
		}
		return false
	}
	idx := ch - 'a'
	if node.children[idx] == nil {
		return false
	}
	return d.searchHelper(node.children[idx], rest)
}

// --- Word Search II (LeetCode #212) — Trie + DFS ---

// FindWords finds all words from the list that can be formed on the board.
func FindWords(board [][]byte, words []string) []string {
	// Build a trie of all words
	root := &TrieNode{}
	// We reuse TrieNode but need a word field; use a parallel map approach.
	// Instead, define a local node with word field.
	type WNode struct {
		children [26]*WNode
		word     string // non-empty at terminal
	}
	wRoot := &WNode{}
	for _, w := range words {
		node := wRoot
		for _, ch := range w {
			idx := ch - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &WNode{}
			}
			node = node.children[idx]
		}
		node.word = w
	}

	rows, cols := len(board), len(board[0])
	found := map[string]bool{}

	var dfs func(node *WNode, r, c int)
	dfs = func(node *WNode, r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] == '#' {
			return
		}
		ch := board[r][c]
		idx := ch - 'a'
		child := node.children[idx]
		if child == nil {
			return
		}
		if child.word != "" {
			found[child.word] = true
		}
		board[r][c] = '#' // mark visited
		dfs(child, r+1, c)
		dfs(child, r-1, c)
		dfs(child, r, c+1)
		dfs(child, r, c-1)
		board[r][c] = ch // restore
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(wRoot, r, c)
		}
	}

	result := make([]string, 0, len(found))
	for w := range found {
		result = append(result, w)
	}
	return result
}

// --- Maximum XOR of Two Numbers using Trie (LeetCode #421) ---

type XorNode struct {
	children [2]*XorNode
}

// FindMaximumXOR finds the maximum XOR of any two numbers in nums.
func FindMaximumXOR(nums []int) int {
	root := &XorNode{}
	// Insert all numbers bit by bit (from MSB to LSB, 31 bits)
	for _, n := range nums {
		node := root
		for i := 31; i >= 0; i-- {
			bit := (n >> i) & 1
			if node.children[bit] == nil {
				node.children[bit] = &XorNode{}
			}
			node = node.children[bit]
		}
	}
	maxXOR := 0
	for _, n := range nums {
		node := root
		curXOR := 0
		for i := 31; i >= 0; i-- {
			bit := (n >> i) & 1
			want := 1 - bit // prefer opposite bit for max XOR
			if node.children[want] != nil {
				curXOR |= (1 << i)
				node = node.children[want]
			} else {
				node = node.children[bit]
			}
		}
		if curXOR > maxXOR {
			maxXOR = curXOR
		}
	}
	return maxXOR
}
