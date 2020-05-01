package shell

import (
	"time"

	"github.com/algorithms/soft"
)

type Shell struct {
	*soft.Soft
}

func NewShell(arr []int) *Shell {
	m := &Shell{
		Soft: &soft.Soft{
			Array: arr,
		},
	}
	return m
}

func (s *Shell) SortTest() {
	before := time.Now()
	sli := shellStepSequence(len(s.Array))
	for _, step := range sli {
		s.Sort(step)
	}
	after := time.Now()
	s.String(after.Sub(before).String())
}
func shellStepSequence(length int) []int {
	sli := make([]int, 0)
	for length > 1 {
		length >>= 1
		sli = append(sli, length)
	}
	return sli
}
func (s *Shell) Sort(step int) {
	for col := 0; col < step; col++ {
		for begin := col + step; begin < len(s.Array); begin += step {
			for cur := begin; cur > col; cur -= step {
				if s.Compare(cur, cur-step) < 0 {
					s.Swap(cur, cur-step)
				} else {
					break
				}
			}
		}
	}
}
