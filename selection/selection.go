package selection

import (
	"github.com/algorithms/heap"
)

func Selection(arr []int) {
	length := len(arr)

	for end := length - 1; end > 0; end-- {
		index := 0
		for begin := 0; begin <= end; begin++ {
			if arr[begin] > arr[index] {
				index = begin
			}
		}

		if index != end {
			arr[index] ^= arr[end]
			arr[end] ^= arr[index]
			arr[index] ^= arr[end]
		}

	}
}

func SelectionByHeap(arr []int) {
	heap := heap.NewHeapSlice(arr)
	length := len(arr)
	for index := length - 1; index >= 0; index-- {
		arr[index] = heap.Remove().Value
	}
}
