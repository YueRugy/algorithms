package graph

type GraphSet struct {
	buckets map[Key]*Edge
}

func NewGraphSet() *GraphSet {
	return &GraphSet{
		buckets: make(map[Key]*Edge, 10),
	}
}

func (g *GraphSet) Add(k Key, e *Edge) {
	g.buckets[k] = e
}

func (g *GraphSet) Values() []*Edge {
	res := make([]*Edge, len(g.buckets))
	index := 0
	for _, v := range g.buckets {
		res[index] = v
		index++
	}
	return res
}
