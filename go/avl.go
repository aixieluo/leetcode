package main

import (
	"fmt"
	"log"
)

func main() {
	t := newAvlTree(1)
	t=t.Insert(2)
	t=t.Insert(3)
	t=t.Insert(4)
	t=t.Insert(5)
	t=t.Insert(6)
	t=t.Insert(7)
	fmt.Println(t)
}

type AvlTree struct {
	Key         int
	High        int
	Left, Right *AvlTree
}

func newAvlTree(key int) *AvlTree {
	return &AvlTree{Key: key}
}

func (node *AvlTree) getHeight() int  {
	if node == nil {
		return 0
	}
	return node.High
}

func (node *AvlTree) leftRotate() *AvlTree {
	right := node.Right
	node.Right = right.Left
	right.Left = node
	return right
}

func (node *AvlTree) rightRotate() *AvlTree {
	left := node.Left
	node.Left = left.Left
	left.Left = node
	return left
}

func (node *AvlTree) rightLeftRotate() *AvlTree {
	node.Left = node.Left.rightRotate()
	return node.leftRotate()
}

func (node *AvlTree) leftRightRotate() *AvlTree {
	node.Right = node.Right.leftRotate()
	return node.rightRotate()
}

func (node *AvlTree) adjust() *AvlTree {
	if node.Left.getHeight() - node.Right.getHeight() == 2 {
		if node.Left.Left.getHeight() > node.Left.Right.getHeight()  {
			node = node.rightRotate()
		} else {
			node = node.leftRightRotate()
		}
	} else if node.Right.getHeight() - node.Left.getHeight() ==2 {
		if node.Right.Right.getHeight() > node.Right.Left.getHeight(){
			node = node.leftRotate()
		} else {
			node = node.rightLeftRotate()
		}
	}
	return node
}

func (node *AvlTree) Insert(key int) *AvlTree {
	if node == nil {
		return &AvlTree{Key: key, High: 1}
	}
	if key < node.Key {
		node.Left = node.Left.Insert(key)
	} else if key > node.Key {
		node.Right = node.Right.Insert(key)
	} else {
		log.Fatal("node already exists!")
	}
	node = node.adjust()
	node.High = max(node.Left.getHeight(), node.Right.getHeight())+1
	return node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
