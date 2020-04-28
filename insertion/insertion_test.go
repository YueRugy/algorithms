package insertion

import (
	"fmt"
	"sort"
	"testing"
)

func TestInsertion_Sort(t *testing.T) {
	arr := []int{12, 34, 5, 9, 79, 65, 23, 11}
	insert := NewInsertion(arr)
	insert.Sort()
	if sort.IntsAreSorted(insert.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(insert.Array)
}
func TestInsertion1_Sort(t *testing.T) {
	arr := []int{12, 34, 5, 9, 79, 65, 23, 11}
	insert := NewInsertion1(arr)
	insert.Sort()
	if sort.IntsAreSorted(insert.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(insert.Array)
}
