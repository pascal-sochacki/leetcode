package main

func main() {
	println(removeDuplicates([]int{1, 1, 2}))
}

func removeDuplicates(nums []int) int {

	var current = nums[0]
	offset := 0
	for i := range nums {

		if nums[offset] != nums[i] {
			offset++
			current = nums[i]
			nums[offset] = current
		}

	}

	return offset + 1

}
