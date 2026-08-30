package matrix_traversal

// =============================================================
// NumIslands — LeetCode #200
// DFS flood-fill: mark visited land cells with '0'.
// Time: O(m·n), Space: O(m·n) recursion stack in worst case.
// =============================================================

// NumIslands counts the number of islands in a 2-D binary grid.
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	count := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0' // mark visited
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}
	return count
}

// =============================================================
// OrangesRotting — LeetCode #994
// Multi-source BFS from all initially rotten oranges.
// Time: O(m·n), Space: O(m·n).
// =============================================================

// OrangesRotting returns the minimum minutes until no fresh orange remains,
// or -1 if impossible.
func OrangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	type point struct{ r, c int }
	queue := []point{}
	fresh := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, point{r, c})
			} else if grid[r][c] == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	minutes := 0

	for len(queue) > 0 && fresh > 0 {
		size := len(queue)
		minutes++
		for i := 0; i < size; i++ {
			p := queue[i]
			for _, d := range dirs {
				nr, nc := p.r+d[0], p.c+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					fresh--
					queue = append(queue, point{nr, nc})
				}
			}
		}
		queue = queue[size:]
	}

	if fresh > 0 {
		return -1
	}
	return minutes
}

// =============================================================
// Solve — LeetCode #130 (Surrounded Regions)
// Reverse DFS: mark all 'O' connected to borders as safe ('S'),
// then flip remaining 'O' → 'X' and restore 'S' → 'O'.
// Time: O(m·n), Space: O(m·n).
// =============================================================

// Solve modifies the board in-place: surrounded 'O' regions become 'X'.
func Solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	rows, cols := len(board), len(board[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != 'O' {
			return
		}
		board[r][c] = 'S' // safe
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	// Mark border-connected 'O' cells.
	for r := 0; r < rows; r++ {
		dfs(r, 0)
		dfs(r, cols-1)
	}
	for c := 0; c < cols; c++ {
		dfs(0, c)
		dfs(rows-1, c)
	}

	// Flip surrounded 'O' → 'X'; restore 'S' → 'O'.
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == 'S' {
				board[r][c] = 'O'
			}
		}
	}
}

// =============================================================
// PacificAtlantic — LeetCode #417
// Reverse DFS/BFS from each ocean's border; collect cells
// reachable from both oceans.
// Time: O(m·n), Space: O(m·n).
// =============================================================

// PacificAtlantic returns all coordinates reachable from both oceans.
func PacificAtlantic(heights [][]int) [][2]int {
	if len(heights) == 0 {
		return nil
	}
	rows, cols := len(heights), len(heights[0])
	pacific := make([][]bool, rows)
	atlantic := make([][]bool, rows)
	for r := 0; r < rows; r++ {
		pacific[r] = make([]bool, cols)
		atlantic[r] = make([]bool, cols)
	}

	var dfs func(r, c int, visited [][]bool, prev int)
	dfs = func(r, c int, visited [][]bool, prev int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || visited[r][c] || heights[r][c] < prev {
			return
		}
		visited[r][c] = true
		dfs(r+1, c, visited, heights[r][c])
		dfs(r-1, c, visited, heights[r][c])
		dfs(r, c+1, visited, heights[r][c])
		dfs(r, c-1, visited, heights[r][c])
	}

	for r := 0; r < rows; r++ {
		dfs(r, 0, pacific, heights[r][0])
		dfs(r, cols-1, atlantic, heights[r][cols-1])
	}
	for c := 0; c < cols; c++ {
		dfs(0, c, pacific, heights[0][c])
		dfs(rows-1, c, atlantic, heights[rows-1][c])
	}

	result := [][2]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pacific[r][c] && atlantic[r][c] {
				result = append(result, [2]int{r, c})
			}
		}
	}
	return result
}
