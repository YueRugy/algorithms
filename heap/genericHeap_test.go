package heap

import (
	"fmt"
	"testing"
)

type arr []*interface{}
type Student struct {
	id int
}

func NewStudent(id int) *Student {
	return &Student{id: id}
}

func (s Student) Compare(v1 Comparable) int {
	if s, ok := v1.(*Student); ok {
		return s.getId() - s.getId()
	}
	panic("unexpect type")
}

func (s Student) getId() int {
	return s.id
}

func TestGenericHeap_Add(t *testing.T) {
	var list = make([]*Student, 10)
	for index := 0; index < len(list); index++ {
		list[index] = NewStudent(index)
	}
	heap := NewGenericHeap(list, compare)
	heap.Remove()
	fmt.Println(heap)
}

func compare(v1, v2 interface{}) int {
	s1, ok := v1.(*Student)
	s2, ok1 := v2.(*Student)
	if !ok || !ok1 {
		panic("unexpect type v2")
	}
	return s1.id - s2.id
}
