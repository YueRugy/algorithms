package insertion

import (
	"fmt"
	"github.com/algorithms/util"
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
func TestInsertion2_Sort(t *testing.T) {
	arr :=util.CreateSlice(100000,1000000)
	insert := NewInsertion2(arr)
	insert.SortTest()
	if sort.IntsAreSorted(insert.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(insert.Array)
}
