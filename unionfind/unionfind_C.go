package unionfind

type UnionFind_C struct {
	parents []int
	rank    []int
}

func NewUnionFind_C(capacity int) *UnionFind_C {
	uf := &UnionFind_C{
		parents: make([]int, capacity),
		rank:    make([]int, capacity),
	}

	for index := 0; index < len(uf.parents); index++ {
		uf.parents[index] = index
		uf.rank[index] = 1
	}
	return uf
}

func (uf *UnionFind_C) Union(v1, v2 int) {
	if uf.checkV(v1) || uf.checkV(v2) || v1 == v2 {
		return
	}

	p1 := uf.Find(v1)
	p2 := uf.Find(v2)
	if p1 == p2 {
		return
	}

	if uf.rank[p1] > uf.rank[p2] {
		uf.parents[p2] = p1
	} else if uf.rank[p1] < uf.rank[p2] {
		uf.parents[p1] = p2
	} else {
		uf.parents[p1] = p2
		uf.rank[p2]++
	}
}

func (uf *UnionFind_C) Find(v int) int {
	if uf.checkV(v) {
		return -1
	}

	if v != uf.parents[v] {
		uf.parents[v] = uf.Find(uf.parents[v])
	}
	return uf.parents[v]
}

func (uf *UnionFind_C) checkV(v int) bool {
	if v >= len(uf.parents) || v < 0 {
		return true
	}
	return false
}

func (uf *UnionFind_C) IsSame(v1, v2 int) bool {
	if uf.checkV(v1) || uf.checkV(v2) {
		return false
	}
	return uf.Find(v1) == uf.Find(v2)
}
