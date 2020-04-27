package selection

import (
	"fmt"
	"sort"
	"testing"
)

func TestSelection_Sort(t *testing.T) {
	arr := []int{12, 34, 5, 9, 79, 65, 23, 11}
	s := NewSelection(arr)
	s.Sort()
	if sort.IntsAreSorted(s.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(s.Array)
}
