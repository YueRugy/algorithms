package general

import (
	"math/rand"
	"testing"
	"time"
	""
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
	panic("unexpect type")
}

func (is *IntSlice) Pop() interface{} {
	old := *is
	n := old.Length()
	x := old[n-1]
	*is = old[:n]
	return x
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func TestInit(t *testing.T) {
	var array  = make(IntSlice, 10)
	for index := 0; index < 10; index++ {
		array = append(array, rand.Intn(100))
	}
	heap.Push()


}
