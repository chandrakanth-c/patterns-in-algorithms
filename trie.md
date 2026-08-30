# 21 - Trie (Prefix Tree)

## 21.1 - Overview

* A **Trie** (derived from "re**trie**val"), or **Prefix Tree**, is a tree-like search data structure used to store and retrieve keys in a dataset of strings.
* Unlike binary search trees, no node in the tree stores the key associated with that node; instead, its position in the tree defines what key it is associated with.
* All descendants of a node share a common string prefix.

---

## 21.2 - Properties of a problem that suggests Trie

* **Prefix-based operations:** Auto-complete suggestions, spell checking, IP routing (longest prefix match).
* Fast dictionary word lookups ($\mathcal{O}(L)$ time where $L$ is word length, independent of the number of words $N$).
* **Binary Trie:** Finding maximum XOR pair in an array of integers by storing 32-bit binary representations.

---

## 21.3 - Classic Example: Implement Trie (Prefix Tree)

### Java Implementation

```java
public class Trie {

    static class TrieNode {
        TrieNode[] children;
        boolean isEndOfWord;

        public TrieNode() {
            children = new TrieNode[26]; // for 'a' through 'z'
            isEndOfWord = false;
        }
    }

    private final TrieNode root;

    public Trie() {
        root = new TrieNode();
    }

    // Inserts a word into the trie
    public void insert(String word) {
        TrieNode current = root;
        for (char c : word.toCharArray()) {
            int index = c - 'a';
            if (current.children[index] == null) {
                current.children[index] = new TrieNode();
            }
            current = current.children[index];
        }
        current.isEndOfWord = true;
    }

    // Returns true if the word is in the trie
    public boolean search(String word) {
        TrieNode node = findPrefixNode(word);
        return node != null && node.isEndOfWord;
    }

    // Returns true if there is any word in the trie that starts with the given prefix
    public boolean startsWith(String prefix) {
        return findPrefixNode(prefix) != null;
    }

    private TrieNode findPrefixNode(String str) {
        TrieNode current = root;
        for (char c : str.toCharArray()) {
            int index = c - 'a';
            if (current.children[index] == null) return null;
            current = current.children[index];
        }
        return current;
    }
}
```

---

## 21.4 - Time & Space Complexity

* **Time Complexity:**
  * `insert(word)`: $\mathcal{O}(L)$ where $L$ is word length.
  * `search(word)`: $\mathcal{O}(L)$
  * `startsWith(prefix)`: $\mathcal{O}(L)$
* **Space Complexity:** $\mathcal{O}(\text{Total Characters} \times \Sigma)$ where $\Sigma$ is alphabet size (26 for lowercase English).

---

## 21.5 - Classic LeetCode Problems

* **Implement Trie (Prefix Tree)** (LeetCode #208)
* **Design Add and Search Words Data Structure** (LeetCode #211)
* **Word Search II** (LeetCode #212)
* **Maximum XOR of Two Numbers in an Array** (Binary Trie) (LeetCode #421)
* **Replace Words** (LeetCode #648)

---

## 21.6 - Sources used for this file:
https://en.wikipedia.org/wiki/Trie <br>
https://leetcode.com/problems/implement-trie-prefix-tree/ <br>
https://www.geeksforgeeks.org/trie-insert-and-search/ <br>
https://techinterviewhandbook.org/algorithms/trie/
