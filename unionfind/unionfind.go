package unionfind

type UnionFind struct {
	parents []int
}

func NewUnionFind(capacity int) *UnionFind {
	uf := &UnionFind{
		parents: make([]int, capacity),
	}

	for index := 0; index < len(uf.parents); index++ {
		uf.parents[index] = index
	}
	return uf
}

func (u *UnionFind) Find(v int) int {
	if u.checkV(v) {
		return -1
	}

	return u.parents[v]
}

func (u *UnionFind) IsSame(v1, v2 int) bool {
	if u.checkV(v1) || u.checkV(v2) || v1 == v2 {
		return false
	}
	return u.Find(v1) == u.Find(v2)
}

func (u *UnionFind) checkV(v int) bool {
	if v >= len(u.parents) || v < 0 {
		return true
	}
	return false
}

func (u *UnionFind) Union(v1, v2 int) {
	if u.checkV(v1) || u.checkV(v2) {
		return
	}
	p1 := u.Find(v1)
	p2 := u.Find(v2)
	if p1 == p2 {
		return
	}
	for index := 0; index < len(u.parents); index++ {
		if u.parents[index] == p1 {
			u.parents[index] = p2
		}
	}
}
