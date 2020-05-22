package general

type Sort interface {
	Less(i, j int) bool
	Swap(i, j int)
	Length() int
	Push(x interface{})
	Pop() interface{}
}

func Init(sort Sort) {
	//heapify
	length := sort.Length()
	for i := length>>1 - 1; i >= 0; i-- {
		down(sort, i, length)
	}
}

func Push(sort Sort, ele interface{}) {
	sort.Push(ele)
	up(sort, sort.Length()-1)
}

func Pop(sort Sort) interface{} {
	length := sort.Length() - 1
	if length == 0 {
		return sort.Pop()
	}
	sort.Swap(0, length)
	down(sort, 0, length)
	return sort.Pop()
}

func Remove(sort Sort, index int) interface{} {
	l := sort.Length() - 1
	if l != index {
		sort.Swap(index, l)
		if !down(sort, index, l) {
			up(sort, index)
		}
	}
	return sort.Pop()
}

func up(sort Sort, i int) {
	for {
		parent := (i - 1) >> 1
		if parent <= 0 || !sort.Less(parent, i) {
			break
		}
		sort.Swap(parent, i)
		i = parent
	}
}

func down(sort Sort, i int, length int) bool {
	index := i
	for {
		left := index<<1 + 1
		if left >= length || left < 0 {
			break
		}

		finalIndex := left
		if right := left + 1; right < length && sort.Less(right, left) {
			finalIndex = right
		}
		if !sort.Less(finalIndex, index) {
			break
		}
		sort.Swap(index, finalIndex)
		index = finalIndex
	}
	return index > i
}
