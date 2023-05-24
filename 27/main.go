package main

func main() {

	println(removeElement([]int{3, 2, 2, 3}, 3))
	println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))

}

func removeElement(nums []int, val int) int {
	offset := 0
	for i := range nums {

		if nums[i] == val {
			offset++
		} else {
			nums[i-offset] = nums[i]
		}

	}

	return offset + 1
}
