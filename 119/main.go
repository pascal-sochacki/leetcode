package main

import "fmt"

func main() {
	//fmt.Printf("%v", getRow(0))
	//fmt.Printf("%v", getRow(1))
	fmt.Printf("%v", getRow(2))
	fmt.Printf("%v", getRow(3))
}

func getRow(rowIndex int) []int {

	if rowIndex == 0 {
		return []int{1}
	}

	last := getRow(rowIndex - 1)
	result := make([]int, rowIndex+1)

	result[0] = 1
	result[rowIndex] = 1

	for i := 1; i < rowIndex; i++ {
		result[i] = last[i-1] + last[i]
	}

	return result

}
