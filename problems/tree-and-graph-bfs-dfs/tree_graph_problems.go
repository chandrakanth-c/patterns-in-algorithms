package tree_graph

// =============================================================
// TreeNode — binary tree node.
// =============================================================

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// =============================================================
// GraphNode — undirected graph node (adjacency list).
// =============================================================

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

// =============================================================
// LevelOrder — LeetCode #102
// BFS using a queue; collect node values level by level.
// Time: O(n), Space: O(n).
// =============================================================

// LevelOrder returns the level-order traversal of a binary tree.
func LevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[i]
			level[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		result = append(result, level)
	}
	return result
}

// =============================================================
// LowestCommonAncestor — LeetCode #236
// DFS post-order: the first node that sees both p and q in its
// subtrees (or is p/q itself) is the LCA.
// Time: O(n), Space: O(h).
// =============================================================

// LowestCommonAncestor finds the LCA of nodes p and q in the tree.
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

// =============================================================
// LadderLength — LeetCode #127
// BFS from beginWord; at each step try all single-character
// mutations and advance to unused words in wordList.
// Time: O(M^2 · N) where M=word length, N=list size.
// =============================================================

// LadderLength returns the length of the shortest transformation sequence.
func LadderLength(beginWord, endWord string, wordList []string) int {
	wordSet := make(map[string]bool, len(wordList))
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return 0
	}

	queue := []string{beginWord}
	visited := map[string]bool{beginWord: true}
	steps := 1

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			word := queue[i]
			bs := []byte(word)
			for j := 0; j < len(bs); j++ {
				orig := bs[j]
				for c := byte('a'); c <= byte('z'); c++ {
					if c == orig {
						continue
					}
					bs[j] = c
					next := string(bs)
					if next == endWord {
						return steps + 1
					}
					if wordSet[next] && !visited[next] {
						visited[next] = true
						queue = append(queue, next)
					}
					bs[j] = orig
				}
			}
		}
		queue = queue[size:]
		steps++
	}
	return 0
}

// =============================================================
// CloneGraph — LeetCode #133
// DFS with a visited map from original node → cloned node.
// Time: O(V+E), Space: O(V).
// =============================================================

// CloneGraph deep-copies an undirected connected graph.
func CloneGraph(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}
	visited := map[*GraphNode]*GraphNode{}
	var dfs func(n *GraphNode) *GraphNode
	dfs = func(n *GraphNode) *GraphNode {
		if clone, ok := visited[n]; ok {
			return clone
		}
		clone := &GraphNode{Val: n.Val}
		visited[n] = clone
		for _, nb := range n.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(nb))
		}
		return clone
	}
	return dfs(node)
}
