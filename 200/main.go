package main

func main() {
	println(numIslands([][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}))

	println(numIslands([][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}))
}

func numIslands(grid [][]int) (max int) {
	for i := range grid {
		for i2 := range grid[i] {
			if grid[i][i2] == 1 {
				tmp := floodFill(grid, i, i2)
				if tmp > max {
					max = tmp
				}
			}

		}
	}
	return max

}

func floodFill(grid [][]int, r int, c int) (result int) {

	grid[r][c] = 0
	result++

	if r-1 >= 0 && grid[r-1][c] == 1 {
		result += floodFill(grid, r-1, c)
	}

	if r+1 < len(grid) && grid[r+1][c] == 1 {
		result += floodFill(grid, r+1, c)
	}

	if c-1 >= 0 && grid[r][c-1] == 1 {
		result += floodFill(grid, r, c-1)
	}

	if c+1 < len(grid[0]) && grid[r][c+1] == 1 {
		result += floodFill(grid, r, c+1)
	}

	return result
}
