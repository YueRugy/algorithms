package heap

import "github.com/algorithms/mi"

const defaultSize = 16

type Heap struct {
	size int
	arr  []*mi.Node
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Empty() bool {
	return h.size == 0
}

func (h *Heap) Add(node *mi.Node) {
	if node == nil {
		return
	}
	h.ensureCap(h.size + 1)
	h.arr[h.size] = node
	h.siftUp(h.size)
	h.size++
}

func (h *Heap) Get() *mi.Node {
	return h.arr[0]
}

func (h *Heap) Clear() {
	h.size = 0
	h.arr = nil
}

func (h *Heap) Remove() *mi.Node {
	if h.Size() == 0 {
		return nil
	}
	node := h.arr[0]
	h.arr[0] = h.arr[h.size-1]
	h.arr[h.size-1] = nil
	h.size--
	h.siftDown()
	return node
}

func (h *Heap) Replace(node *mi.Node) *mi.Node {
	oldNode := h.arr[0]
	h.arr[0] = node
	h.siftDown()
	return oldNode

}

func (h *Heap) ensureCap(num int) {
	if num > len(h.arr) {
		newArr := make([]*mi.Node, len(h.arr)<<1)
		copy(newArr, h.arr)
		h.arr = newArr
	}

}

func (h *Heap) siftUp(selfIndex int) {
	node := h.arr[h.size]
	for selfIndex > 0 {
		pin := (selfIndex - 1) >> 1
		if node.Value > h.arr[pin].Value {
			h.arr[selfIndex] = h.arr[pin]
			selfIndex = pin
		} else {
			break
		}
	}
	h.arr[selfIndex] = node
}

func (h *Heap) siftDown() {
	node := h.arr[0]
	selfIndex := 0
	for {
		maxIndex := h.maxIndex(selfIndex)
		if maxIndex < 0 {
			break
		}
		h.arr[selfIndex] = h.arr[maxIndex]
		selfIndex = maxIndex
	}
	h.arr[selfIndex] = node
}

func (h *Heap) maxIndex(index int) int {
	left := index<<1 + 1
	if left > h.size {
		return -1
	} else if left == h.size-1 {
		return left
	} else {
		if h.arr[left].Value > h.arr[left+1].Value {
			return left
		}
		return left + 1
	}

}

func NewHeap() *Heap {
	return &Heap{
		size: 0,
		arr:  make([]*mi.Node, defaultSize),
	}
}
