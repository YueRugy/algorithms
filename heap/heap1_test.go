package heap

import (
	"fmt"
	"github.com/algorithms/mi"
	"sort"
	"testing"
)

func TestHeap_Size(t *testing.T) {
	heap := NewHeap1()
	num := heap.Size()
	if num == 0 {
		t.Log("success")
	} else {
		t.Log("error")
	}
}
func TestHeap_Add(t *testing.T) {

	array := []int{54, 80, 29, 79, 6, 58, 93, 86, 51, 65, 34, 39, 85, 26, 28, 95}
	heap := NewHeap1()
	for index, v := range array {
		node := &mi.Node{
			Index: index,
			Value: v,
		}
		heap.Add(node)
	}
	//	array1 := []int{6, 29, 26, 51, 34, 54, 28, 86, 80, 79, 65, 58, 85, 93, 39, 95}
	for _, node := range heap.arr {
		fmt.Print(node.Value, "\t")
	}
	t.Log("success")
}

//fmt.Println(heap.arr)
//flag := true
//for index := 0; index < len(array); index++ {
//	if heap.array[index] != array1[index] {
//		flag = false
//	}
//}
//if flag {
//	t.Log("success")
//} else {
//	t.Error("failed")
//}

func TestHeap_Remove(t *testing.T) {
	array := []int{54, 80, 29, 79, 6, 58, 93, 86, 51, 65, 34, 39, 85, 26, 28, 95}
	heap := NewHeap1()
	for index, v := range array {
		node := &mi.Node{
			Index: index,
			Value: v,
		}
		heap.Add(node)
	}
	sli := make([]*mi.Node, heap.Size())
	length := heap.Size()
	for i := 0; i < length; i++ {
		sli[i] = heap.Remove()
	}
	for _, node := range sli {
		fmt.Print(node.Value, "\t")
	}
	t.Log("success")
}

//func TestSlice(t *testing.T) {
//	array := []int{54, 80, 29, 79}
//	a2 := make([]int, len(array))
//	a2 = append(a2, array...)
//	t.Log(&array[0], &a2[0])
//	a1 := array[1:3]
//	a1 = append(a1, 6)
//	a1 = append(a1, 16)
//	//a1 = append(a1, array...)
//	//a1 = append(a1, 4)
//	t.Log(&array[1], &a1[0])
//	t.Log(array)
//
//}
//func TestNewHeapSlice(t *testing.T) {
//	array := []int{54, 80, 29, 79, 6, 58, 93, 86, 51, 65, 34, 39, 85, 26, 28, 95}
//	//array1 := []int{6, 29, 26, 51, 34, 54, 28, 86, 80, 79, 65, 58, 85, 93, 39, 95}
//	heap := NewHeapSlice(array)
//	t.Log(heap.array)
//	s := make([]int, len(array))
//	l := heap.size
//	for index := 0; index < l; index++ {
//		s[index] = heap.Remove()
//	}
//	t.Log(s)
//	t.Log(heap.array)
//}

//func TestTopK(t *testing.T) {
//	array := []int{54, 80, 29, 79, 6, 58, 93, 86, 51, 65, 34, 39, 85, 26, 28, 95}
//	sli := TopK(3, array)
//	fmt.Println(sli)
//}
//
//func TestHeapSort(t *testing.T) {
//	array := []int{54, 80, 29, 79, 6, 58, 93, 86, 51, 65, 34, 39, 85, 26, 28, 95}
//	array = HeapSort(array)
//	t.Log(array)
//}
func TestCopy(t *testing.T) {
	array := []int{54, 80, 29, 79, 6, 58, 93, 86, 51, 65, 34, 39, 85, 26, 28, 95}
	array1 := make([]int, len(array))
	copy(array1, array)
	fmt.Println(array1)
}

func TestNewHeapSlice(t *testing.T) {
	array := []int{54, 80, 29, 79, 6, 58, 93, 86, 51, 65, 34, 39, 85, 26, 28, 95}
	heap := NewHeap1Slice(array)
	for _, node := range heap.arr {
		fmt.Print(node.Value, "\t")
	}

}
func TestHeap_Remove2(t *testing.T) {
	array := []int{54, 80, 29, 79, 6, 58, 93, 86, 51, 65, 34, 39, 85, 26, 28, 95}
	heap := NewHeap1Slice(array)

	length := heap.Size()
	sli := make([]*mi.Node, length)
	for i := 0; i < length; i++ {
		sli[i] = heap.Remove()
	}
	for _, node := range sli {
		fmt.Print(node.Value, "\t")
	}
	t.Log("success")

}

func TestNewHeap(t *testing.T) {
	//95	86	93	80	65	85	29	79	51	6	34	39	58	26	28	54
	array := []int{54, 80, 29, 79, 6, 58, 93, 86, 51, 65, 34, 39, 85, 26, 28, 95}
	h := NewHeap(array)
	fmt.Println(h.Array)
}

func TestHeap_Sort(t *testing.T) {
	array := []int{54, 80, 29, 79, 6, 58, 93, 86, 51, 65, 34, 39, 85, 26, 28, 95}
	h := NewHeap(array)
	h.Sort()

	if sort.IntsAreSorted(h.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(h.Array)
}
