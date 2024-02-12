package main

import (
	"fmt"
	"math"
)

type Node struct {
	data        int
	left, right *Node
}

type BST struct {
	root *Node
}

func NewNode(data int) *Node {
	return &Node{data: data, left: nil, right: nil}
}

func (bst *BST) Insert(data int) {
	if bst.root == nil {
		bst.root = NewNode(data)
	} else {
		InsertNode(bst.root, data)
	}
}

func InsertNode(node *Node, data int) {
	if data <= node.data {
		if node.left == nil {
			node.left = NewNode(data)
		} else {
			InsertNode(node.left, data)
		}
	} else {
		if node.right == nil {
			node.right = NewNode(data)
		} else {
			InsertNode(node.right, data)
		}
	}
}

func InOrder(root *Node) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Println("", root.data)
	InOrder(root.right)
}

func (bst *BST) Validate() bool {
	return bst.ValidateUtil(bst.root, math.MinInt, math.MaxInt)
}

func (bst *BST) ValidateUtil(node *Node, min, max int) bool {
	if node == nil {
		return true
	}
	if node.data <= min || node.data >= max {
		return false
	}

	return bst.ValidateUtil(node.left, min, node.data) && bst.ValidateUtil(node.right, node.data, max)
}
func main() {
	bst := BST{}
	bst.Insert(5)
	bst.Insert(1)
	bst.Insert(44)
	bst.Insert(22)
	bst.Insert(3)
	bst.Insert(0)
	InOrder(bst.root)

	fmt.Println("tree bst stat:", bst.Validate())

}
