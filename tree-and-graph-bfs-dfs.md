# 16 - Tree & Graph Traversals (BFS & DFS)

## 16.1 - Overview

* **Breadth-First Search (BFS):** Explores neighbors level-by-level using a **FIFO Queue**. It is optimal for finding the **shortest path in unweighted graphs** and level-order metrics in trees.
* **Depth-First Search (DFS):** Explores as deep as possible down each branch before backtracking, implemented using recursion or an explicit **LIFO Stack**. It is ideal for tree traversals (preorder, inorder, postorder), connectivity, and cycle detection.

---

## 16.2 - Properties of a problem that suggests BFS / DFS

* Traversing hierarchical structures (Binary Trees, N-ary Trees) or networks (Graphs).
* **BFS:** Shortest transformation sequence (e.g. Word Ladder), minimum steps to reach target, level-by-level processing.
* **DFS:** Path sum problems, finding all paths from source to target, validating Binary Search Trees (BST), finding Connected Components.

---

## 16.3 - Classic Examples

### Java Implementation: Binary Tree Level Order Traversal (BFS)

```java
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.List;
import java.util.Queue;

public class TreeGraphTraversals {

    static class TreeNode {
        int val;
        TreeNode left, right;
        TreeNode(int val) { this.val = val; }
    }

    public static List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> result = new ArrayList<>();
        if (root == null) return result;

        Queue<TreeNode> queue = new ArrayDeque<>();
        queue.offer(root);

        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            List<Integer> currentLevel = new ArrayList<>(levelSize);

            for (int i = 0; i < levelSize; i++) {
                TreeNode currentNode = queue.poll();
                currentLevel.add(currentNode.val);

                if (currentNode.left != null) queue.offer(currentNode.left);
                if (currentNode.right != null) queue.offer(currentNode.right);
            }

            result.add(currentLevel);
        }

        return result;
    }
}
```

---

## 16.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(V + E)$ where $V$ is the number of vertices and $E$ is the number of edges (for trees, $\mathcal{O}(N)$).
* **Space Complexity:**
  * BFS: $\mathcal{O}(W)$ where $W$ is the maximum width/breadth of the tree/graph.
  * DFS: $\mathcal{O}(H)$ where $H$ is the maximum depth/height of the tree/graph.

---

## 16.5 - Classic LeetCode Problems

* **Binary Tree Level Order Traversal** (LeetCode #102)
* **Binary Tree Zigzag Level Order Traversal** (LeetCode #103)
* **Maximum Depth of Binary Tree** (LeetCode #104)
* **Path Sum I, II & III** (LeetCode #112, #113, #437)
* **Lowest Common Ancestor of a Binary Tree** (LeetCode #236)
* **Word Ladder** (LeetCode #127)
* **Clone Graph** (LeetCode #133)

---

## 16.6 - Sources used for this file:
https://leetcode.com/explore/interview/card/cheatsheets/720/resources/4723/ <br>
https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d38a08d2bb2d978e1cd7 <br>
https://www.geeksforgeeks.org/breadth-first-search-or-bfs-for-a-graph/ <br>
https://www.geeksforgeeks.org/depth-first-search-or-dfs-for-a-graph/ <br>
https://techinterviewhandbook.org/algorithms/graph/
