package main

func main() {
	println(searchInsert([]int{}, 3))
	println(searchInsert([]int{1, 3, 5, 6}, 3))
	println(searchInsert([]int{1, 3, 5, 6}, 7))
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] < target {
			return 1
		} else {
			return 0
		}
	}

	middle := len(nums) / 2
	value := nums[middle]
	if value == target {
		return middle
	}

	if value < target {
		return searchInsert(nums[middle:], target) + middle
	} else {
		return searchInsert(nums[0:middle], target)
	}

}
