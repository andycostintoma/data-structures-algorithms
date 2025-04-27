package main

import (
	"errors"
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

// O(1)
func pop(arr []int) ([]int, int, error) {
	if len(arr) == 0 {
		return nil, 0, errors.New("cannot pop from empty array")
	}
	value := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	return arr, value, nil
}

// ---- Main Demo ----

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("After initialization:", arr)

	arr, err := insert(arr, 5, 99)
	if err != nil {
		fmt.Println("Insert error:", err)
		return
	}
	fmt.Println("After insert at index 5:", arr)

	arr, err = delete(arr, 2)
	if err != nil {
		fmt.Println("Delete error:", err)
		return
	}
	fmt.Println("After delete at index 2:", arr)

	// O(1)
	arr = append(arr, 100)
	fmt.Println("After appending 100:", arr)

	var popped int
	arr, popped, err = pop(arr)
	if err != nil {
		fmt.Println("Pop error:", err)
		return
	}
	fmt.Printf("Popped value: %d\n", popped)
	fmt.Println("After popping:", arr)

	foundIndex := search(arr, 1)
	fmt.Printf("Found 1 at index: %d\n", foundIndex)

	fmt.Println("Final array:", arr)
}
