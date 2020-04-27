package bubble

import (
	"github.com/algorithms/soft"
)

type Bubble struct {
	*soft.Soft
}

func NewBubble(arr []int) *Bubble {
	return &Bubble{
		&soft.Soft{
			Array: arr,
		},
	}
}

func (b *Bubble) Sort() {
	length := len(b.Array)
	for end := length - 1; end > 1; end-- {
		for begin := 0; begin < end; begin++ {
			if b.Compare(begin, begin+1) > 0 {
				b.Swap(begin, begin+1)
			}
		}
	}
}
