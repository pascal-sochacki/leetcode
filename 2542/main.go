package main

import (
	"container/heap"
	"sort"
)

func main() {
	println(maxScore([]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3))
}

type Pair struct {
	num1 int
	num2 int
}

type MinHeap []Pair

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i].num1 > m[j].num1
}

func (m MinHeap) Swap(i, j int) {
	m[j], m[i] = m[j], m[j]
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(Pair))
}

func (h *MinHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {

	arr := make([]Pair, len(nums1))

	for i := range nums1 {
		arr[i] = Pair{
			num1: nums1[i],
			num2: nums2[i],
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].num2 > arr[j].num2
	})

	minheap := &MinHeap{}
	sum := 0

	for i := 0; i < k; i++ {
		heap.Push(minheap, arr[i])
		sum += arr[i].num1
	}

	res := sum * arr[k-1].num2

	for i := k; i < len(arr); i++ {
		smallNum := heap.Pop(minheap).(Pair)
		sum += arr[i].num1 - smallNum.num1

		// O(logK)
		heap.Push(minheap, arr[i])

		res = max(res, sum*arr[i].num2)
	}

	return int64(res)

}
