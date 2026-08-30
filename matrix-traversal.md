# 17 - Matrix Traversal (Grid BFS & DFS / Flood Fill)

## 17.1 - Overview

* The **Matrix Traversal** pattern treats a 2D grid matrix of size $M \times N$ as an implicit graph where each cell $(r, c)$ is connected to up to 4 orthogonal neighbors:
  $$\{(r-1, c), (r+1, c), (r, c-1), (r, c+1)\}$$
  (or 8 neighbors if diagonals are allowed).
* Common techniques include:
  1. **Flood Fill / DFS:** Sink visited components in-place (e.g. mutate `'1'` to `'0'`) or maintain a `boolean[][] visited`.
  2. **Multi-Source BFS:** Add all initial source locations to a queue simultaneously (e.g. Rotting Oranges, Walls and Gates) to calculate uniform distance fronts.

---

## 17.2 - Properties of a problem that suggests Matrix Traversal

* Input is a 2D array / grid with boundaries.
* Counting connected "islands", enclosures, or perimeter lengths.
* Shortest path on a maze / grid or time taken for a state to spread to all cells.

---

## 17.3 - Classic Example: Number of Islands (DFS)

### Java Implementation

```java
public class MatrixTraversal {

    public static int numIslands(char[][] grid) {
        if (grid == null || grid.length == 0) return 0;

        int islands = 0;
        int rows = grid.length, cols = grid[0].length;

        for (int r = 0; r < rows; r++) {
            for (int c = 0; c < cols; c++) {
                if (grid[r][c] == '1') {
                    islands++;
                    dfsSink(grid, r, c); // Sink the entire connected island
                }
            }
        }

        return islands;
    }

    private static void dfsSink(char[][] grid, int r, int c) {
        // Boundary and water check
        if (r < 0 || r >= grid.length || c < 0 || c >= grid[0].length || grid[r][c] != '1') {
            return;
        }

        grid[r][c] = '0'; // Mark cell as visited in-place

        // Explore 4 directions
        dfsSink(grid, r + 1, c);
        dfsSink(grid, r - 1, c);
        dfsSink(grid, r, c + 1);
        dfsSink(grid, r, c - 1);
    }
}
```

---

## 17.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(M \times N)$ where each cell is visited at most a constant number of times (once per neighbor).
* **Space Complexity:** $\mathcal{O}(M \times N)$ in the worst case for the recursion stack (or BFS queue) if the entire grid is land.

---

## 17.5 - Classic LeetCode Problems

* **Number of Islands** (LeetCode #200)
* **Max Area of Island** (LeetCode #695)
* **Rotting Oranges** (Multi-source BFS) (LeetCode #994)
* **Surrounded Regions** (LeetCode #130)
* **Pacific Atlantic Water Flow** (LeetCode #417)
* **01 Matrix** (LeetCode #542)
* **Shortest Path in Binary Matrix** (LeetCode #1091)

---

## 17.6 - Sources used for this file:
https://leetcode.com/problems/number-of-islands/ <br>
https://www.geeksforgeeks.org/flood-fill-algorithm/ <br>
https://techinterviewhandbook.org/algorithms/matrix/ <br>
https://leetcode.com/explore/interview/card/cheatsheets/720/resources/4723/
