package selection

import "github.com/algorithms/soft"

type Selection struct {
	*soft.Soft
}

func NewSelection(arr []int) *Selection {
	return &Selection{
		&soft.Soft{
			Array: arr,
		},
	}
}

func (s *Selection) Sort() {
	length := len(s.Array)

	for end := length - 1; end > 0; end-- {
		maxIndex := end
		for begin := 0; begin < end; begin++ {
			if s.Compare(begin, maxIndex) > 0 {
				maxIndex = begin
			}
		}
		if maxIndex != end {
			s.Swap(maxIndex, end)
		}
	}
}
