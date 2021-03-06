package graph

import (
	"fmt"
	"github.com/algorithms/general"
	"math"
)

type Graph struct {
	size     int
	vertices map[string]*vertex
	edges    map[*Edge]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]*vertex, 20),
		edges:    make(map[*Edge]struct{}, defaultCap),
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

type ValueInfo struct {
	distance int
	key      *vertex
	paths    []*Edge
}

func NewValueInfo(distance int, key *vertex, paths []*Edge) *ValueInfo {
	return &ValueInfo{distance: distance, key: key, paths: paths}
}

type ValueInfoSlice []*ValueInfo

func (v ValueInfoSlice) Less(i, j int) bool {
	return v[i].distance < v[j].distance
}

func (v ValueInfoSlice) Swap(i, j int) {
	if i == j {
		return
	}
	v[i], v[j] = v[j], v[i]
}

func (v ValueInfoSlice) Length() int {
	return len(v)
}

func (v *ValueInfoSlice) Push(x interface{}) {
	*v = append(*v, x.(*ValueInfo))
}

func (v *ValueInfoSlice) Pop() interface{} {
	old := *v
	n := old.Length()
	value := old[n-1]
	*v = old[:n-1]
	return value
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
	g.edges[e] = struct{}{}
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

func (g *Graph) Dijkstra(k string) map[*vertex]*ValueInfo {

	ver, ok := g.vertices[k]
	if len(g.vertices) < 2 || !ok {
		return nil
	}
	heap := make(ValueInfoSlice, 0)
	dm := make(map[*vertex]*ValueInfo)
	if ver.outEdges != nil {
		for _, edge := range ver.outEdges.buckets {
			vi := NewValueInfo(edge.weight, edge.to, make([]*Edge, 0))
			vi.paths = append(vi.paths, edge)
			dm[edge.to] = vi
			general.Push(&heap, vi)
		}
		//fmt.Println(heap)
	} else {
		return nil
	}
	selectPath := make(map[*vertex]*ValueInfo)
	begin := NewValueInfo(math.MaxInt64, ver, make([]*Edge, 0))
	selectPath[ver] = begin
	for heap.Length() > 0 && len(selectPath) <= len(g.vertices) {
		v := general.Pop(&heap).(*ValueInfo)
		selectPath[v.key] = v
		if v.key.outEdges != nil {
			for _, e := range v.key.outEdges.buckets {
				if selectPath[e.to] != nil && selectPath[e.from] != nil {
					continue
				}
				info := dm[e.to]
				if info == nil {
					info = NewValueInfo(e.weight+v.distance, e.to, make([]*Edge, 0))
					for _, p := range v.paths {
						info.paths = append(info.paths, p)
					}
					info.paths = append(info.paths, e)
					general.Push(&heap, info)
					dm[e.to] = info
				} else {
					distance := v.distance + e.weight
					if distance >= info.distance {
						continue
					} else {
						info.distance = distance
						info.paths = make([]*Edge, 0)
						for _, p := range v.paths {
							info.paths = append(info.paths, p)
						}
						info.paths = append(info.paths, e)
						general.Init(&heap)
					}
				}
			}
		}
	}

	delete(selectPath, ver)
	return selectPath
}

func (g *Graph) BellmanFord(key string) map[*vertex]*ValueInfo {
	vert, ok := g.vertices[key]
	if len(g.vertices) < 2 || !ok {
		return nil
	}

	res := make(map[*vertex]*ValueInfo)
	if vert.outEdges == nil {
		return nil
	}

	for _, e := range vert.outEdges.buckets {
		vi := NewValueInfo(e.weight, e.to, make([]*Edge, 0))
		vi.paths = append(vi.paths, e)
		res[e.to] = vi
	}

	vertexSize := len(g.vertices)
	for index := 0; index < vertexSize; index++ {
		for edge := range g.edges {
			g.relax(res, edge)
		}
	}

	for edge := range g.edges {
		if g.relax(res,edge){
			fmt.Println("有负权环")
			return nil
		}
	}

	return res
}

func (g *Graph) relax(res map[*vertex]*ValueInfo, edge *Edge) bool {
	if res[edge.from] != nil {
		distance := res[edge.from].distance + edge.weight
		if res[edge.to] == nil {
			vi := NewValueInfo(distance, edge.to, make([]*Edge, 0))
			for _, ep := range res[edge.from].paths {
				vi.paths = append(vi.paths, ep)
			}
			vi.paths = append(vi.paths, edge)
			res[edge.to] = vi
			return true
		} else {
			if distance < res[edge.to].distance {
				res[edge.to].distance = distance
				res[edge.to].paths = make([]*Edge, 0)
				for _, ep := range res[edge.from].paths {
					res[edge.to].paths = append(res[edge.to].paths, ep)
				}
				res[edge.to].paths = append(res[edge.to].paths, edge)
				return true
			}
		}
	}
	return false
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

func (g *Graph) MstKruskal() map[*Edge]struct{} {
	if len(g.vertices) <= 1 {
		return nil
	}
	res := make(map[*Edge]struct{})
	uf := NewUnionfindGraph()
	for _, v := range g.vertices {
		uf.Add(v)
	}
	collection := make([]*Edge, len(g.edges))
	index := 0
	for v := range g.edges {
		collection[index] = v
		index++
	}
	mh := NewHeapGraph(collection, func(e1, e2 *Edge) int {
		return e1.weight - e2.weight
	})
	for !mh.IsEmpty() && len(res) < len(g.vertices)-1 {
		e := mh.Remove()
		if uf.IsSame(e.from, e.to) {
			continue
		}
		res[e] = struct{}{}
		uf.Union(e.from, e.to)
	}
	return res
}

func (g *Graph) MstPrim() map[*Edge]struct{} {
	if len(g.vertices) == 1 || len(g.vertices) == 0 {
		return nil
	}
	vertexSet := make(map[*vertex]struct{}, len(g.vertices))
	res := make(map[*Edge]struct{})
	var beginVertex *vertex
	for _, value := range g.vertices {
		beginVertex = value
		break
	}
	if beginVertex == nil {
		return nil
	}
	vertexSet[beginVertex] = struct{}{}
	heap := NewHeapGraph(beginVertex.outEdges.Values(), func(e1, e2 *Edge) int {
		return e1.weight - e2.weight
	})
	for heap.Size() > 0 && len(vertexSet) < len(g.vertices) {
		edge := heap.Remove()

		_, o := vertexSet[edge.from]
		_, o1 := vertexSet[edge.to]
		if o && o1 {
			continue
		}
		vertexSet[edge.to] = struct{}{}
		res[edge] = struct{}{}
		if edge.to.outEdges != nil {
			for _, v := range edge.to.outEdges.buckets {
				if _, ok := res[v]; ok {
					continue
				}
				_, ok := vertexSet[v.from]
				_, ok1 := vertexSet[v.to]
				if ok && ok1 {
					continue
				}
				heap.Add(v)
			}
		}
	}
	return res
	//heap:=NewHeapGraph(beginVertex.outEdges.buckets)
}

func (g *Graph) MstPrim1() map[*Edge]struct{} {
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
