package soft

import (
	"fmt"
	"strconv"
)

type Soft struct {
	Array        []int
	SwapCount    int
	CompareCount int
}

func (soft *Soft) String(time string) {
	fmt.Println("use " + time + " compare= " + strconv.Itoa(soft.CompareCount) + "  swap= " +
		strconv.Itoa(soft.SwapCount))
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
func (soft *Soft) CompareNum(num1, num2 int) int {
	soft.CompareCount++
	return num1 - num2
}
