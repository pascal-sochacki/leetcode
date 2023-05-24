package main

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for i, s := range nums {

		if val, ok := m[target-s]; ok {
			return []int{val, i}
		}
		m[s] = i
	}

	return []int{}
}

func main() {

	test := []int{2, 7, 11, 15}
	sum := twoSum(test, 9)
	print(sum)
}
