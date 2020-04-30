package quick

import (
	"math/rand"
	"time"

	"github.com/algorithms/soft"
)

type Quick struct {
	*soft.Soft
}

func init() {
	rand.Seed(time.Now().Unix())
}

func NewQuick(arr []int) *Quick {
	m := &Quick{
		Soft: &soft.Soft{
			Array: arr,
		},
	}
	return m
}

func (q *Quick) SortTest() {
	before := time.Now()

	q.Sort(0, len(q.Array))
	after := time.Now()
	q.String(after.Sub(before).String())
}

func (q *Quick) Sort(begin, end int) {
	if end-begin < 2 {
		return
	}

	pivotIndex := q.pivotIndex(begin, end)
	q.Sort(begin, pivotIndex)
	q.Sort(pivotIndex+1, end)
}

func (q *Quick) pivotIndex(begin, end int) int {
	end--
	randIndex := rand.Intn(end-begin) + begin
	q.Swap(randIndex, begin)

	node := q.Array[begin]
	for begin < end {
		for begin < end {
			if node < q.Array[end] {
				end--
			} else {
				q.Array[begin] = q.Array[end]
				begin++
				break
			}
		}

		for begin < end {
			if node > q.Array[begin] {
				begin++
			} else {
				q.Array[end] = q.Array[begin]
				end--
				break
			}
		}
	}
	q.Array[begin] = node
	return begin
}
