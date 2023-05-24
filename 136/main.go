package main

func main() {
	println(singleNumber([]int{1, 2, 2, 1, 3}))
}

func singleNumber(nums []int) int {
	a := 2
	a ^= 3

	s := map[int]bool{}

	for i := range nums {

		_, ok := s[nums[i]]
		if ok {
			delete(s, nums[i])
		} else {
			s[nums[i]] = true
		}

	}

	for k := range s {
		return k
	}
	return 0
}
