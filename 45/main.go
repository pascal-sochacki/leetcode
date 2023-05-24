package main

import "math"

func main() {
	//print(jump([]int{2, 3, 1, 1, 4}))
	//print(jump([]int{2, 3, 0, 1, 4}))
	//print(jump([]int{1, 2, 3}))
	print(jump([]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}))
	//print(jump([]int{8, 0, 8, 5, 0, 1, 1, 0, 9, 4, 1, 5, 1, 3, 8, 0, 3, 2, 2, 7, 1, 7, 3, 4, 5, 3, 6, 1, 1, 6, 7, 9, 7, 5, 5, 9, 8, 6, 9, 7, 2, 2, 5, 5, 5, 3, 8, 0, 5, 0, 9, 9, 0, 3}))
}

func jump(nums []int) int {
	memory := make([]int, len(nums))
	res, _ := canJumpInner(nums, memory, 0)

	return res
}

func canJumpInner(nums []int, memory []int, pos int) (int, bool) {
	if pos >= len(nums)-1 {
		return 0, true
	}

	if nums[pos] == 0 {
		return math.MaxInt, false
	}

	if memory[pos] != 0 {
		return memory[pos], true
	}

	var max int
	if nums[pos] > len(nums) {
		max = len(nums)
	} else {
		max = nums[pos]
	}

	minJump := math.MaxInt
	for i := max; i > 0; i-- {
		temp, reached := canJumpInner(nums, memory, pos+i)
		if !reached {
			continue
		}
		if minJump >= temp {
			minJump = temp
			memory[pos] = minJump + 1
		}
	}
	if minJump == math.MaxInt {
		return memory[pos], false
	} else {
		return memory[pos], true
	}
}
