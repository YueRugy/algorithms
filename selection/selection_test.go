package selection

import (
	"sort"
	"testing"
)

func TestSelection(t *testing.T) {
	arr := []int{12, 34, 5, 9, 79, 65, 23, 11}
	Selection(arr)
	if sort.IntsAreSorted(arr) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
func TestSelectionByHeap(t *testing.T) {
	arr := []int{12, 34, 5, 9, 79, 65, 23, 11}
	SelectionByHeap(arr)
	if sort.IntsAreSorted(arr) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
