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

* **Number of Islands** (LeetCode #200)
* **Rotting Oranges** (Multi-Source BFS) (LeetCode #994)
* **Max Area of Island** (LeetCode #695)
* **Surrounded Regions** (LeetCode #130)
* **Pacific Atlantic Water Flow** (LeetCode #417)
* **01 Matrix** (LeetCode #542)

---

## 17.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 20: Elementary Graph Algorithms (Grid representations pp. 589–593)
* https://leetcode.com/problems/number-of-islands/
* https://techinterviewhandbook.org/algorithms/matrix/
