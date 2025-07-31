package main

import (
	"fmt"
)

// Set represents a generic set of elements of type T
// T must be comparable so it can be used as a map key
type Set[T comparable] struct {
	m map[T]struct{} // underlying map stores keys with empty struct{} as dummy value
}

// NewSet creates and returns a new Set with initial capacity
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

// -------------------- Insertion --------------------

// Average Time Complexity: O(1)
func (s *Set[T]) Add(key T) {
	s.m[key] = struct{}{}
}

// -------------------- Search --------------------

// Average Time Complexity: O(1)
func (s *Set[T]) Contains(key T) bool {
	_, exists := s.m[key]
	return exists
}

// -------------------- Deletion --------------------

// Average Time Complexity: O(1)
func (s *Set[T]) Remove(key T) {
	delete(s.m, key)
}

// -------------------- Utilities --------------------

// Returns the number of elements in the set
func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) ToSlice() []T {
	out := make([]T, 0, len(s.m))
	for k := range s.m {
		out = append(out, k)
	}
	return out
}

// Returns string representation of the set
func (s *Set[T]) String() string {
	keys := make([]T, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}
	return fmt.Sprintf("%v", keys)
}
