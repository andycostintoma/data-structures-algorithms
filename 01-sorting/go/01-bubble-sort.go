package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		n--
	}
}

func main() {
	arr := []int{3, 7, 4, 12, 6, 5}
	fmt.Println("Initial array:", arr)
	bubbleSort(arr)
	fmt.Println("Sorted array:", arr)
}
