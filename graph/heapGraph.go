package graph

const defaultHeapCap = 16

type HeapGraph struct {
	size    int
	buckets []*Edge
	compare func(e1, e2 *Edge) int
}

func NewHeapGraph(collection []*Edge, compare func(e1, e2 *Edge) int) *HeapGraph {
	heap := &HeapGraph{compare: compare}
	if collection == nil {
		heap.buckets = make([]*Edge, defaultHeapCap)
		return heap
	}
	if len(collection) <= defaultHeapCap {
		heap.buckets = make([]*Edge, defaultHeapCap)
	} else {
		heap.buckets = make([]*Edge, len(collection)+len(collection)>>1)
	}
	copy(heap.buckets, collection)
	heap.size = len(collection)
	heap.heapify()
	return heap
}

func (h *HeapGraph) Add(e *Edge) {
	if e == nil {
		return
	}
	h.ensureCapacity(h.size + 1)
	if h.size == 0 {
		h.buckets[h.size] = e
		return
	}
	h.buckets[h.size] = e
	h.size++
	h.siftUp()
}

func (h *HeapGraph) Remove() *Edge {
	if h.size == 0 {
		return nil
	}

	if h.size == 1 {
		h.size--
		return h.buckets[0]
	}

	e := h.buckets[0]
	h.buckets[0] = h.buckets[h.size-1]
	h.size--
	h.siftDown(0)
	return e
}

func (h *HeapGraph) Size() int {
	return h.size
}

func (h *HeapGraph) IsEmpty() bool {
	return h.size == 0
}

func (h *HeapGraph) ensureCapacity(newCapacity int) {
	if newCapacity <= h.length() {
		return
	}
	newBuckets := make([]*Edge, h.length()+h.length()>>1)
	copy(newBuckets, h.buckets)
	h.buckets = newBuckets
}

func (h *HeapGraph) length() int {
	return len(h.buckets)
}

func (h *HeapGraph) siftUp() {
	e := h.buckets[h.size-1]
	si := h.size - 1
	for {
		if si <= 0 {
			break
		}
		fi := (si - 1) >> 1
		if h.compare(e, h.buckets[fi]) >= 0 {
			break
		} else {
			h.buckets[si] = h.buckets[fi]
			si = fi
		}
	}
	if si != h.size-1 {
		h.buckets[si] = e
	}
}

func (h *HeapGraph) siftDown(si int) {
	e := h.buckets[si]
	for {
		mi := h.findMaxOrMinIndex(si, e)
		if mi < 0 {
			break
		}
		h.buckets[si] = h.buckets[mi]
		si = mi
	}
	h.buckets[si] = e
}

func (h *HeapGraph) heapify() {
	for index := h.size >> 1; index >= 0; index-- {
		h.siftDown(index)
	}
}

func (h *HeapGraph) findMaxOrMinIndex(si int, e *Edge) int {
	left := si<<1 + 1
	if left > h.size-1 {
		return -1
	}
	if left == h.size-1 {
		if h.compare(e, h.buckets[left]) > 0 {
			return left
		}
		return -1
	} else {
		right := left + 1
		mi := left
		if h.compare(h.buckets[left], h.buckets[right]) > 0 {
			mi = right
		}
		if h.compare(e, h.buckets[mi]) > 0 {
			return mi
		}
		return -1
	}
}
