package main

import (
	"fmt"
)

func partition(arr []int) int {
	pivot := arr[len(arr)-1]
	i := 0
	for j := 0; j < len(arr)-1; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	return i
}

func quicksort(arr []int) {
	if len(arr) < 2 {
		return
	}
	pivotIdx := partition(arr)
	quicksort(arr[:pivotIdx])
	quicksort(arr[pivotIdx+1:])
}

func main() {
	arr := []int{8, 3, 1, 7, 0, 10, 2}
	fmt.Println("Before:", arr)
	quicksort(arr)
	fmt.Println("After: ", arr)
}
