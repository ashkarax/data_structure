package main

import "fmt"

type Node struct{
	data int
	Next *Node
}

type stack struct{
head *Node
}

func (ll *stack) Push(data int){
	if ll.head == nil{
		newNode:=&Node{data: data,Next: nil}
		ll.head=newNode
		return
	}
	newNode:=&Node{data: data,Next: ll.head}
ll.head=newNode
}

func (ll *stack)Display(){
	current:=ll.head
	for current!=nil{
		fmt.Println(current.data)
current=current.Next
	}
}

func (ll *stack) Pop(){
	ll.head=ll.head.Next
}

func main(){
var sls stack
	
sls.Push(5)
sls.Push(4)
sls.Push(3)
sls.Display()
sls.Pop()
sls.Pop()
sls.Display()


}