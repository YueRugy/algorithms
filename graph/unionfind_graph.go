package graph

import "reflect"

type UnionfindGraph struct {
	ufg map[*vertex]*Node
}

type Node struct {
	rank   int
	value  *vertex
	parent *Node
}

func NewNode(value *vertex) *Node {
	node := &Node{
		value: value,
		rank:  1,
	}
	node.parent = node
	return node
}

func (uf *UnionfindGraph) Add(v *vertex) {
	if _, ok := uf.ufg[v]; ok {
		return
	}
	node := NewNode(v)
	uf.ufg[v] = node
}

func (uf *UnionfindGraph) Find(v *vertex) *vertex {
	node := uf.findNode(v)
	if node == nil {
		return nil
	}
	return node.value
}


func (uf *UnionfindGraph) Union(v1,v2 *vertex) {
	p1:=findNode(v1)
	p2:=findNode(v2)
	if p1||p2||reflect.DeepEqual(p1,p2){
		return
	}
	if p1.rank<p2.rank{
		p1.parent=p2
	}else if p1.rank>p2.rank{
		p2.parent=p1
	}else {
		p1.parent=p2
		p2.rank++
	}
}
func (uf *UnionfindGraph) IsSame(v1,v2 *vertex) bool{
	p1:=findNode(v1)
	p2:=findNode(v2)
	return reflect.DeepEqual(p1,p2)
}

func (uf *UnionfindGraph) findNode(v *vertex) *Node {
	if v == nil {
		return nil
	}

	node := uf.ufg[v]
	if node == nil {
		return nil
	}
	for !reflect.DeepEqual(node, node.parent) {
		node.parent = node.parent.parent
		node = node.parent
	}
	return node
}
