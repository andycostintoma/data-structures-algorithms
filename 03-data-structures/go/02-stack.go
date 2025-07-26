package main

import (
	"errors"
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
