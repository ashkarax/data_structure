package main

import "fmt"

type Node struct{
	data int
	Next *Node
}

type Queue struct{
	head *Node
}

func (q *Queue)enque(data int){
	newNode:=&Node{data: data,Next: nil}
	if q.head==nil{
		q.head=newNode
		return
	}
current:=q.head
for current.Next!= nil{
	current=current.Next
}
current.Next=newNode
}

func (q *Queue) Display(){
	current:=q.head
	for current!=nil{
		fmt.Println(current.data)
		current=current.Next
	}
}

func (q *Queue) deque(){
	q.head=q.head.Next
}

func main(){
	var q Queue
	q.enque(4)
	q.enque(5)
	q.enque(6)
	q.enque(7)
	q.Display()
	q.deque()
	q.Display()
}