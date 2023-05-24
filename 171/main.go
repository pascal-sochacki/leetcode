package main

import "math"

func main() {
	println(titleToNumber("A"))
	println(titleToNumber("B"))
	println(titleToNumber("C"))
	println(titleToNumber("AA"))
	println(titleToNumber("AB"))
}

func titleToNumber(columnTitle string) int {

	result := 0

	length := len(columnTitle) - 1

	for index, _ := range columnTitle {

		result += int(+columnTitle[length-index]-'A'+1) * int(math.Pow(26, float64(index)))
	}

	return result
}
