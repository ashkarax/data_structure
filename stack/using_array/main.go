package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {

	if len(s.items) == 0 {
		fmt.Println("stack is empty")
		return
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	fmt.Println("removed item is:", item)
}
func (s *Stack) Display() {
	fmt.Println(s.items)
}
func main() {
	var stack Stack

	stack.Display()
	stack.Push(5)
	stack.Push(5)
	stack.Push(5)
	stack.Pop()
	stack.Display()
}
