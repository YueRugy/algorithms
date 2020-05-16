package graph

import (
	"testing"
)

func TestGraph_AddVertex(t *testing.T) {
	g:=NewGraph()
	g.AddVertex("A0")
	g.AddVertex("A1")
	g.AddVertex("A2")
	g.AddVertex("A3")
	g.AddVertex("A4")
	g.PrintVertices()
}
func TestGraph_AddEdge(t *testing.T) {
	g:=NewGraph()
	g.AddEdge("V1","V0",9)
	g.AddEdge("V1","V2",3)
	g.AddEdge("V2","V0",2)
	g.AddEdge("V2","V3",5)
	g.AddEdge("V3","V4",1)
	g.AddEdge("V0","V4",6)
	g.AddEdge("V0","V4",6)
	g.PrintVertices()
}