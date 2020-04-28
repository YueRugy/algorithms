package insertion

import "github.com/algorithms/soft"

type Insertion1 struct {
	*soft.Soft
}

func NewInsertion1(arr []int) *Insertion1 {
	i := &Insertion1{
		Soft: &soft.Soft{
			Array: arr,
		},
	}
	return i
}
func (i *Insertion1) Sort() {
	for begin := 1; begin < len(i.Array); begin++ {
		node := i.Array[begin]
		lastIndex := begin
		for cur := begin; cur > 0; cur-- {
			if i.Compare(begin, cur-1) < 0 {
				lastIndex = cur - 1
			}
		}
		for index := begin; index > lastIndex; index-- {
			i.Array[index] = i.Array[index-1]
		}
		i.Array[lastIndex] = node
	}

}
