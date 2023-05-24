package heap

import "errors"

type Heap struct {
	array     []int
	maxLength int
	length    int
}

func (h *Heap) len() int {
	return h.length
}

func (h *Heap) add(i int) {

	if h.length == h.maxLength {
		h.maxLength *= 2
		temp := make([]int, h.maxLength)
		copy(temp, h.array)
		h.array = temp

	}

	h.array[h.length] = i
	h.bubbleUp(h.length)
	h.length = h.length + 1
}

func (h *Heap) bubbleUp(pos int) {
	if pos == 0 {
		return
	}
	if h.parentValue(pos) > h.array[pos] {
		h.swapWithParent(pos)
		h.bubbleUp(h.getParentIndex(pos))
	}
}

func (h *Heap) heapifyDown(pos int) {
	if pos >= h.length {
		return
	}

	if h.getLeftChildIndex(pos) >= h.length {
		return
	}

	leftChildIndex := h.getLeftChildIndex(pos)
	if leftChildIndex == h.length-1 {
		if h.array[pos] > h.array[leftChildIndex] {
			h.swap(pos, leftChildIndex)
			h.heapifyDown(leftChildIndex)
			return
		}
	} else {
		rightChildIndex := h.getRightChildIndex(pos)
		if h.array[rightChildIndex] > h.array[leftChildIndex] {
			if h.array[pos] > h.array[leftChildIndex] {
				h.swap(pos, leftChildIndex)
				h.heapifyDown(leftChildIndex)
				return
			}
		} else {
			if h.array[pos] > h.array[rightChildIndex] {
				h.swap(pos, rightChildIndex)
				h.heapifyDown(rightChildIndex)
				return
			}
		}

	}

}
func (h *Heap) parentValue(pos int) int {
	return h.array[h.getParentIndex(pos)]
}

func (h *Heap) setParentValue(pos int, val int) {
	h.array[h.getParentIndex(pos)] = val
}

func (h *Heap) getRightChildIndex(pos int) int {
	return pos*2 + 2
}

func (h *Heap) getLeftChildIndex(pos int) int {
	return pos*2 + 1
}

func (h *Heap) getParentIndex(pos int) int {
	return (pos - 1) / 2
}

func (h *Heap) swapWithParent(pos int) {
	h.swap(pos, h.getParentIndex(pos))
}

func (h *Heap) swap(pos1 int, pos2 int) {
	temp := h.array[pos1]
	h.array[pos1] = h.array[pos2]
	h.array[pos2] = temp
}

func (h *Heap) pop() (int, error) {
	if h.length == 0 {
		return 0, errors.New("no Items left")
	}

	result := h.array[0]

	h.length = h.length - 1
	top := h.array[h.length]
	h.array[h.length] = 0
	h.array[0] = top
	h.heapifyDown(0)
	return result, nil

}

func NewHeap() *Heap {
	return &Heap{
		length:    0,
		maxLength: 5,
		array:     make([]int, 5),
	}
}
