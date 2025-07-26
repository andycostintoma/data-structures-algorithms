package main

import (
	"fmt"
)

// O(1)
func get(arr []int, index int) (int, error) {
	if index < 0 || index >= len(arr) {
		return 0, fmt.Errorf("index %d out of bounds", index)
	}
	return arr[index], nil
}

// O(n)
func search(arr []int, value int) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

// O(n)
func insert(arr []int, index int, value int) ([]int, error) {
	if index < 0 || index > len(arr) {
		return nil, fmt.Errorf("insert index %d out of bounds", index)
	}
	arr = append(arr, 0)             // grow slice
	copy(arr[index+1:], arr[index:]) // shift elements right
	arr[index] = value
	return arr, nil
}

// O(n)
func delete(arr []int, index int) ([]int, error) {
	if index < 0 || index >= len(arr) {
		return nil, fmt.Errorf("delete index %d out of bounds", index)
	}
	copy(arr[index:], arr[index+1:]) // shift elements left
	arr = arr[:len(arr)-1]           // shrink slice
	return arr, nil
}
