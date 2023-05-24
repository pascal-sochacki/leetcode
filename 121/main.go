package main

func main() {
	profit := maxProfit([]int{7, 1, 5, 3, 6, 4})
	println(profit)
}

func maxProfit(prices []int) (maxProfit int) {

	if len(prices) == 0 {
		return 0
	}

	low := prices[0]
	high := prices[0]

	for i := range prices {

		if low > prices[i] {
			low = prices[i]
			high = prices[i]
		}

		if high < prices[i] {
			high = prices[i]
		}

		if high-low > maxProfit {
			maxProfit = high - low
		}

	}

	return maxProfit

}
