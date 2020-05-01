package shell

import (
	"fmt"
	"math"
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
	//sli := shellStepSequence(len(s.Array))
	//for _, step := range sli {
	//	s.Sort(step)
	//}
	sli := robertStepSequence(len(s.Array))

	for index := len(sli) - 1; index >= 0; index-- {
		s.Sort(sli[index])
	}

	fmt.Println(sli)
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

func robertStepSequence(length int) []int {
	k, step := 0, 0
	sli := make([]int, 0)
	for {
		if k&1 == 0 {
			step = int(9*(math.Pow(2.0, float64(k))-math.Pow(2.0, float64(k>>1)))) + 1
		} else {
			step = int(8*math.Pow(2.0, float64(k)) - 6*math.Pow(2.0, float64((k+1)>>1)) + 1)
		}
		if step >= length {
			break
		}
		k++
		sli = append(sli, step)
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
