package bs

import (
	"github.com/algorithms/insertion"
	"sort"
	"testing"
)

func TestIndex(t *testing.T) {
	arr := []int{12, 34, 5, 9, 79, 65, 23, 11}
	insert := insertion.NewInsertion1(arr)
	insert.Sort()
	if sort.IntsAreSorted(insert.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	t.Log(insert.Array)
	index := Index(insert.Array, 12)
	if index == 3 {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
