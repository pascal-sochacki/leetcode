package main

import "container/heap"

func main() {

	println(matchPlayersAndTrainers([]int{4, 7, 9}, []int{8, 2, 5, 8}))

}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]

	return x
}

func matchPlayersAndTrainers(players []int, trainers []int) (result int) {

	playerHeap := &IntHeap{}
	heap.Init(playerHeap)

	for _, player := range players {
		heap.Push(playerHeap, player)
	}

	trainerHeap := &IntHeap{}
	heap.Init(trainerHeap)

	for _, trainer := range trainers {
		heap.Push(trainerHeap, trainer)
	}

	for {
		if playerHeap.Len() == 0 || trainerHeap.Len() == 0 {
			break
		}

		player := heap.Pop(playerHeap).(int)

		for {
			if trainerHeap.Len() == 0 {
				return
			}

			trainer := heap.Pop(trainerHeap).(int)

			if trainer >= player {
				result += 1
				break
			}

		}
	}

	return
}
