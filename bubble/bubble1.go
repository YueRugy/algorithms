package bubble

import "github.com/algorithms/soft"

type Bubble1 struct {
	*soft.Soft
}

func NewBubble1(arr []int) *Bubble1 {
	return &Bubble1{
		&soft.Soft{
			Array: arr,
		},
	}
}

func (b *Bubble1) Sort() {
	length := len(b.Array)
	for end := length - 1; end > 0; end-- {
		flag := true
		for begin := 0; begin < end; begin++ {
			if b.Compare(begin, begin+1) > 0 {
				flag = false
				b.Swap(begin, begin+1)
			}
		}
		if flag {
			break
		}
	}
}
