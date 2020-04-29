package merge

import (
	"github.com/algorithms/soft"
	"time"
)

type Merge struct {
	*soft.Soft
	leftArray []int
}

func NewMerge(arr []int) *Merge {
	m := &Merge{
		Soft: &soft.Soft{
			Array: arr,
		},
	}
	m.leftArray = make([]int, (len(m.Array)+1)>>1)
	return m
}
func (m *Merge) SortTest() {
	before := time.Now()

	m.Sort(0, len(m.Array)>>1, len(m.Array))
	after := time.Now()
	m.String(after.Sub(before).String())
}
func (m *Merge) Sort(begin, mid, end int) {
	if end-begin < 2 {
		return
	}
	m.Sort(begin, (begin+mid)>>1, mid)
	m.Sort(mid, (mid+end)>>1, end)
	m.Merge(begin, mid, end)
}

func (m *Merge) Merge(begin, mid, end int) {
	//拷贝left
	for i := 0; i < mid-begin; i++ {
		m.leftArray[i] = m.Array[begin+i]
	}
	ls, le := 0, mid-begin
	rs, re := mid, end
	ai := begin
	for ls < le {
		if rs >= re || m.CompareNum(m.leftArray[ls], m.Array[rs]) <= 0 {
			m.Array[ai] = m.leftArray[ls]
			ai++
			ls++
		} else {
			m.Array[ai] = m.Array[rs]
			rs++
			ai++
		}
	}
}
