package general

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	//"github.com/algorithms/general"
)

type IntSlice []int

func (is IntSlice) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is IntSlice) Swap(i, j int) {
	if i == j || is[i] == is[j] {
		return
	}
	is[i] ^= is[j]
	is[j] ^= is[i]
	is[i] ^= is[j]
}

func (is IntSlice) Length() int {
	return len(is)
}

func (is *IntSlice) Push(x interface{}) {
	if v, ok := x.(int); ok {
		*is = append(*is, v)
	}
}

func (is *IntSlice) Pop() interface{} {
	old := *is
	n := old.Length()
	x := old[n-1]
	*is = old[:n-1]
	return x
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func TestInit(t *testing.T) {
	var array = &IntSlice{2, 7, 3, 8}
	Init(array)
	Push(array, 1)
	for array.Length()>0{
		fmt.Printf("%d ",Pop(array))
	}
}
