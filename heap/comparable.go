package heap

type Comparable interface {
	Compare(v1 Comparable) int
}

func foo(e *Comparable)  {
}

