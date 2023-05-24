package main

func main() {

	print(canJump([]int{2, 3, 1, 1, 4}))
	print(canJump([]int{3, 2, 1, 0, 4}))
	print(canJump([]int{0}))
}

func canJump(nums []int) bool {
	memory := make([]bool, len(nums))
	return canJumpInner(nums, memory, 0)
}

func canJumpInner(nums []int, memory []bool, pos int) bool {
	if pos >= len(nums)-1 {
		return true
	}

	if memory[pos] {
		return false
	}
	var max int
	if nums[pos] > len(nums) {
		max = len(nums)
	} else {
		max = nums[pos]
	}

	for i := max; i > 0; i-- {
		if canJumpInner(nums, memory, pos+i) {
			return true
		}
	}
	memory[pos] = true
	return false
}
