package unionfind

import (
	"fmt"
	"testing"
)

func TestUnionFind_Union(t *testing.T) {
	uf := NewUnionFind(12)
	uf.Union(0, 1)
	uf.Union(0, 3)
	uf.Union(0, 4)
	uf.Union(2, 3)
	uf.Union(2, 5)
	uf.Union(6, 7)
	uf.Union(8, 10)
	uf.Union(9, 10)
	uf.Union(9, 11)

	uf.Union(4, 6)
	fmt.Println(uf.IsSame(2, 7))
	fmt.Println(uf.parents)
}

func TestUnionFind1_Union(t *testing.T) {
	uf := NewUnionFind1(12)
	uf.Union(0, 1)
	uf.Union(0, 3)
	uf.Union(0, 4)
	uf.Union(2, 3)
	uf.Union(2, 5)
	uf.Union(6, 7)
	uf.Union(8, 10)
	uf.Union(9, 10)
	uf.Union(9, 11)

	uf.Union(4, 6)
	fmt.Println(uf.IsSame(2, 7))
	fmt.Println(uf.parents)
}
