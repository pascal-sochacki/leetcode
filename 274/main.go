package main

import "sort"

func main() {
	println(hIndex([]int{3, 0, 6, 1, 5}))
	println(hIndex([]int{3, 1, 1}))
	println(hIndex([]int{0}))
	println(hIndex([]int{100}))

}

func hIndex(citations []int) (result int) {

	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	for i := 1; i < len(citations)+1; i++ {
		if citations[i-1] == i {
			return i
		}
		if citations[i-1] < i {
			return i - 1
		}

		result += 1
	}

	return result
}
