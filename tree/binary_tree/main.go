package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func main() {
	node1 := &Node{data: 11, right: nil, left: nil}
	node2 := &Node{data: 12, right: nil, left: nil}
	node3 := &Node{data: 13, right: nil, left: nil}
	node4 := &Node{data: 14, right: nil, left: nil}
	node5 := &Node{data: 15, right: nil, left: nil}
	node6 := &Node{data: 16, right: nil, left: nil}
	node7 := &Node{data: 17, right: nil, left: nil}

	bt := &BinaryTree{}

	bt.root = node1
	bt.root.left = node2
	bt.root.right = node3

	bt.root.left.left = node4
	bt.root.left.right = node5

	bt.root.right.left = node6
	bt.root.right.right = node7

	fmt.Print("Pre-Order Traversal :-")
	PreOrder(bt.root)
	fmt.Println()

	fmt.Print("InOrder Traversal :-")
	InOrder(bt.root)
	fmt.Println()

	fmt.Print("PostOrder Traversal :-")
	PostOrder(bt.root)
	fmt.Println()

}

func PreOrder(root *Node) {
	if root == nil {
		return
	}
	PreOrder(root.left)
	PreOrder(root.right)
	fmt.Print(root.data, " ")
}

func InOrder(root *Node) {
	if root == nil {
		return
	}
	PreOrder(root.left)
	fmt.Print(root.data, " ")
	PreOrder(root.right)
}

func PostOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.data, " ")
	PreOrder(root.left)
	PreOrder(root.right)
}
