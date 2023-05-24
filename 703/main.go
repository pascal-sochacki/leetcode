package main

import (
	"math"
	"sort"
)

func main() {
	test := Constructor(2, []int{0})
	println(test.Add(-1))
}

type KthLargest struct {
	list []int
	max  int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return KthLargest{
		list: nums,
		k:    k,
	}
}

func (this *KthLargest) Add(val int) int {
	var newList = make([]int, len(this.list)+1)
	newList[len(this.list)] = math.MinInt
	copy(newList, this.list)

	if len(this.list) == 0 {
		this.list = []int{val}
		return val
	}

	for i, _ := range newList {
		if newList[i] > val {
			continue
		}

		if i == len(this.list) {
			newList[i] = val
			break
		}

		newList[i] = val
		copy(newList[i+1:], this.list[i:])
		break
	}

	this.list = newList

	return this.list[this.k-1]
}
