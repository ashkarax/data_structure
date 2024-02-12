package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) pop() int {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

type Queue struct {
	inStack  Stack
	outStack Stack
}

func (q *Queue) enque(val int) {
	q.inStack.push(val)
}

func (q *Queue) deque() {
	if q.outStack.isEmpty() {
		for !q.inStack.isEmpty() {
			item := q.inStack.pop()
			q.outStack.push(item)
		}
	}
	item := q.outStack.pop()
	fmt.Println("popped item:", item)

}

func main() {

	q := Queue{}
	q.enque(5)
	q.enque(4)
	q.enque(3)
	q.enque(2)
	q.deque()
	q.deque()
	q.deque()

}
