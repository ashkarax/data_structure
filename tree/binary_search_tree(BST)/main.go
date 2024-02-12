package main

import "fmt"

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

func InOrder(node *Node) {
	if node == nil {
		return
	}
	InOrder(node.left)
	fmt.Printf("%d  ", node.data)
	InOrder(node.right)
}

func (bst *BST) Delete(data int) *Node {
	bst.root = bst.deleteNode(bst.root, data)
	return bst.root
}

func (bst *BST) deleteNode(root *Node, data int) *Node {
	if root == nil {
		return nil
	}

	if data < root.data {
		root.left = bst.deleteNode(root.left, data)
	} else if data > root.data {
		root.right = bst.deleteNode(root.right, data)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		successor := bst.minimumValueNode(root.right)
		root.data = successor.data
		root.right = bst.deleteNode(root.right, successor.data)
	}
	return root
}

func (bst *BST) minimumValueNode(node *Node) *Node {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func main() {
	new := &BST{}
	new.Insert(11)
	new.Insert(4)
	new.Insert(12)
	new.Insert(5)
	new.Insert(13)
	new.Insert(6)
	new.Insert(14)
	new.Insert(7)
	new.Insert(15)
	new.Insert(0)
	new.Insert(99)
	new.Insert(5)
	new.Insert(8)
	InOrder(new.root)
	fmt.Println("======================================")
	new.Delete(15)
	fmt.Println("======================================")

	InOrder(new.root)

}
