package main

func majorityElement(nums []int) (result int) {

	count := map[int]int{}

	for _, i := range nums {
		count[i]++

		if count[i] > count[result] {
			result = i
		}
	}
	return result
}
