package quick

import (
	"sort"
	"testing"

	"github.com/algorithms/util"
)

func TestQuick_SortTest(t *testing.T) {
	arr := util.CreateSlice(100000, 1000000)
	quick := NewQuick(arr)
	quick.SortTest()
	if sort.IntsAreSorted(quick.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	//fmt.Println(quick.Array)
}
