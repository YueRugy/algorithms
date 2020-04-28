package insertion

import "github.com/algorithms/soft"

type Insertion struct {
	*soft.Soft
}

func NewInsertion(arr []int) *Insertion {
	i := &Insertion{
		Soft: &soft.Soft{
			Array: arr,
		},
	}
	return i
}
func (i *Insertion) Sort() {
	for begin := 1; begin < len(i.Array); begin++ {
		for cur := begin; cur > 0; cur-- {
			if i.Compare(cur, cur-1) < 0 {
				i.Swap(cur, cur-1)
			}
		}
	}

}
