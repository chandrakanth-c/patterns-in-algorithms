# 17 - Matrix Traversal (Grid BFS & DFS / Flood Fill)

## 17.1 - Overview & Theoretical Foundations (CLRS Chapter 20)

* A 2D matrix of size $M \times N$ represents an **implicit planar graph**:
  * Total vertices: $V = M \times N$.
  * Total edges: $E \le 4MN$ (under 4-directional connectivity) or $8MN$ (under 8-directional connectivity).
* **Multi-Source BFS:**
  * When finding minimum distance from *multiple simultaneous start points* (e.g. fire spreading, rotting oranges), initializing the BFS queue with all source cells at $t=0$ explores distance frontiers in $\mathcal{O}(M \times N)$ total time.
* **Connected Component Counting (Flood Fill):**
  * Sinking visited cells in-place (mutating `'1'` to `'0'`) allows component isolation without auxiliary `visited[][]` tables.

---

## 17.2 - Properties of a problem that suggests Matrix Traversal

* Grid/maze navigation, island counting, terrain simulation, shortest path on unweighted 2D boards.

---

## 17.3 - Classic Example: Number of Islands (DFS Flood Fill)

### Java Implementation

```java
public class MatrixTraversal {

    public static int numIslands(char[][] grid) {
        if (grid == null || grid.length == 0) return 0;

        int rows = grid.length, cols = grid[0].length;
        int islandCount = 0;

        for (int r = 0; r < rows; r++) {
            for (int c = 0; c < cols; c++) {
                if (grid[r][c] == '1') {
                    islandCount++;
                    dfsSink(grid, r, c); // Sink the connected island
                }
            }
        }

        return islandCount;
    }

    private static void dfsSink(char[][] grid, int r, int c) {
        if (r < 0 || r >= grid.length || c < 0 || c >= grid[0].length || grid[r][c] != '1') {
            return;
        }

        grid[r][c] = '0'; // Mark visited

        dfsSink(grid, r + 1, c);
        dfsSink(grid, r - 1, c);
        dfsSink(grid, r, c + 1);
        dfsSink(grid, r, c - 1);
    }
}
```

---

### Go Implementation

```go
package main

// NumIslands counts the number of connected 1s in a grid
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	var dfsSink func(r, c int)
	dfsSink = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != '1' {
			return
		}

		grid[r][c] = '0' // Sink visited land

		dfsSink(r+1, c)
		dfsSink(r-1, c)
		dfsSink(r, c+1)
		dfsSink(r, c-1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfsSink(r, c)
			}
		}
	}

	return count
}
```

---

## 17.4 - Time & Space Complexity Analysis

* **Time Complexity:** $\Theta(M \times N)$ — Each cell in the matrix is processed at most a constant number of times (once per incoming neighbor direction).
* **Space Complexity:** $\mathcal{O}(M \times N)$ in the worst case for recursion depth or BFS queue (e.g. grid completely filled with land).

---

## 17.5 - Classic LeetCode & Benchmark Problems

### 1. Number of Islands (LeetCode #200)
* **Problem Statement**: Given an `m x n` 2D binary grid `grid` which represents a map of '1's (land) and '0's (water), return the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
* **Solution Link**: [problems/matrix-traversal/matrix_traversal_problems.go](problems/matrix-traversal/matrix_traversal_problems.go) (`NumIslands`)
* **Explanation**: Iterates through each cell in the grid. When a '1' is encountered, it increments the island count and performs a DFS to "sink" (mark as '0') all connected land cells.
* **Conceptual Link**: Uses the **Flood Fill** technique to isolate and count connected components in a planar graph.

### 2. Rotting Oranges (LeetCode #994)
* **Problem Statement**: You are given an `m x n` grid where each cell can have one of three values: 0 (empty), 1 (fresh orange), or 2 (rotten orange). Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten. Return the minimum number of minutes that must elapse until no cell has a fresh orange.
* **Solution Link**: [problems/matrix-traversal/matrix_traversal_problems.go](problems/matrix-traversal/matrix_traversal_problems.go) (`OrangesRotting`)
* **Explanation**: Employs **Multi-Source BFS**. All initially rotten oranges are added to a queue. The BFS explores the grid layer by layer, where each layer represents one minute of spreading rot.
* **Conceptual Link**: Demonstrates how Multi-Source BFS can find the shortest time for a process to spread across a grid from multiple origins simultaneously.

### 3. Surrounded Regions (LeetCode #130)
* **Problem Statement**: Given an `m x n` matrix `board` containing 'X' and 'O', capture all regions that are 4-directionally surrounded by 'X'. A region is captured by flipping all 'O's into 'X's in that surrounded region.
* **Solution Link**: [problems/matrix-traversal/matrix_traversal_problems.go](problems/matrix-traversal/matrix_traversal_problems.go) (`Solve`)
* **Explanation**: Uses **Reverse DFS**. Any 'O' on the border, and any 'O' connected to such a border 'O', cannot be captured. These "safe" cells are marked. Finally, all unmarked 'O's are flipped to 'X's.
* **Conceptual Link**: Illustrates the strategy of identifying "escapable" or "connected-to-boundary" nodes in a grid to solve containment problems.

### 4. Pacific Atlantic Water Flow (LeetCode #417)
* **Problem Statement**: There is an `m x n` rectangular island that borders both the Pacific Ocean and Atlantic Ocean. Find all grid coordinates from which water can flow to both the Pacific and Atlantic oceans.
* **Solution Link**: [problems/matrix-traversal/matrix_traversal_problems.go](problems/matrix-traversal/matrix_traversal_problems.go) (`PacificAtlantic`)
* **Explanation**: Performs DFS starting from the cells adjacent to each ocean and moving "uphill". Cells reachable from both oceans are the result.
* **Conceptual Link**: Another application of **Reverse Traversal**, where we start from the goal (the oceans) to find all valid starting positions.

* **Max Area of Island** (LeetCode #695)
* **01 Matrix** (LeetCode #542)

---

## 17.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 20: Elementary Graph Algorithms (Grid representations pp. 589–593)
* https://leetcode.com/problems/number-of-islands/
* https://techinterviewhandbook.org/algorithms/matrix/
