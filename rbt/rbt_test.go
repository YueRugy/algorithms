package rbt

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRedBlackTree_Add(t *testing.T) {
	rbt := NewRedBlackTree()
	arr := []int{94, 28, 70, 86, 89, 72, 24, 7, 75, 33, 23, 9, 55, 22, 80, 30, 18}
	for _, v := range arr {
		rbt.Add(v)
	}
	rbt.LevelRange(visitor)
	fmt.Println()
	rbt.LevelRange(visitor1)
}

func TestRedBlackTree_Remove(t *testing.T) {
	rbt := NewRedBlackTree()
	arr := []int{94, 28, 70, 86, 89, 72, 24, 7, 75, 33, 23, 9, 55, 22, 80, 30, 18}
	for _, v := range arr {
		rbt.Add(v)
	}

	for _, v := range arr {
		rbt.Remove(v)
		rbt.LevelRange(visitor)
		fmt.Println()
		rbt.LevelRange(visitor1)
		fmt.Println()
	}

}

func visitor(node *Node) {
	fmt.Print(strconv.Itoa(node.value) + "\t")
}

func visitor1(node *Node) {
	if node.color == red {
		fmt.Print(strconv.Itoa(node.value) + "\t")
	}
}
