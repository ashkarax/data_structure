package main

import "fmt"

type MinHeap struct {
	arr []int
}

func (mh *MinHeap) Insert(data int) {
	mh.arr = append(mh.arr, data)
	mh.buildHeap()
}

func (mh *MinHeap) buildHeap() {
	n := len(mh.arr)
	for i := n/2 - 1; i >= 0; i-- {
		mh.minHeapify(i)
	}
}

func (mh *MinHeap) minHeapify(i int) {
	var k int
	n := len(mh.arr)
	for {
		l := i*2 + 1
		r := i*2 + 2
		if l >= n {
			break
		}
		if r < n && mh.arr[r] < mh.arr[l] {
			k = r
		} else {
			k = l
		}
		if mh.arr[k] < mh.arr[i] {
			mh.arr[k], mh.arr[i] = mh.arr[i], mh.arr[k]
			i = k
		} else {
			break
		}
	}
}

func main() {
	var mh MinHeap
	mh.Insert(4)
	mh.Insert(2)
	mh.Insert(0)
	mh.Insert(88)
	mh.Insert(11)
	mh.Insert(6)

	fmt.Println(mh.arr)

}
