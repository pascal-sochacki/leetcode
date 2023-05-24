package main

func main() {
	climbStairs(44)
}

func climbStairs(n int) int {

	save := make([]int, 45)
	return climbStairss(n, save)
}

func climbStairss(n int, save []int) int {
	if save[n] != 0 {
		return save[n]
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	save[n] = climbStairss(n-2, save) + climbStairss(n-1, save)

	return save[n]
}
