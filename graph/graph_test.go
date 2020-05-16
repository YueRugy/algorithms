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
	if g.size==6{
		t.Log("success")
	}else {
		t.Error("failed")
	}
}
func TestGraph_RemoveEdge(t *testing.T) {
	g:=NewGraph()
	g.AddEdge("V1","V0",9)
	g.AddEdge("V1","V2",3)
	g.AddEdge("V2","V0",2)
	g.AddEdge("V2","V3",5)
	g.AddEdge("V3","V4",1)
	g.AddEdge("V0","V4",6)
	g.RemoveEdge("V0","V4")
	if g.size==5{
		t.Log("success")
	}else {
		t.Error("failed")
	}
}
func TestGraph_RemoveVertex(t *testing.T) {
	g:=NewGraph()
	g.AddEdge("V1","V0",9)
	g.AddEdge("V1","V2",3)
	g.AddEdge("V2","V0",2)
	g.AddEdge("V2","V3",5)
	g.AddEdge("V3","V4",1)
	g.AddEdge("V0","V4",6)
	g.RemoveVertex("V0")
	if g.size==3{
		t.Log("success")
	}else {
		t.Error("failed")
	}
}
