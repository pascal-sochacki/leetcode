package main

import "math"

func main() {

	println(stoneGameIII([]int{-1, -2, -3}))

}

func stoneGameIII(piles []int) string {
	i, j := inner(piles, true, 1, 0)

	println(i, j)

	if i > j {
		return "Alice"
	}

	if i < j {
		return "Bob"
	}

	return "Tie"
}

func inner(piles []int, player bool, m int, d int) (playerValue int, enemyValue int) {

	if len(piles) == 0 {
		return 0, 0
	}
	movesToMake := 3

	if len(piles) < movesToMake {
		movesToMake = len(piles)
	}

	max := math.MinInt
	bestev := 0
	bestpv := 0

	for i := 0; i < movesToMake; i++ {

		v := sum(piles[:i+1])

		pv, ev := inner(piles[i+1:], !player, i+1, d+1)

		if d == 0 {
			println(pv+v, " ", ev)
		}

		if player {
			if max < (pv + v) {
				bestpv = pv + v
				bestev = ev
				max = bestpv
			}
		} else {
			if max < (ev + v) {
				bestpv = pv
				bestev = ev + v
				max = bestev
			}
		}
	}

	return bestpv, bestev
}

func sum(piles []int) (result int) {
	for _, pile := range piles {
		result += pile
	}
	return
}
