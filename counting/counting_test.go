package counting

import (
	"sort"
	"testing"

	"github.com/algorithms/util"
)

func TestCounting_SortTest(t *testing.T) {
	arr := util.CreateSlice(100000, 1000000)
	counting := NewCounting(arr)
	counting.SortTest()
	if sort.IntsAreSorted(counting.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	//fmt.Println(counting.Array)
}
