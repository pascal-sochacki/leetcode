package main

func main() {
	grid1 := [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
	grid2 := [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}
	println(countSubIslands(grid1, grid2))
}

func countSubIslands(grid1 [][]int, grid2 [][]int) (result int) {
	for i := range grid2 {
		for j := range grid2[i] {
			if grid2[i][j] == 1 {

				missing := removeIsland(grid1, grid2, i, j)

				if missing == 0 {
					result++
				}

			}
		}
	}
	return result
}

func removeIsland(grid1 [][]int, grid2 [][]int, c int, r int) (result int) {

	if c < 0 || r < 0 || c >= len(grid2) || r >= len(grid2[0]) {
		return 0
	}

	if grid2[c][r] != 1 {
		return 0
	}

	if grid1[c][r] == 0 {
		result++
	}

	grid2[c][r] = 0

	result += removeIsland(grid1, grid2, c-1, r)
	result += removeIsland(grid1, grid2, c+1, r)
	result += removeIsland(grid1, grid2, c, r-1)
	result += removeIsland(grid1, grid2, c, r+1)

	return result
}
