package main

func numEnclaves(grid [][]int) (result int) {
	c := len(grid) - 1
	r := len(grid[0]) - 1

	for i := range grid {
		if grid[i][0] == 1 {
			removeIslandSafe(grid, i, 0)
		}
		if grid[i][r] == 1 {
			removeIslandSafe(grid, i, r)
		}
	}

	for j := range grid[0] {
		if grid[0][j] == 1 {
			removeIslandSafe(grid, 0, j)
		}
		if grid[c][j] == 1 {
			removeIslandSafe(grid, c, j)
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				result += removeIslandUnSafe(grid, i, j)
			}
		}
	}
	return result
}

func removeIslandSafe(grid [][]int, c int, r int) int {
	if c < 0 || r < 0 || c == len(grid) || r == len(grid[0]) {
		return 0
	}

	if grid[c][r] == 0 {
		return 0
	}
	grid[c][r] = 0
	return 1 + removeIslandSafe(grid, c, r+1) + removeIslandSafe(grid, c+1, r) + removeIslandSafe(grid, c, r-1) + removeIslandSafe(grid, c-1, r)
}

func removeIslandUnSafe(grid [][]int, c int, r int) int {
	if grid[c][r] == 1 {
		return 0
	}
	grid[c][r] = 0
	return 1 + removeIslandUnSafe(grid, c, r+1) + removeIslandUnSafe(grid, c+1, r) + removeIslandUnSafe(grid, c, r-1) + removeIslandUnSafe(grid, c-1, r)
}
