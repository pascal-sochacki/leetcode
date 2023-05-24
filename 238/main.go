package main

func main() {
	//Input: nums = [1,2,3,4]
	//toRight: nums = [24,12,4,1]
	//toLeft: nums = [1,1,2,6]
	//Output: [24,12,8,6]
	t := productExceptSelf([]int{1, 2, 3, 4})
	println(t)
	//Input: nums =   [-1, 1, 0,-3, 3]
	//toRight: nums = [ 0, 0,-9, 3, 1]
	//toRight: nums = [ 1,-1,-1, 0, 0]
	//Output: [24,12,8,6]
	t = productExceptSelf([]int{-1, 1, 0, -3, 3})
	println(t)
}

func productExceptSelf(nums []int) []int {
	//toLeft := make([]int, len(nums))
	result := make([]int, len(nums))

	toRight := make([]int, len(nums))
	for i := len(toRight) - 1; i >= 0; i-- {

		if i == len(toRight)-1 {
			toRight[i] = 1
		} else if i == len(toRight)-2 {
			toRight[i] = nums[i+1]
		} else {
			toRight[i] = nums[i+1] * toRight[i+1]
		}

	}
	toLeft := make([]int, len(nums))
	for i := 0; i < len(toLeft); i++ {

		if i == 0 {
			toLeft[i] = 1

		} else if i == 1 {
			toLeft[i] = nums[i-1]
		} else {
			toLeft[i] = nums[i-1] * toLeft[i-1]
		}
	}
	for i := 0; i < len(result); i++ {
		result[i] = toLeft[i] * toRight[i]

	}
	return result
}
