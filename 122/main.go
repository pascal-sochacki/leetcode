package main

func main() {

	print(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	print(maxProfit([]int{1, 2, 3, 4, 5}))
	print(maxProfit([]int{7, 6, 3, 2, 1}))
}

func maxProfit(prices []int) (result int) {

	max := len(prices) - 1
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			result += prices[i+1] - prices[i]
		}
	}
	return

}
