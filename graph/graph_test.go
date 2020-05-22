package graph

import (
	"fmt"
	"testing"
)

func TestGraph_Dijkstra(t *testing.T) {
	g := NewGraph()
	foo4(g)
	res := g.Dijkstra("A")
	parse(res)
	fmt.Println(res)
}

func TestGraph_MstPrim(t *testing.T) {
	g := NewGraph()
	//foo2(g)
	foo3(g)
	//list := g.MstPrim()
	list := g.MstKruskal()
	showEdge(list)
}

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
	//g.Dijkstra("A")
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

func TestGraph_TopologicalSort(t *testing.T) {
	g := NewGraph()
	foo1(g)
	list := g.TopologicalSort()
	if list == nil {
		t.Log("不是有向无环图")
	}
	for _, str := range list {
		fmt.Printf("%s ", str)
	}
	t.Log("success")
}

func foo1(g *Graph) {
	var l = []Temp{
		{from: "0", to: "2", weight: 0},
		{from: "1", to: "0", weight: 0},
		{from: "2", to: "5", weight: 0},
		{from: "2", to: "6", weight: 0},
		{from: "3", to: "1", weight: 0},
		{from: "3", to: "5", weight: 0},
		{from: "3", to: "7", weight: 0},
		{from: "5", to: "7", weight: 0},
		{from: "6", to: "4", weight: 0},
		{from: "7", to: "6", weight: 0},
	}
	for _, v := range l {
		g.AddEdge(v.from, v.to, v.weight)
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
	for _, v := range l {
		g.AddEdge(v.from, v.to, v.weight)
	}
}

func foo3(g *Graph) {
	var l = []Temp{
		{from: "0", to: "2", weight: 2},
		{from: "0", to: "4", weight: 7},
		{from: "1", to: "2", weight: 3},
		{from: "1", to: "5", weight: 1},
		{from: "1", to: "6", weight: 7},
		{from: "2", to: "4", weight: 4},
		{from: "2", to: "5", weight: 3},
		{from: "2", to: "6", weight: 6},
		{from: "3", to: "7", weight: 9},
		{from: "4", to: "6", weight: 8},
		{from: "5", to: "6", weight: 4},
		{from: "5", to: "7", weight: 5},
	}

	for _, v := range l {
		g.AddEdge(v.from, v.to, v.weight)
		g.AddEdge(v.to, v.from, v.weight)
	}
}

func foo2(g *Graph) {
	var l = []Temp{
		{from: "A", to: "B", weight: 17},
		{from: "A", to: "F", weight: 1},
		{from: "A", to: "E", weight: 16},
		{from: "B", to: "C", weight: 6},
		{from: "B", to: "D", weight: 5},
		{from: "B", to: "F", weight: 11},
		{from: "C", to: "D", weight: 10},
		{from: "D", to: "E", weight: 4},
		{from: "D", to: "F", weight: 14},
		{from: "E", to: "F", weight: 33},
	}
	for _, v := range l {
		g.AddEdge(v.from, v.to, v.weight)
		g.AddEdge(v.to, v.from, v.weight)
	}
}

type Temp struct {
	from, to string
	weight   int
}

func showEdge(list map[*Edge]struct{}) {
	if list == nil || len(list) == 0 {
		return
	}
	for k := range list {
		fmt.Printf("from=%s  to=%s  weight=%d \n", k.from.value, k.to.value, k.weight)
	}
}

func foo4(g *Graph) {
	var l = []Temp{
		{from: "A", to: "B", weight: 10},
		{from: "A", to: "D", weight: 30},
		{from: "A", to: "E", weight: 100},
		{from: "B", to: "C", weight: 50},
		{from: "C", to: "E", weight: 10},
		{from: "D", to: "C", weight: 20},
		{from: "D", to: "E", weight: 60},
	}

	for _, v := range l {
		g.AddEdge(v.from, v.to, v.weight)
	}
}

func parse(result map[*vertex]*ValueInfo)  {
	for _, v := range result {
		fmt.Printf("终点是%s  最短路径=%d  ",v.key.value,v.distance)
		for _, paths := range v.paths {
			fmt.Printf("  %s---%s  ",paths.from.value,paths.to.value)
		}
		fmt.Println()
	}
}
