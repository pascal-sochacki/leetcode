package main

func main() {
	//r := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	r := dailyTemperatures([]int{30, 40, 50, 60})
	print(r)
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		for j := i; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				result[i] = j - i
				break
			}

		}
	}
	return result
}
