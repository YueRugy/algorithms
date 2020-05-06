package rbt

import (
	"reflect"
)

const (
	unknown = -1
	left    = 0
	right   = 1
	black   = 0
	red     = 1
)

type RedBlackTree struct {
	size int
	root *Node
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

func (rbt *RedBlackTree) Remove(value int) {
	node := rbt.getNode(value)
	if node == nil {
		return
	}

	if node.left != nil && node.right != nil {
		predecessorNode := node.predecessor()
		node.value = predecessorNode.value
		node = predecessorNode
	}

	if node == rbt.root {
		rbt.root = nil
		rbt.size--
		return
	}
	var replaceNode *Node
	dire := node.dire()
	if node.left != nil {
		replaceNode = node.left
	} else if node.right != nil {
		replaceNode = node.right
	}
	if node.dire() == left {
		node.parent.left = replaceNode
	} else {
		node.parent.right = replaceNode
	}
	if replaceNode != nil {
		replaceNode.parent = node.parent
	}
	rbt.size--
	rbt.removeAfter(node, replaceNode, dire)
}

func (rbt *RedBlackTree) removeAfter(node, replaceNode *Node, dire int) {
	if rbt.isRed(node) {
		return
	}

	if rbt.isRed(replaceNode) {
		rbt.black(replaceNode)
		return
	}

	pNode := node.parent
	var siblingNode *Node
	if dire == left {
		siblingNode = pNode.right
	} else {
		siblingNode = pNode.left
	}
	if dire == left {
		if rbt.isRed(siblingNode) {
			rotateLeft(pNode, siblingNode)
			rbt.red(pNode)
			rbt.black(siblingNode)
			if siblingNode.parent == nil {
				rbt.root = siblingNode
			}
			siblingNode = pNode.right
		}
		if rbt.isBlack(siblingNode.left) && rbt.isBlack(siblingNode.right) {
			bFlag := rbt.isBlack(pNode)
			rbt.black(pNode)
			rbt.red(siblingNode)
			if bFlag {
				rbt.removeAfter(pNode, nil, pNode.dire())
			}
		} else {
			if rbt.isBlack(siblingNode.right) {
				rotateRight(siblingNode, siblingNode.left)
				siblingNode = pNode.right
			}
			rotateLeft(pNode, siblingNode)
			if siblingNode.parent == nil {
				rbt.root = siblingNode
			}
			rbt.color(siblingNode, pNode.color)
			rbt.black(pNode)
			rbt.black(siblingNode.right)
		}

	} else {
		if rbt.isRed(siblingNode) {
			rotateRight(pNode, siblingNode)
			if siblingNode.parent == nil {
				rbt.root = siblingNode
			}
			rbt.black(siblingNode)
			rbt.red(pNode)
			siblingNode = pNode.left
		}

		if siblingNode.left == nil && siblingNode.right == nil {
			bFlag := rbt.isBlack(pNode)
			rbt.black(pNode)
			rbt.red(siblingNode)
			if bFlag {
				rbt.removeAfter(pNode, nil, pNode.dire())
			}
		} else {
			if rbt.isBlack(siblingNode.left) {
				rotateLeft(siblingNode, siblingNode.right)
				siblingNode = pNode.left
			}
			rotateRight(pNode, siblingNode)
			if siblingNode.parent == nil {
				rbt.root = siblingNode
			}
			rbt.color(siblingNode, pNode.color)
			rbt.black(pNode)
			rbt.black(siblingNode.left)
		}
	}

}

func (rbt *RedBlackTree) Add(value int) {
	node := &Node{
		value: value,
	}

	if rbt.root == nil {
		rbt.root = node
		rbt.size++
		rbt.addAfter(rbt.root)
		return
	}
	res := rbt.root
	for temp := rbt.root; temp != nil; {
		res = temp
		if compare(value, temp.value) > 0 {
			temp = temp.right
		} else if compare(value, temp.value) < 0 {
			temp = temp.left
		} else {
			break
		}
	}

	if compare(value, res.value) > 0 {
		res.right = node
	} else if compare(value, res.value) < 0 {
		res.left = node
	} else {
		res.value = node.value
		return
	}

	node.parent = res
	rbt.addAfter(node)
	rbt.size++
}

func (rbt *RedBlackTree) addAfter(node *Node) {
	if node == rbt.root {
		rbt.black(node)
		return
	}

	rbt.red(node)
	if rbt.isBlack(node.parent) {
		return
	}
	//double red
	parent := node.parent
	uncle := parent.sibling()
	grand := parent.parent
	if rbt.isRed(uncle) {
		rbt.black(uncle)
		rbt.black(parent)
		rbt.red(grand)
		rbt.addAfter(grand)
	} else {
		if parent.dire() == left {
			if node.dire() == right {
				rotateLeft(parent, node)
				parent = node
			}
			rotateRight(grand, parent)
		} else {
			if node.dire() == left {
				rotateRight(parent, node)
				parent = node
			}
			rotateLeft(grand, parent)
		}
		rbt.red(grand)
		rbt.black(parent)
		if parent.parent == nil {
			rbt.root = parent
		}
	}
}

func (rbt *RedBlackTree) Size() int {
	return rbt.size
}

func (rbt *RedBlackTree) Empty() bool {
	return rbt.size == 0
}

func (rbt *RedBlackTree) Clear() {
	rbt.size = 0
	rbt.root = nil
}

func (rbt *RedBlackTree) red(node *Node) {
	if node != nil {
		node.color = red
	}
}

func (rbt *RedBlackTree) black(node *Node) {
	if node != nil {
		node.color = black
	}
}

func (rbt *RedBlackTree) color(node *Node, color int) {
	node.color = color
}

func (rbt *RedBlackTree) isRed(node *Node) bool {
	return node != nil && node.color == red
}

func (rbt *RedBlackTree) isBlack(node *Node) bool {
	return node == nil || node.color == black
}

func (rbt *RedBlackTree) getNode(value int) *Node {
	if rbt.root == nil {
		return nil
	}

	for temp := rbt.root; temp != nil; {
		if compare(value, temp.value) > 0 {
			temp = temp.right
		} else if compare(value, temp.value) < 0 {
			temp = temp.left
		} else {
			return temp
		}
	}

	return nil
}

func (rbt *RedBlackTree) LevelRange(visitor func(*Node)) {
	if rbt.root != nil {
		queue := make([]*Node, 0)
		queue = append(queue, rbt.root)
		for len(queue) > 0 {
			temp := queue[0]
			queue = queue[1:]
			visitor(temp)
			if temp.left != nil {
				queue = append(queue, temp.left)
			}
			if temp.right != nil {
				queue = append(queue, temp.right)
			}
		}
	}
}

type Node struct {
	value  int
	color  int
	left   *Node
	right  *Node
	parent *Node
}

func (node *Node) sibling() *Node {
	if node == nil || node.parent == nil {
		return nil
	}

	if node.dire() == left {
		return node.parent.right
	} else {
		return node.parent.left
	}
}

func (node *Node) predecessor() *Node {
	if node == nil {
		return nil
	}

	if node.left != nil {
		res := node.left
		for temp := node.left; temp != nil; {
			res = temp
			temp = temp.right
		}
		return res
	} else if node.parent != nil {
		for temp := node.parent; temp != nil; {
			if reflect.DeepEqual(temp, temp.parent.right) {
				return temp
			}
			temp = temp.parent
		}
		return nil
	}
	return nil
}

func (node *Node) dire() int {
	if node.parent == nil {
		return unknown
	}

	if reflect.DeepEqual(node, node.parent.left) {
		return left
	}
	return right
}

//------------------------------------------------
func compare(n1, n2 int) int {
	return n1 - n2
}

func rotateRight(p, n *Node) {
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

func rotateLeft(p, n *Node) {
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
