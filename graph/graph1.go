package graph

import "fmt"

type Graph1 struct {
	edgesNum int
	vertex   []string
	edges    [][]int
}

func NewGraph1(num int) *Graph1 {
	g := &Graph1{}
	g.edges = make([][]int, num)
	for index := 0; index < len(g.edges); index++ {
		g.edges[index] = make([]int, num)
	}
	return g
}

func (g *Graph1) EdgesSize() int {
	return g.edgesNum
}

func (g *Graph1) VertexSize() int {
	return len(g.vertex)
}

func (g *Graph1) AddVertex(v string) {
	g.vertex = append(g.vertex, v)
}

func (g *Graph1) add(from, to, weight int) {
	g.edges[from][to] = weight
	g.edges[to][from] = weight
	g.edgesNum++
}

func (g *Graph1) AddEdge(from, to string, weight int) {
	fromIndex := g.getVertexIndex(from)
	toIndex := g.getVertexIndex(to)
	g.add(fromIndex, toIndex, weight)
}

func (g *Graph1) Print() {
	for index := 0; index < len(g.edges); index++ {
		fmt.Println(g.edges[index])
	}
}

func (g *Graph1) getVertexIndex(str string) int {
	for index := 0; index < len(g.vertex); index++ {
		if g.vertex[index] == str {
			return index
		}
	}
	g.vertex = append(g.vertex, str)
	return len(g.vertex) - 1
}
