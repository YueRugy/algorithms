package graph

type Graph2 struct {
	*Graph1
}

func NewGraph2(num int) *Graph2 {
	return &Graph2{
		Graph1: NewGraph1(num),
	}
}

func (g *Graph2) AddEdge(from, to string, weight int) {
	fromIndex := g.getVertexIndex(from)
	toIndex := g.getVertexIndex(to)
	g.edges[fromIndex][toIndex] = weight
	g.edgesNum++
}
