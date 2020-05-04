package hash

import (
	"reflect"
)

const (
	unknow      = -1
	left        = 0
	right       = 1
	black       = 0
	red         = 1
	defaultCap  = 16
	balanceFact = 0.75
)

type Hash struct {
	buckets []*RedBlackTree
}

func NewHash() *Hash {
	return &Hash{buckets: make([]*RedBlackTree, defaultCap)}
}

type RedBlackTree struct {
	size int
	root *Node
}

func (rb *RedBlackTree) add(node *Node) {
	if rb.root == nil {
		rb.root = node
		rb.size++
	}
	// find pos add
	res := rb.root
	for temp := rb.root; temp != nil; {
		res = temp
		if compare(node.k, temp.k) > 0 {
			temp = temp.right
		} else if compare(node.k, temp.k) < 0 {
			temp = temp.left
		} else {
			break
		}
	}

	if compare(node.k, res.k) > 0 {
		res.right = node
		node.parent = res
		rb.size++
	} else if compare(node.k, res.k) < 0 {
		res.left = node
		node.parent = res
		rb.size++
	} else {
		res.v = node.v
		return
	}

	rb.addAfter(node)
}

func (rb *RedBlackTree) addAfter(node *Node) {
	if node == rb.root {
		node.color = black
		return
	}
	node.color = red
	if rb.isBlack(node.parent) {
		return
	}

	p := node.parent
	uncle := p.sibling()
	g := p.parent
	if rb.isRed(uncle) {
		rb.black(uncle)
		rb.black(p)
		rb.addAfter(p)
	} else {
		if p.dire() == left {
			if node.dire() == right {
				rotationLeft(node, p)
				p = node
			}
			rotationRight(p, g)
		} else {
			if node.dire() == left {
				rotationRight(node, p)
				p = node
			}
			rotationLeft(p, g)
		}
		if p.parent == nil {
			rb.root = p
		}
		rb.black(p)
		rb.red(g)
	}

}

func (rb *RedBlackTree) remove(k int) {
	if rb.root == nil {
		return
	}
	node := rb.getNode(k)
	if node == nil {
		return
	}

	if node.left != nil && node.right != nil {
		predecessorNode := node.predecessor()
		node.v = predecessorNode.v
		node.k = predecessorNode.k
		node = predecessorNode
	}

	dire := node.dire()
	var childNode *Node = nil
	if node.left != nil {
		if dire == left {
			node.parent.left = node.left
		} else if dire == right {
			node.parent.right = node.left
		} else {
			rb.root = node.left
		}
		node.left.parent = node.parent
		childNode = node.left
	} else if node.right != nil {
		if dire == left {
			node.parent.left = node.right
		} else if dire == right {
			node.parent.right = node.right
		} else {
			rb.root = node.right
		}
		node.right.parent = node.parent
		childNode = node.right
	} else {

		if dire == left {
			node.parent.left = nil
		} else if dire == right {
			node.parent.right = nil
		} else {
			rb.root = nil
		}
	}
	rb.size--
	rb.removeAfter(node, childNode, dire)

}

func (rb *RedBlackTree) removeAfter(node, childNode *Node, dire int) {
	if rb.root == nil || rb.isRed(node) {
		return
	}

	if rb.isRed(childNode) || childNode == rb.root {
		rb.black(childNode)
		return
	}

	p := node.parent
	var sibling *Node
	if dire == left {
		sibling = p.right
	} else {
		sibling = p.right
	}

	if dire == left {
		if rb.isRed(sibling) {
			rotationLeft(sibling, p)
			rb.black(sibling)
			rb.red(p)
			if sibling.parent == nil {
				rb.root = sibling
			}
			sibling = p.right
		}

		if sibling.left == nil && sibling.right == nil {
			rb.black(p)
			rb.red(sibling)
			rb.removeAfter(p, nil, p.dire())
		} else {
			if rb.isBlack(sibling.right) {
				rotationRight(sibling.left, sibling)
				sibling = p.right
			}
			rotationLeft(sibling, p)
			if sibling.parent == nil {
				rb.root = sibling
			}

			rb.colorOf(sibling, p.color)
			rb.black(p)
			rb.black(sibling.left)
		}
	} else {
		if rb.isRed(sibling) {
			rotationRight(sibling, p)
			rb.red(p)
			rb.black(sibling)
			if sibling.parent == nil {
				rb.root = sibling
			}
			sibling = p.left
		}
		if sibling.left == nil && sibling.right == nil {
			rb.black(p)
			rb.red(sibling)
			rb.removeAfter(p, nil, p.dire())
		} else {
			if rb.isBlack(sibling.left) {
				rotationLeft(sibling.right, sibling)
				sibling = p.left
			}
			rotationRight(sibling, p)
			if sibling.parent == nil {
				rb.root = sibling
			}

			rb.colorOf(sibling, p.color)
			rb.black(p)
			rb.black(p.right)
		}

	}

}

func (rb *RedBlackTree) getNode(k int) *Node {
	if rb.root == nil {
		return nil
	}

	for temp := rb.root; temp != nil; {
		if compare(k, temp.k) < 0 {
			temp = temp.left
		} else if compare(k, temp.k) > 0 {
			temp = temp.right
		} else {
			return temp
		}
	}

	return nil
}

func (rb *RedBlackTree) Size() int {
	return rb.size
}

func (rb *RedBlackTree) Clear() {
	rb.size = 0
	rb.root = nil
}

func (rb *RedBlackTree) Empty() bool {
	return rb.size == 0
}

func (rb *RedBlackTree) black(node *Node) {
	if node != nil {
		node.color = black
	}
}

func (rb *RedBlackTree) red(node *Node) {
	if node != nil {
		node.color = red
	}
}

func (rb *RedBlackTree) isBlack(node *Node) bool {
	return node == nil || node.color == black
}

func (rb *RedBlackTree) isRed(node *Node) bool {
	return node != nil && node.color == red
}

func (rb *RedBlackTree) colorOf(node *Node, color int) {
	if node == nil {
		return
	}
	node.color = color
}

type Node struct {
	k, v   int
	color  int
	parent *Node
	left   *Node
	right  *Node
}

func (n *Node) sibling() *Node {
	if n == nil || n.parent == nil {
		return nil
	}

	if n.dire() == left {
		return n.parent.right
	}
	return n.parent.left

}

func (n *Node) dire() int {
	p := n.parent
	if p == nil {
		return unknow
	}
	if reflect.DeepEqual(p.left, n) {
		return left
	} else {
		return right
	}
}

func (n *Node) predecessor() *Node {
	if n == nil {
		return nil
	}
	if n.left != nil {
		res := n.left
		for temp := res; temp != nil; {
			res = temp
			temp = temp.right
		}
		return res
	}

	if n.parent != nil {
		temp := n
		for temp.parent != nil {
			if temp.dire() == left {
				return temp.parent
			}
			temp = temp.parent
		}
	}

	return nil
}

//-----------------------------
func compare(k1, k2 int) int {
	return k1 - k2
}

func rotationLeft(n, p *Node) {
	p.right = n.left
	if n.left != nil {
		n.left.parent = p
	}
	n.left = p

	n.parent = p.parent
	if p.parent != nil {
		if p.dire() == left {
			p.parent.left = n
		} else {
			p.parent.right = n
		}
	}
	p.parent = n
}

func rotationRight(n, p *Node) {
	p.left = n.right
	if n.right != nil {
		n.right.parent = p
	}
	n.right = p

	n.parent = p.parent
	if p.parent != nil {
		if p.dire() == left {
			p.parent.left = n
		} else {
			p.parent.right = n
		}
	}

	p.parent = n
}
