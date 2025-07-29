package main

import (
	"fmt"
)

// Set represents a generic set of elements of type string
type Set struct {
	m map[string]struct{} // underlying map stores keys with empty struct{} as dummy value
}

// NewSet creates and returns a new Set with initial capacity
func NewSet() *Set {
	return &Set{
		m: make(map[string]struct{}),
	}
}

// -------------------- Insertion --------------------

// Average Time Complexity: O(1)
func (s *Set) Add(key string) {
	s.m[key] = struct{}{}
}

// -------------------- Search --------------------

// Average Time Complexity: O(1)
func (s *Set) Contains(key string) bool {
	_, exists := s.m[key]
	return exists
}

// -------------------- Deletion --------------------

// Average Time Complexity: O(1)
func (s *Set) Remove(key string) {
	delete(s.m, key)
}

// -------------------- Utilities --------------------

// Returns the number of elements in the set
func (s *Set) Size() int {
	return len(s.m)
}

// -------------------- Debug --------------------

// Returns string representation of the set
func (s *Set) String() string {
	keys := make([]string, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}
	return fmt.Sprintf("%v", keys)
}

func main() {
	set := NewSet()
	set.Add("apple")
	set.Add("banana")
	fmt.Println(set)                    // [apple banana]
	fmt.Println(set.Contains("banana")) // true
	set.Remove("banana")
	fmt.Println(set)                    // [apple]
	fmt.Println(set.Contains("banana")) // false
}
