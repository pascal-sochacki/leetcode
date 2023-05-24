package main

import "container/heap"

func main() {
	//topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	//topKFrequent([]int{1, 1, 1, 2, 2, 3333}, 2)
	//topKFrequent([]int{5, 3, 1, 1, 1, 3, 73, 1}, 2)
	topKFrequent([]int{5, -3, 9, 1, 7, 7, 9, 10, 2, 2, 10, 10, 3, -1, 3, 7, -9, -1, 3, 3}, 3)
}

type IntHeap [][]int

func (h IntHeap) Less(a, b int) bool { return h[a][1] < h[b][1] }
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }

func (h *IntHeap) Push(a interface{}) {
	*h = append(*h, a.([]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]

	return x
}

func topKFrequent(nums []int, k int) []int {

	set := make(map[int]int, 0)
	for i := range nums {
		set[nums[i]] += 1
	}

	h := &IntHeap{}
	heap.Init(h)

	for key, val := range set {
		heap.Push(h, []int{key, val})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, 0)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).([]int)[0])
	}

	return res
}
