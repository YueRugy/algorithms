package bubble

import "github.com/algorithms/soft"

type Bubble2 struct {
	*soft.Soft
}

func NewBubble2(arr []int) *Bubble2 {
	return &Bubble2{
		&soft.Soft{
			Array: arr,
		},
	}
}

func (b *Bubble2) Sort() {
	length := len(b.Array)
	for end := length - 1; end > 0; {
		lastIndex := 0
		for begin := 0; begin < end; begin++ {
			if b.Compare(begin, begin+1) > 0 {
				b.Swap(begin, begin+1)
				lastIndex = begin
			}
		}
		end = lastIndex
	}
}
