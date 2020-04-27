package mi

type Node struct {
	Index, Value int
}

type HeapI interface {
	Size() int
	Empty() bool
	Add(*Node)
	Get() *Node
	Clear()
	Remove() *Node
	Replace(*Node) *Node
}
