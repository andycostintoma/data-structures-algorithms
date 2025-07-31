package main

import (
	"errors"
)

// Stack is a generic stack data structure
type Stack[T comparable] struct {
	items []T
}

// NewStack creates and returns a new empty Stack
func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

// Size returns the size of the stack (O(1))
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Peek returns the top item of the stack (O(1))
// Returns an error if the stack is empty
func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if len(s.items) == 0 {
		return zero, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// Search returns the index of the item if found (O(n))
// Returns -1 if the item is not found
func (s *Stack[T]) Search(item T) int {
	for i, elem := range s.items {
		if elem == item {
			return i
		}
	}
	return -1
}

// Push adds an item to the stack (O(1))
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack (O(1))
// Returns an error if the stack is empty
func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(s.items) == 0 {
		return zero, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}
