package main

import (
	"errors"
	"fmt"
)

// Stack represents a stack data structure
type Stack struct {
	items []int
}

// Size returns the size of the stack (O(1))
func (s *Stack) Size() int {
	return len(s.items)
}

// Peek returns the top item of the stack (O(1))
// Returns an error if the stack is empty
func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// Search returns the index of the item if found (O(n))
// Returns -1 if the item is not found
func (s *Stack) Search(item int) int {
	for i, elem := range s.items {
		if elem == item {
			return i
		}
	}
	return -1 // return -1 if item is not found
}

// Push adds an item to the stack (O(1))
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack (O(1))
// Returns an error if the stack is empty
func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func main() {
	stack := &Stack{}

	// Insertion (push items to stack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Size of the stack
	fmt.Println("Stack size:", stack.Size())

	// Peek operation (access the top item)
	top, err := stack.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Top of stack:", top)
	}

	// Search for an item (return index where found)
	index := stack.Search(2)
	if index != -1 {
		fmt.Printf("Item 2 found at index %d\n", index)
	} else {
		fmt.Println("Item 2 not found")
	}

	// Pop operation (remove the top item)
	poppedItem, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Popped item:", poppedItem)
	}
	fmt.Println("Stack size after pop:", stack.Size())
}
