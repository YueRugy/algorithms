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

//func (i *Insertion1) Sort() {
//	for begin := 1; begin < len(i.Array); begin++ {
//		node := i.Array[begin]
//		cur := begin
//		for ; cur > 0 && i.CompareNum(node, i.Array[cur-1]) < 0; {
//			i.Array[cur] = i.Array[cur-1]
//			cur--
//		}
//		i.Array[cur] = node
//	}
//}
func (i *Insertion1) Sort() {
	for begin := 1; begin < len(i.Array); begin++ {
		node := i.Array[begin]
		cur := begin
		for ; cur > 0 && i.CompareNum(node, i.Array[cur-1]) < 0; cur-- {
			i.Array[cur] = i.Array[cur-1]
		}
		i.Array[cur] = node
	}
}
