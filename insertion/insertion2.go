package insertion

import (
	"github.com/algorithms/soft"
	"time"
)

type Insertion2 struct {
	*soft.Soft
}

func NewInsertion2(arr []int) *Insertion2 {
	i := &Insertion2{
		Soft: &soft.Soft{
			Array: arr,
		},
	}
	return i
}
func (i *Insertion2) SortTest() {
	before := time.Now()
	i.Sort()
	after := time.Now()
	i.String(after.Sub(before).String())
}
func (i *Insertion2) Sort() {
	for begin := 1; begin < len(i.Array); begin++ {
		index := i.Index(begin)
		num := i.Array[begin]
		for tempIndex := begin; tempIndex > index; tempIndex-- {
			i.Array[tempIndex] = i.Array[tempIndex-1]
		}
		i.Array[index] = num
	}
}
func (i *Insertion2) Index(index int) int {
	begin, end, num := 0, index, i.Array[index]
	for begin < end {
		mid := (begin + end) >> 1
		if i.CompareNum(num, i.Array[mid]) < 1 {
			end = mid
		} else {
			begin = mid + 1
		}
	}
	return begin
}
