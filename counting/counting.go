package counting

import (
	"time"

	"github.com/algorithms/soft"
)

type Counting struct {
	*soft.Soft
}

func NewCounting(arr []int) *Counting {
	m := &Counting{
		Soft: &soft.Soft{
			Array: arr,
		},
	}
	return m
}

func (c *Counting) SortTest() {
	before := time.Now()

	c.Sort()
	after := time.Now()
	c.String(after.Sub(before).String())
}

func (c *Counting) Sort() {
	max := c.Array[0]
	for index := 1; index < len(c.Array); index++ {
		if c.CompareNum(max, c.Array[index]) < 0 {
			max = c.Array[index]
		}
	}

	sli := make([]int, max+1)
	for index := 0; index < len(c.Array); index++ {
		sli[c.Array[index]]++
	}

	count := 0
	for index, v := range sli {
		for v > 0 {
			c.Array[count] = index
			count++
			v--
		}
	}
}
