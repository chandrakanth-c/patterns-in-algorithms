# 16 - Tree & Graph Traversals (BFS & DFS)

## 16.1 - Overview & Theoretical Foundations (CLRS Chapter 20)

* **Breadth-First Search (BFS, CLRS 20.2):**
  * Explores the graph outward in concentric frontier waves from a source vertex $s$ using a **FIFO Queue**.
  * **Shortest Path Property (CLRS Theorem 20.5):** For unweighted graphs, BFS is guaranteed to discover the shortest path (minimum number of edges) from source $s$ to all reachable vertices.
* **Depth-First Search (DFS, CLRS 20.3):**
  * Searches deeper in the graph whenever possible using a **LIFO Stack** or system call stack.
  * **Parenthesis Theorem (CLRS Theorem 20.7):** In any DFS of graph $G$, vertex $v$ is a descendant of vertex $u$ in the DFS forest if and only if $[v.d, v.f] \subset [u.d, u.f]$ (where $d$ is discovery time and $f$ is finish time).
  * **Edge Classification:** Tree edges, Back edges (indicate cycles in directed graphs), Forward edges, and Cross edges.

---

## 16.2 - Properties of a problem that suggests BFS / DFS

* **BFS:** Shortest path on unweighted graphs, minimum transformations, level-order metric computations.
* **DFS:** Cycle detection, connectivity, reachability, tree preorder/inorder/postorder properties, topological ordering.

---

## 16.3 - Classic Examples: Tree Level-Order (BFS) & Graph Connected Components (DFS)

### Java Implementation

```java
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.List;
import java.util.Queue;

public class TreeGraphTraversals {

    public static class TreeNode {
        public int val;
        public TreeNode left, right;
        public TreeNode(int val) { this.val = val; }
    }

    // 1. BFS Level-Order Traversal
    public static List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> result = new ArrayList<>();
        if (root == null) return result;

        Queue<TreeNode> queue = new ArrayDeque<>();
        queue.offer(root);

        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            List<Integer> currentLevel = new ArrayList<>(levelSize);

            for (int i = 0; i < levelSize; i++) {
                TreeNode node = queue.poll();
                currentLevel.add(node.val);

                if (node.left != null) queue.offer(node.left);
                if (node.right != null) queue.offer(node.right);
            }
            result.add(currentLevel);
        }
        return result;
    }

    // 2. DFS Recursive Traversal on Graph (Adjacency List)
    public static void dfs(int u, List<List<Integer>> adj, boolean[] visited, List<Integer> traversal) {
        visited[u] = true;
        traversal.add(u);

        for (int v : adj.get(u)) {
            if (!visited[v]) {
                dfs(v, adj, visited, traversal);
            }
        }
    }
}
```

---

### Go Implementation

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LevelOrder performs BFS level-order traversal on a binary tree
func LevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:] // Dequeue

			currentLevel = append(currentLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

// DFS performs Depth-First Search on an adjacency list graph
func DFS(u int, adj [][]int, visited []bool, traversal *[]int) {
	visited[u] = true
	*traversal = append(*traversal, u)

	for _, v := range adj[u] {
		if !visited[v] {
			DFS(v, adj, visited, traversal)
		}
	}
}
```

---

## 16.4 - Time & Space Complexity Analysis

* **Time Complexity:** $\Theta(V + E)$ where $V$ is number of vertices and $E$ is number of edges (for trees, $\Theta(N)$). Every vertex is enqueued/visited once and every adjacency edge inspected.
* **Space Complexity:** $\mathcal{O}(V)$ for the queue (BFS) or recursion call stack (DFS).

---

## 16.5 - Classic LeetCode & CLRS Benchmarks

* **Breadth-First Search & Depth-First Search** (CLRS 20.2, 20.3)
* **Binary Tree Level Order Traversal** (LeetCode #102)
* **Lowest Common Ancestor of a Binary Tree** (LeetCode #236)
* **Word Ladder** (LeetCode #127)
* **Clone Graph** (LeetCode #133)
* **Course Schedule** (LeetCode #207)

---

## 16.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 20: Elementary Graph Algorithms (pp. 589–623)
  * Section 20.2: Breadth-first search (pp. 594–603)
  * Section 20.3: Depth-first search (pp. 603–613)
* https://leetcode.com/explore/interview/card/cheatsheets/720/resources/4723/
* https://techinterviewhandbook.org/algorithms/graph/
