package main

func main() {

	println(stoneGameII([]int{1, 2, 3, 4, 5, 100}))
	println(stoneGameII([]int{77, 12, 64, 35, 28, 4, 87, 21, 20}))

}

func stoneGameII(piles []int) int {
	i, _ := inner(piles, true, 1, 0)

	return i
}

func inner(piles []int, player bool, m int, d int) (playerValue int, enemyValue int) {

	if len(piles) == 0 {
		return 0, 0
	}
	movesToMake := getMoves(m)

	if len(piles) <= movesToMake {
		score := sum(piles)
		if player {
			return score, 0
		} else {
			return 0, score
		}
	}

	bestev := -10
	bestpv := -10

	for i := 0; i < movesToMake; i++ {

		v := sum(piles[:i+1])

		pv, ev := inner(piles[i+1:], !player, i+1, d+1)

		if d == 0 {
			println(pv+v, " ", ev)
		}

		if player {
			if (pv + v) >= bestpv {
				bestpv = pv + v
				bestev = ev
			}
		} else {
			if (ev + v) >= bestev {
				bestpv = pv
				bestev = ev + v
			}
		}
	}

	return bestpv, bestev
}

func getMoves(m int) int {
	return 2 * m
}

func sum(piles []int) (result int) {
	for _, pile := range piles {
		result += pile
	}
	return
}
