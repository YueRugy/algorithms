package unionfind

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const max = 50000

func TestUnionFind_Union(t *testing.T) {
	uf := NewUnionFind(max)
	foo(uf)
}

func TestUnionFind1_Union(t *testing.T) {
	uf := NewUnionFind1(max)
	foo(uf)
}

func TestUnionFind_S_Union(t *testing.T) {
	uf:=NewUnionFind_S(max)
	foo(uf)
}

func TestUnionFind_R_Union(t *testing.T) {
	uf := NewUnionFind_R(max)
	foo(uf)
}

func TestUnionFind_C_Union(t *testing.T) {
	uf:=NewUnionFind_C(max)
	foo(uf)
}

func foo(uf Union) {
	rand.Seed(time.Now().Unix())
	uf.Union(0, 1)
	uf.Union(0, 3)
	uf.Union(0, 4)
	uf.Union(2, 3)
	uf.Union(2, 5)
	uf.Union(6, 7)
	uf.Union(8, 10)
	uf.Union(9, 10)
	uf.Union(9, 11)

	fmt.Println(uf.IsSame(2, 7))
	uf.Union(4, 6)
	fmt.Println(uf.IsSame(2, 7))
	fmt.Println(uf.Find(0))
	begin := time.Now()
	for index := 0; index < max; index++ {
		uf.Union(rand.Intn(max), rand.Intn(max))
	}
	for index := 0; index < max; index++ {
		uf.IsSame(rand.Intn(max), rand.Intn(max))
	}

	str := time.Now().Sub(begin).String()
	fmt.Println("use " + str)
}
