package main

import "fmt"

func bubbleSort(arr []int) {
	swapping := true
	for swapping {
		swapping = false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapping = true
			}
		}
	}
}

func main() {
	arr := []int{3, 7, 4, 12, 6, 5}
	bubbleSort(arr)
	fmt.Println(arr)
}
