package main

import "fmt"

type Queue struct{
	arr []int
}

func (q *Queue)enque(data int){
q.arr=append(q.arr, data)
}

func (q *Queue) Display(){
	for _,item:=range q.arr{
		fmt.Println(item)
	}
}

func (q *Queue)deque(){
	if len(q.arr)==0{
		fmt.Println("invalid")
		return
	}
	item:=q.arr[0]
	q.arr=q.arr[1:]
	fmt.Println("remove item is :",item)

}

func main(){
	var q Queue
	q.enque(3)
	q.enque(4)
	q.enque(5)
	q.enque(6)
	q.Display()
	q.deque()
	q.deque()
	q.enque(111)

	q.Display()



}