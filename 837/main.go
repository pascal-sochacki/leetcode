package main

func main() {

	println(new21Game(21, 17, 10))

}

func new21Game(n int, k int, maxPts int) float64 {
	// Corner cases
	if k == 0 {
		return 1.0
	}
	if n >= k-1+maxPts {
		return 1.0
	}

	// dp[i] is the probability of getting point i.
	dp := make([]float64, n+1)

	probability := 0.0 // dp of N or less points.
	windowSum := 1.0   // Sliding required window sum
	dp[0] = 1.0

	for i := 1; i <= n; i++ {
		dp[i] = windowSum / float64(maxPts)

		if i < k {
			windowSum += dp[i]
		} else {
			probability += dp[i]
		}

		if i-maxPts >= 0 {
			windowSum -= dp[i-maxPts]
		}
	}

	return probability
}
