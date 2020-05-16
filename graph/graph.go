package graph

import "fmt"

type Graph struct {
	size     int
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
	if oe := fv.outEdges.buckets[k]; oe != nil {
		oe.weight = weight
		return
	}
	e := &Edge{
		weight: weight,
		from:   fv,
		to:     tv,
	}
	fv.outEdges.Add(k, e)
	tv.inEdges.Add(k, e)
	g.size++
}

func (g *Graph) RemoveEdge(from, to string) {
	fv := g.vertices[from]
	if fv == nil {
		return
	}
	tv := g.vertices[to]
	if tv == nil {
		return
	}
	k := Key{
		k1: from,
		k2: to,
	}
	delete(fv.outEdges.buckets, k)
	delete(tv.inEdges.buckets, k)
	g.size--
}

func (g *Graph) RemoveVertex(value string) {
	v := g.vertices[value]
	if v == nil {
		return
	}

	for _, e := range v.outEdges.buckets {
		k := Key{
			k1: value,
			k2: e.to.value,
		}
		delete(e.to.inEdges.buckets, k)
		g.size--
	}
	for _, e := range v.inEdges.buckets {
		k := Key{
			k1: e.from.value,
			k2: value,
		}
		delete(e.from.outEdges.buckets, k)
		g.size--
	}
	v.outEdges = nil
	v.inEdges = nil
	delete(g.vertices, value)
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
