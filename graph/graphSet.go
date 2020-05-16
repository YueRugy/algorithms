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
