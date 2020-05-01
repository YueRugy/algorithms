package counting

import (
	"time"

	"github.com/algorithms/soft"
)

type Counting1 struct {
	*soft.Soft
}

func NewCounting1(arr []int) *Counting1 {
	m := &Counting1{
		Soft: &soft.Soft{
			Array: arr,
		},
	}
	return m
}

func (c *Counting1) SortTest() {
	before := time.Now()

	c.Sort()
	after := time.Now()
	c.String(after.Sub(before).String())
}

func (c *Counting1) Sort() {
	max, min := c.Array[0], c.Array[0]
	for index := 1; index < len(c.Array); index++ {
		if c.CompareNum(max, c.Array[index]) < 0 {
			max = c.Array[index]
		} else if c.CompareNum(min, c.Array[index]) > 0 {
			min = c.Array[index]
		}
	}
	sli := make([]int, max-min+1)
	for index := 0; index < len(c.Array); index++ {
		sli[c.Array[index]-min]++
	}
	for index := 1; index < len(sli); index++ {
		sli[index] = sli[index] + sli[index-1]
	}
	newArray := make([]int, len(c.Array))
	for index := len(c.Array) - 1; index >= 0; index-- {
		newIndex := sli[c.Array[index]-min] - 1
		newArray[newIndex] = c.Array[index]
		sli[c.Array[index]-min]--
	}
	c.Array = newArray
}
