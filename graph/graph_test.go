package graph

import (
	"testing"
)

func TestGraph_AddVertex(t *testing.T) {
	g := NewGraph()
	g.AddVertex("A0")
	g.AddVertex("A1")
	g.AddVertex("A2")
	g.AddVertex("A3")
	g.AddVertex("A4")
	g.PrintVertices()
}
func TestGraph_AddEdge(t *testing.T) {
	g := NewGraph()
	g.AddEdge("V1", "V0", 9)
	g.AddEdge("V1", "V2", 3)
	g.AddEdge("V2", "V0", 2)
	g.AddEdge("V2", "V3", 5)
	g.AddEdge("V3", "V4", 1)
	g.AddEdge("V0", "V4", 6)
	g.AddEdge("V0", "V4", 6)
	g.PrintVertices()
	if g.size == 6 {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
func TestGraph_RemoveEdge(t *testing.T) {
	g := NewGraph()
	g.AddEdge("V1", "V0", 9)
	g.AddEdge("V1", "V2", 3)
	g.AddEdge("V2", "V0", 2)
	g.AddEdge("V2", "V3", 5)
	g.AddEdge("V3", "V4", 1)
	g.AddEdge("V0", "V4", 6)
	g.RemoveEdge("V0", "V4")
	if g.size == 5 {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
func TestGraph_RemoveVertex(t *testing.T) {
	g := NewGraph()
	g.AddEdge("V1", "V0", 9)
	g.AddEdge("V1", "V2", 3)
	g.AddEdge("V2", "V0", 2)
	g.AddEdge("V2", "V3", 5)
	g.AddEdge("V3", "V4", 1)
	g.AddEdge("V0", "V4", 6)
	g.RemoveVertex("V0")
	if g.size == 3 {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}

func TestGraph_BFS(t *testing.T) {
	g := NewGraph()
	//g.AddEdge("V1", "V0", 9)
	//g.AddEdge("V1", "V2", 3)
	//g.AddEdge("V2", "V0", 2)
	//g.AddEdge("V2", "V3", 5)
	//g.AddEdge("V3", "V4", 1)
	//g.AddEdge("V0", "V4", 6)
	//g.BFS("V1")
	foo(g)
	g.BFS("A")
}

func TestGraph_DFS(t *testing.T) {

	g := NewGraph()
	g.AddEdge("V1", "V0", 9)
	g.AddEdge("V1", "V2", 3)
	g.AddEdge("V2", "V0", 2)
	g.AddEdge("V2", "V3", 5)
	g.AddEdge("V3", "V4", 1)
	g.AddEdge("V0", "V4", 6)
	g.DFS("V1")
}

func foo1(g *Graph)  {
	var l=[]Temp{
		
{from:"0", to:"2",weight:0,},
{from:"1", to:"0",weight:0,},
{from:"2", to:"5",weight:0,},
{from:"2", to:"6",weight:0,},
{from:"3", to:"1",weight:0,}, 
{from:"3", to:"5",weight:0,},
{from:"3", to:"7",weight:0,}
{from:"5", to:"7",weight:0,},
{from:"6", to:"4",weight:0,},
{from:"7", to:"6",weight:0,},
	}
}

func foo(g *Graph) {
	var l = []Temp{
		{
			from:   "A",
			to:     "B",
			weight: 17,
		},
		{
			from:   "A",
			to:     "F",
			weight: 1,
		},
		{
			from:   "A",
			to:     "E",
			weight: 16,
		},
		{
			from:   "B",
			to:     "C",
			weight: 6,
		},
		{
			from:   "B",
			to:     "D",
			weight: 5,
		},
		{
			from:   "B",
			to:     "F",
			weight: 11,
		},
		{
			from:   "C",
			to:     "D",
			weight: 10,
		},
		{
			from:   "D",
			to:     "E",
			weight: 4,
		},
		{
			from:   "D",
			to:     "F",
			weight: 14,
		},
		{
			from:   "E",
			to:     "F",
			weight: 33,
		},
	}
	for _,v:=range l{
		g.AddEdge(v.from,v.to,v.weight)
	}
}

type Temp struct {
	from, to string
	weight   int
}
