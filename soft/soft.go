package soft

import (
	"fmt"
	"reflect"
)

type Soft struct {
	Array        []int
	SwapCount    int
	CompareCount int
}

func (soft *Soft) String(time string) {
	fmt.Println("struct  " + reflect.TypeOf(soft).String() + "  use " + time)
}

func (soft *Soft) Compare(index1, index2 int) int {
	soft.CompareCount++
	return soft.Array[index1] - soft.Array[index2]
}

func (soft *Soft) Swap(index1, index2 int) {
	soft.SwapCount++
	soft.Array[index1] ^= soft.Array[index2]
	soft.Array[index2] ^= soft.Array[index1]
	soft.Array[index1] ^= soft.Array[index2]
}
