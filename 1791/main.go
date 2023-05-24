package main

func findCenter(edges [][]int) int {
	set := make([]uint32, 100_000)

	max := 0
	for i := range edges {
		set[edges[i][0]]++
		set[edges[i][1]]++

		if set[max] < set[edges[i][0]] {
			max = edges[i][0]
		}

		if set[max] < set[edges[i][1]] {
			max = edges[i][1]
		}

	}

	return max

}
