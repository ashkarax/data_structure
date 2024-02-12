package main

import "fmt"

type Node struct {
	data int
	Next *Node
}

type Stack struct {
	head *Node
	tail *Node
}

func (s *Stack) push(data int) {

	if s.head == nil {
		newNode := &Node{data: data, Next: nil}
		s.head = newNode
		s.tail = newNode
		return
	}
	newNode := &Node{data: data, Next: s.head}
	s.head = newNode
}

func (s *Stack) Display() {
	current := s.head
	for current != nil {
		fmt.Println(current.data)
		current = current.Next
	}
}

func (s *Stack) pop() int {
	item := s.head.data
	s.head = s.head.Next
	fmt.Println("popped one:", item)
	return item

}

func (s *Stack) removeParticular(val int) {
	current := s.head
	var temparr []int
	for current.Next != nil {
		if current.data != val {
			temparr = append(temparr, s.pop())
		} else {
			found := s.pop()
			fmt.Println(found)
			break
		}
		current = current.Next
	}
	fmt.Println("----------------88888888", temparr)
	if len(temparr) > 0 {
		for i := len(temparr) - 1; i >= 0; i-- {
			fmt.Println("--------------------", temparr[i])
			s.push(temparr[i])
		}
	}

}

func main() {
	var stack Stack
	// stack.push(8)

	// stack.push(7)
	stack.push(6)
	stack.push(5)
	stack.push(4)
	stack.push(3)
	stack.push(2)
	stack.push(1)
	stack.Display()
	fmt.Println("-----------------------")
	stack.removeParticular(4)
	fmt.Println("-----------------------")

	stack.Display()

}
