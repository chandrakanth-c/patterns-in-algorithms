# 21 - Trie (Prefix Tree)

## 21.1 - Overview & Theoretical Foundations (CLRS Chapter 31 & Radix Trees)

* A **Trie** (derived from re**trie**val), or **Prefix Tree**, is an $m$-ary ordered search tree where keys are strings over an alphabet $\Sigma$.
* Unlike binary search trees, nodes do not store the explicit search key. Instead, the position of a node within the tree defines its corresponding string prefix.
* All descendants of a node share a common prefix, and the root is associated with the empty string $\epsilon$.
* **Lookup Invariance:** The search time for a key of length $L$ depends strictly on $L$ ($\mathcal{O}(L)$ time), completely independent of the total number of keys $N$ stored in the dataset.

---

## 21.2 - Properties of a problem that suggests Trie

* **Prefix matching**, autocomplete suggestions, dictionary spell-checking, IP routing (longest prefix match).
* Bitwise maximum XOR pairs (storing binary sequences of fixed bit width in a Binary Trie).

---

## 21.3 - Classic Example: Implement Trie (Prefix Tree)

### Java Implementation

```java
public class Trie {

    private static class TrieNode {
        TrieNode[] children = new TrieNode[26]; // for 'a' through 'z'
        boolean isEndOfWord = false;
    }

    private final TrieNode root;

    public Trie() {
        root = new TrieNode();
    }

    // Inserts a word into the trie in O(L)
    public void insert(String word) {
        TrieNode current = root;
        for (char c : word.toCharArray()) {
            int idx = c - 'a';
            if (current.children[idx] == null) {
                current.children[idx] = new TrieNode();
            }
            current = current.children[idx];
        }
        current.isEndOfWord = true;
    }

    // Returns true if the word is in the trie in O(L)
    public boolean search(String word) {
        TrieNode node = findPrefixNode(word);
        return node != null && node.isEndOfWord;
    }

    // Returns true if there is any word starting with given prefix in O(L)
    public boolean startsWith(String prefix) {
        return findPrefixNode(prefix) != null;
    }

    private TrieNode findPrefixNode(String prefix) {
        TrieNode current = root;
        for (char c : prefix.toCharArray()) {
            int idx = c - 'a';
            if (current.children[idx] == null) return null;
            current = current.children[idx];
        }
        return current;
    }
}
```

---

### Go Implementation

```go
package main

type TrieNode struct {
	children    [26]*TrieNode
	isEndOfWord bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

// Insert inserts a word into the trie
func (t *Trie) Insert(word string) {
	current := t.root
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if current.children[idx] == nil {
			current.children[idx] = &TrieNode{}
		}
		current = current.children[idx]
	}
	current.isEndOfWord = true
}

// Search returns true if the word is present
func (t *Trie) Search(word string) bool {
	node := t.findPrefixNode(word)
	return node != nil && node.isEndOfWord
}

// StartsWith returns true if there is any word starting with the prefix
func (t *Trie) StartsWith(prefix string) bool {
	return t.findPrefixNode(prefix) != nil
}

func (t *Trie) findPrefixNode(str string) *TrieNode {
	current := t.root
	for i := 0; i < len(str); i++ {
		idx := str[i] - 'a'
		if current.children[idx] == nil {
			return nil
		}
		current = current.children[idx]
	}
	return current
}
```

---

## 21.4 - Time & Space Complexity Analysis

* **Time Complexity:** $\mathcal{O}(L)$ for `insert`, `search`, and `startsWith`, where $L$ is the length of the string.
* **Space Complexity:** $\mathcal{O}(N \times L \times |\Sigma|)$ worst case where $|\Sigma| = 26$ is alphabet size.

---

## 21.5 - Classic LeetCode & Benchmark Problems

* **Implement Trie (Prefix Tree)** (LeetCode #208)
* **Design Add and Search Words Data Structure** (LeetCode #211)
* **Word Search II** (LeetCode #212)
* **Maximum XOR of Two Numbers in an Array** (LeetCode #421)
* **Replace Words** (LeetCode #648)

---

## 21.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 31: String Matching & Radix tree exercises (pp. 985–1012)
* https://en.wikipedia.org/wiki/Trie
* https://leetcode.com/problems/implement-trie-prefix-tree/
