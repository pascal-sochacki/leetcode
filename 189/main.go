package main

func main() {
	//	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	//	rotate([]int{-1, -100, 3, 99}, 2)
	rotate([]int{1, 2}, 3)
}

func rotate(nums []int, k int) {
	if len(nums) <= 0 {
		return
	}

	temp := make([]int, len(nums))
	for i, _ := range nums {

		temp[(i+k)%len(nums)] = nums[i]

	}
	copy(nums, temp)
}
