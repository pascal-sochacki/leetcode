package main

func main() {
	ints := []int{1, 2, 3, 0, 0, 0}
	merge(ints, 3, []int{2, 5, 6}, 3)

	println(ints)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	i := 0
	j := 0
	k := 0

	nums1Copy := make([]int, m)
	copy(nums1Copy, nums1[:m])

	for k < m+n {

		if i == m {
			nums1[k] = nums2[j]
			j++
		} else if j == n {
			nums1[k] = nums1Copy[i]
			i++
		} else {
			if nums1Copy[i] < nums2[j] {
				nums1[k] = nums1Copy[i]
				i++
			} else {
				nums1[k] = nums2[j]
				j++
			}
		}

		println(nums1[k], i, j)

		k++

	}
}
