package graph

import (
	"fmt"
	"testing"
)

func TestGraph1_AddVertex(t *testing.T) {
	g := NewGraph1(5)
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	fmt.Println(g.VertexSize())
	fmt.Println(g.vertex)
}

func TestGraph1_AddEdge(t *testing.T) {
	g := NewGraph1(5)
	g.AddEdge("A", "B", 1)
	g.AddEdge("B", "C", 2)
	g.AddEdge("C", "D", 3)
	g.AddEdge("A", "D", 4)
	g.AddEdge("A", "C", 5)

	fmt.Println(g.vertex)
	g.Print()
	fmt.Println(g.EdgesSize())
	fmt.Println(g.VertexSize())
}
