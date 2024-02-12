package main

import "fmt"

func main() {
	arr := []int{5, 4, 10, 1, 6, 2}
	//insertionSort(arr)
	//selectionSort(arr)
	fmt.Println(arr)

	//quickSort(arr, 0, len(arr)-1)
	sortedArray := mergeSort(arr)
	fmt.Println(sortedArray)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)

}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result

}

func quickSort(arr []int, lb, ub int) {
	if lb < ub {
		loc := partition(arr, lb, ub)
		quickSort(arr, lb, loc-1)
		quickSort(arr, loc+1, ub)
	}
}
func partition(arr []int, lb, ub int) int {
	pivot := arr[lb]
	start := lb
	end := ub
	for start <= end {
		for start <= end && arr[start] <= pivot {
			start++
		}
		for start <= end && arr[end] > pivot {
			end--
		}
		if start < end {
			arr[start], arr[end] = arr[end], arr[start]
		}
	}
	arr[lb], arr[end] = arr[end], arr[lb] 
	return end
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
		fmt.Println(arr)

	}
}

func insertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
		fmt.Println(arr)
	}
}
