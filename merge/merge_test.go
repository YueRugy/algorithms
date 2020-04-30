package merge

import (
	"github.com/algorithms/util"
	"sort"
	"testing"
)

func TestMerge_SortTest(t *testing.T) {
	arr := util.CreateSlice(100000, 1000000)
	merge := NewMerge(arr)
	merge.SortTest()
	if sort.IntsAreSorted(merge.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	//fmt.Println(merge.Array)
}
