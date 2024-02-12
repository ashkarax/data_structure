package main

import "fmt"

type MaxHeap struct {
	arr []int
}

func (mh *MaxHeap) Insert(data int) {
	mh.arr = append(mh.arr, data)
	mh.buildHeap()
}

func (mh *MaxHeap) buildHeap() {
	n := len(mh.arr)
	for i := n/2 - 1; i >= 0; i-- {
		mh.maxHeapify(i)
	}
}

func (mh *MaxHeap) maxHeapify(i int) {
	var k int
	n := len(mh.arr)
	for {
		l := 2*i + 1
		r := 2*i + 2
		if l >= n {
			break
		}
		if r < n && mh.arr[r] > mh.arr[l] {
			k = r
		} else {
			k = l
		}
		if mh.arr[k] > mh.arr[i] {
			mh.arr[k], mh.arr[i] = mh.arr[i], mh.arr[k]
			i = k
		} else {
			break
		}

	}
}

func (mh *MaxHeap) Delete() int {
	if len(mh.arr) == 0 {
		return -1
	}
	maxValue := mh.arr[0]
	mh.arr[0] = mh.arr[len(mh.arr)-1]
	mh.arr = mh.arr[:len(mh.arr)-1]
	mh.maxHeapify(0)
	return maxValue
}

func main() {
	var maxheap MaxHeap
	maxheap.Insert(99)
	maxheap.Insert(5)
	maxheap.Insert(4)
	maxheap.Insert(3)
	maxheap.Insert(2)
	maxheap.Insert(1)

	fmt.Println(maxheap.arr)

	maxheap.Delete()

	fmt.Println(maxheap.arr)

}
