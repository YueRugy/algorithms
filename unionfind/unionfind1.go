package unionfind

type UnionFind1 struct {
	parents []int
}

func NewUnionFind1(capacity int) *UnionFind1 {
	uf := &UnionFind1{
		parents: make([]int, capacity),
	}

	for index := 0; index < len(uf.parents); index++ {
		uf.parents[index] = index
	}
	return uf
}

func (uf *UnionFind1) Union(v1, v2 int) {
	if uf.checkV(v1) || uf.checkV(v2) || v1 == v2 {
		return
	}

	p1 := uf.Find(v1)
	p2 := uf.Find(v2)
	if p1 == p2 {
		return
	}
	uf.parents[p1] = p2
}

func (uf *UnionFind1) Find(v int) int {
	if uf.checkV(v) {
		return -1
	}

	for uf.parents[v] != v {
		v = uf.parents[v]
	}
	return v
}

func (uf *UnionFind1) IsSame(v1, v2 int) bool {
	if uf.checkV(v1) || uf.checkV(v2) {
		return false
	}
	return uf.Find(v1) == uf.Find(v2)
}

func (uf *UnionFind1) checkV(v int) bool {
	if v >= len(uf.parents) || v < 0 {
		return true
	}
	return false
}
