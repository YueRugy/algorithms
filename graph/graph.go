package graph

import "fmt"

type Graph struct {
	vertices map[string]*vertex
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]*vertex, 20),
	}
}

type Key struct {
	k1, k2 string
}

type vertex struct {
	value    string
	inEdges  *GraphSet
	outEdges *GraphSet
}

type Edge struct {
	weight int
	from   *vertex
	to     *vertex
}

func (g *Graph) AddEdge(from, to string, weight int) {
	fv := g.AddVertex(from)
	tv := g.AddVertex(to)
	e := &Edge{
		weight: weight,
		from:   fv,
		to:     tv,
	}
	if fv.outEdges == nil {
		fv.outEdges = NewGraphSet()
	}
	if tv.inEdges == nil {
		tv.inEdges = NewGraphSet()
	}
	k := Key{
		k1: from,
		k2: to,
	}
	fv.outEdges.Add(k,e)
	tv.inEdges.Add(k,e)
}

func (g *Graph) AddVertex(value string) *vertex {
	if _, ok := g.vertices[value]; !ok {
		g.vertices[value] = &vertex{
			value: value,
		}
	}
	return g.vertices[value]
}

func (g *Graph) PrintVertices() {
	for k, v := range g.vertices {
		fmt.Print(k)
		fmt.Printf("   %s", v.value)
		fmt.Println()
	}
}
