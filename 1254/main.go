package main

func main() {
	/*	println(closedIsland([][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}))*/

	println(closedIsland([][]int{
		{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
		{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 0}},
	))

}

func closedIsland(grid [][]int) (result int) {

	c := len(grid) - 1
	r := len(grid[0]) - 1

	for i := range grid {
		if grid[i][0] == 0 {
			removeIsland(grid, i, 0)
		}
		if grid[i][r] == 0 {
			removeIsland(grid, i, r)
		}
	}

	for j := range grid[0] {
		if grid[0][j] == 0 {
			removeIsland(grid, 0, j)
		}
		if grid[c][j] == 0 {
			removeIsland(grid, c, j)
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				removeIsland(grid, i, j)
				result++
			}
		}
	}
	return result
}

func removeIsland(grid [][]int, c int, r int) {
	if c < 0 || r < 0 || c == len(grid) || r == len(grid[0]) {
		return
	}

	if grid[c][r] == 1 {
		return
	}
	grid[c][r] = 1

	removeIsland(grid, c, r+1)
	removeIsland(grid, c+1, r)
	removeIsland(grid, c, r-1)
	removeIsland(grid, c-1, r)
}
