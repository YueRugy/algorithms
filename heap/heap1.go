package heap

import "github.com/algorithms/mi"

const defaultSize = 16

type Heap1 struct {
	size int
	arr  []*mi.Node
}

func (h *Heap1) Size() int {
	return h.size
}

func (h *Heap1) Empty() bool {
	return h.size == 0
}

func (h *Heap1) Add(node *mi.Node) {
	if node == nil {
		return
	}
	h.ensureCap(h.size + 1)
	h.arr[h.size] = node
	h.siftUp(h.size)
	h.size++
}

func (h *Heap1) Get() *mi.Node {
	return h.arr[0]
}

func (h *Heap1) Clear() {
	h.size = 0
	h.arr = nil
}

func (h *Heap1) Remove() *mi.Node {
	if h.Size() == 0 {
		return nil
	}
	node := h.arr[0]
	h.arr[0] = h.arr[h.size-1]
	h.arr[h.size-1] = nil
	h.size--
	h.siftDown(0)
	return node
}

func (h *Heap1) Replace(node *mi.Node) *mi.Node {
	oldNode := h.arr[0]
	h.arr[0] = node
	h.siftDown(0)
	return oldNode

}

func (h *Heap1) ensureCap(num int) {
	if num > len(h.arr) {
		newArr := make([]*mi.Node, len(h.arr)<<1)
		copy(newArr, h.arr)
		h.arr = newArr
	}

}

func (h *Heap1) siftUp(selfIndex int) {
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

func (h *Heap1) siftDown(selfIndex int) {
	node := h.arr[selfIndex]
	for {
		maxIndex := h.maxIndex(selfIndex)
		if maxIndex < 0 || h.arr[maxIndex].Value < node.Value {
			break
		}
		h.arr[selfIndex] = h.arr[maxIndex]
		selfIndex = maxIndex
	}
	h.arr[selfIndex] = node
}

func (h *Heap1) maxIndex(index int) int {
	left := index<<1 + 1
	if left >= h.size {
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

func (h *Heap1) heapify() {
	for index := h.size >> 1; index >= 0; index-- {
		h.siftDown(index)
	}
}

func NewHeap1() *Heap1 {
	return &Heap1{
		size: 0,
		arr:  make([]*mi.Node, defaultSize),
	}
}

func NewHeap1Slice(arr []int) *Heap1 {
	nodeSlice := make([]*mi.Node, len(arr))
	count := 0
	for index, v := range arr {
		node := &mi.Node{
			Index: index,
			Value: v,
		}
		nodeSlice[count] = node
		count++
	}
	heap := NewHeap1()
	heap.size = len(nodeSlice)
	heap.arr = make([]*mi.Node, heap.size)
	copy(heap.arr, nodeSlice)
	heap.heapify()
	return heap
}
