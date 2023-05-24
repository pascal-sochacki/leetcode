package main

func main() {
	//println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	//println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
	println(removeDuplicates([]int{1, 1, 1}))

}

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	curr := nums[0]
	currIndex := 0
	k := len(nums)

	for i := 1; i < k; i++ {

		if i == k-1 {
			if curr == nums[i] {
				if i-currIndex >= 2 {
					return k - (i - 1 - currIndex)

				}
			}
		}

		if curr != nums[i] {

			if i-2 <= currIndex {

				curr = nums[i]
				currIndex = i
				continue

			} else {
				copy(nums[currIndex:], nums[i-2:])

				removed := i - 2 - currIndex

				k -= removed

				i = currIndex + 2
				curr = nums[i]
				currIndex = i
				continue
			}

		}

	}
	return k

}
