package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// move elements bigger than key one position to the right
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		// insert key in the right position
		arr[j+1] = key
	}
}

func main() {
	arr := []int{3, 7, 4, 12, 6, 4, 5}
	fmt.Println("Initial array:", arr)
	insertionSort(arr)
	fmt.Println("Sorted array:", arr)
}
