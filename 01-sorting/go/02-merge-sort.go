package main

import "fmt"

func merge(left []int, right []int) []int {
	i, j := 0, 0
	result := make([]int, 0, len(right)+len(left))
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
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

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left, right := arr[:mid], arr[mid:]
	return merge(mergeSort(left), mergeSort(right))

}

func main() {
	arr := []int{2, 7, 5, 3, 1, 9, 8}
	sorted := mergeSort(arr)
	fmt.Println(sorted)
}
