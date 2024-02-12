package main

import "fmt"

func heapSort(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for p := n - 1; p >= 0; p-- {
		arr[0], arr[p] = arr[p], arr[0]
		heapify(arr, p, 0)
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	l := i*2 + 1
	r := i*2 + 2
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		heapify(arr, n, largest)
	}
}

func main() {
	var arr = []int{2, 5, 1, 88, 11, 5, 2, 1}
	fmt.Println("before sorting:", arr)
	heapSort(arr)
	fmt.Println("after sorting:", arr)
}
