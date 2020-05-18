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

func (g *Graph) BFS(ver string) {
	v, ok := g.vertices[ver]
	if !ok {
		return
	}
	queue := make([]*vertex, 0)
	queue = append(queue, v)
	set := make(map[*vertex]struct{})
	set[v] = struct{}{}
	for len(queue) != 0 {
		vert := queue[0]
		fmt.Println(vert.value)
		queue = queue[1:]
		if vert.outEdges != nil {
			for _, value := range vert.outEdges.buckets {
				if _, ok := set[value.to]; !ok {
					set[value.to] = struct{}{}
					queue = append(queue, value.to)
				}
			}
		}
	}
}

func (g *Graph) DFS(ver string) {
	s := make(map[*vertex]struct{})
	vert, ok := g.vertices[ver]
	if ok {
		g.dfs(vert, s)
	}
}

func (g *Graph) dfs(vert *vertex, s map[*vertex]struct{}) {
	s[vert] = struct{}{}
	fmt.Println(vert.value)
	if vert.outEdges == nil {
		return
	}
	for _, value := range vert.outEdges.buckets {
		if _, ok := s[value.to]; !ok {
			s[value.to] = struct{}{}
			g.dfs(value.to, s)
		}
	}
}

func (g *Graph) PrintVertices() {
	for k, v := range g.vertices {
		fmt.Print(k)
		fmt.Printf("   %s", v.value)
		fmt.Println()
	}
}

func (g *Graph) TopologicalSort() []string {
	mv := make(map[*vertex]int, len(g.vertices))
	queue := make([]*vertex, 0)
	for _, value := range g.vertices { // init mv
		if value.inEdges == nil {
			queue = append(queue, value)
		} else {
			mv[value] = len(value.inEdges.buckets)
		}
	}

	if len(queue) == 0 {
		return nil
	}

	var res = make([]string, len(g.vertices))
	var index = 0
	for len(queue) > 0 {
		temp := queue[0]
		res[index] = temp.value
		index++
		queue = queue[1:]
		if temp.outEdges != nil {
			for _, v := range temp.outEdges.buckets {
				mv[v.to]--
				if mv[v.to] == 0 {
					queue = append(queue, v.to)
				}
			}
		}
	}
	if len(res) != len(g.vertices) {
		return nil
	}
	return res
}

func (g *Graph) MstPrim() map[*Edge]struct{} {
	length := len(g.vertices)
	if length == 0 {
		return nil
	}

	es := make(map[*Edge]struct{}, length-1)
	vs := make(map[*vertex]struct{}, length)
	for _, value := range g.vertices {
		vs[value] = struct{}{}
		break
	}

	for len(vs) < length {
		var wl *Edge = nil
		for value := range vs {
			if value.outEdges.buckets != nil {
				for _, eg := range value.outEdges.buckets {
					_, ok := vs[eg.from]
					_, ok1 := vs[eg.to]
					if ok && ok1 {
						continue
					}
					if wl == nil {
						wl = eg
					} else {
						if wl.weight > eg.weight {
							wl = eg
						}
					}
				}
			}
		}
		if wl != nil {
			vs[wl.to] = struct{}{}
			es[wl] = struct{}{}
		}
	}

	return es
}
