package heap

import "reflect"

const defaultCap = 16

type GenericHeap struct {
	size    int
	buckets []interface{}
	compare func(v1, v2 interface{}) int
}

func NewGenericHeap(c interface{},compare func(v1,v2 interface{}) int) *GenericHeap {
	collection:= toSlice(c)

	heap := &GenericHeap{}
	heap.compare=compare
	if collection == nil {
		heap.buckets = make([]interface{}, defaultCap)
		return heap
	}
	if len(collection) <= defaultCap {
		heap.buckets = make([]interface{}, defaultCap)
	} else {
		heap.buckets = make([]interface{}, len(collection)+len(collection)>>1)
	}
	copy(heap.buckets, collection)
	heap.size = len(collection)
	heap.heapify()
	return heap
}

func (h *GenericHeap) Add(e *interface{}) {
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

func (h *GenericHeap) Remove() interface{} {
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

func (h *GenericHeap) Size() int {
	return h.size
}

func (h *GenericHeap) IsEmpty() bool {
	return h.size == 0
}

func (h *GenericHeap) heapify() {
	for index := h.size >> 1; index >= 0; index-- {
		h.siftDown(index)
	}
}

func (h *GenericHeap) siftDown(si int) {
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

func (h *GenericHeap) findMaxOrMinIndex(si int, e interface{}) int {
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

func (h *GenericHeap) ensureCapacity(newCapacity int) {
	if newCapacity <= h.length() {
		return
	}
	newBuckets := make([]interface{}, h.length()+h.length()>>1)
	copy(newBuckets, h.buckets)
	h.buckets = newBuckets
}

func (h *GenericHeap) length() int {
	return len(h.buckets)
}

func (h *GenericHeap) siftUp() {
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
func toSlice(c interface{})  []interface{}{
	v:=reflect.ValueOf(c)
	if v.Kind()!=reflect.Slice{
		panic("unexpect type")
	}
	l:=v.Len()
	ret:=make([]interface{},l)
	for index:=0;index<l;index++{
		ret[index]=v.Index(index).Interface()
	}
	return ret
}
