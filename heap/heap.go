package heap

import "github.com/algorithms/soft"

type Heap struct {
	*soft.Soft
	size int
}

func (h *Heap) heapify() {
	for index := h.size >> 1; index >= 0; index-- {
		h.siftDown(index)
	}
}

func (h *Heap) siftDown(index int) {
	num := h.Array[index]
	for {
		maxIndex := h.findMaxIndex(index)
		if maxIndex < 0 || num > h.Array[maxIndex] {
			break
		}
		h.Array[index] = h.Array[maxIndex]
		index = maxIndex
	}
	h.Array[index] = num
}

func (h *Heap) findMaxIndex(index int) int {
	left := index<<1 + 1
	if left > h.size-1 {
		return -1
	} else if left == h.size-1 {
		return left
	} else {
		if h.Array[left] > h.Array[left+1] {
			return left
		}
		return left + 1
	}
}

func NewHeap(arr []int) *Heap {
	h := &Heap{
		size: len(arr),
		Soft: &soft.Soft{
			Array: arr,
		},
	}

	h.heapify()
	return h
}

func (h *Heap) Sort() {
	for h.size > 1 {
		h.size--
		h.Swap(0, h.size)
		h.siftDown(0)
	}
}
