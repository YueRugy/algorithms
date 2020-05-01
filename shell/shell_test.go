package shell

import (
	"fmt"
	"sort"
	"testing"

	"github.com/algorithmsRepeat/util"
)

func TestShell_SortTest(t *testing.T) {
	arr := util.NewSlice(10000, 100000)
	s := NewShell(arr)
	s.SortTest()
	if sort.IntsAreSorted(s.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(s.Array)
}
