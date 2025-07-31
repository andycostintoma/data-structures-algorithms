package main

import (
	"errors"
	"fmt"
)

type entry struct {
	key     string
	value   any
	deleted bool
}

type HashMap struct {
	entries []entry
	size    int // count of active (non-deleted) entries
}

// NewHashMap creates a new hashmap with given initial capacity.
func NewHashMap(capacity int) *HashMap {
	return &HashMap{
		entries: make([]entry, capacity),
		size:    0,
	}
}

// keyToIndex computes hash index by summing ASCII codes mod capacity.
func (h *HashMap) keyToIndex(key string) int {
	total := 0
	for _, c := range key {
		total += int(c)
	}
	return total % len(h.entries)
}

// Get returns the value for a given key or error if not found.
func (h *HashMap) Get(key string) (any, error) {
	index := h.keyToIndex(key)
	originalIndex := index
	firstIteration := true

	for {
		e := h.entries[index]
		if e.key == "" && !e.deleted {
			// empty slot, key not found
			break
		}
		if !e.deleted && e.key == key {
			return e.value, nil
		}

		index = (index + 1) % len(h.entries)
		if !firstIteration && index == originalIndex {
			break
		}
		firstIteration = false
	}

	return nil, errors.New("key not found")
}

// Insert adds or updates the key with given value.
func (h *HashMap) Insert(key string, value any) {
	h.resizeIfNeeded()
	h.insertWithoutResize(key, value)
}

// insertWithoutResize inserts without checking resize.
func (h *HashMap) insertWithoutResize(key string, value any) {
	index := h.keyToIndex(key)
	for {
		e := h.entries[index]
		if e.key == "" || e.deleted || e.key == key {
			// Insert new or update existing
			if e.key == "" || e.deleted {
				h.size++
			}
			h.entries[index] = entry{key: key, value: value, deleted: false}
			return
		}
		index = (index + 1) % len(h.entries)
	}
}

// Delete marks the key as deleted or errors if not found.
func (h *HashMap) Delete(key string) error {
	index := h.keyToIndex(key)
	originalIndex := index
	firstIteration := true

	for {
		e := h.entries[index]
		if e.key == "" && !e.deleted {
			// empty slot, key not found
			break
		}
		if !e.deleted && e.key == key {
			h.entries[index].deleted = true
			h.size--
			return nil
		}

		index = (index + 1) % len(h.entries)
		if !firstIteration && index == originalIndex {
			break
		}
		firstIteration = false
	}

	return errors.New("key not found")
}

// Size returns number of active key-value pairs.
func (h *HashMap) Size() int {
	return h.size
}

// loadFactor returns current load factor (active / capacity).
func (h *HashMap) loadFactor() float64 {
	return float64(h.size) / float64(len(h.entries))
}

// resizeIfNeeded doubles capacity if load factor > 0.7
func (h *HashMap) resizeIfNeeded() {
	if h.loadFactor() <= 0.7 {
		return
	}

	oldEntries := h.entries
	newCapacity := len(oldEntries) * 2
	h.entries = make([]entry, newCapacity)
	h.size = 0

	for _, e := range oldEntries {
		if e.key != "" && !e.deleted {
			h.insertWithoutResize(e.key, e.value)
		}
	}
}

// String returns a string representation of active entries.
func (h *HashMap) String() string {
	result := "["
	first := true
	for _, e := range h.entries {
		if e.key != "" && !e.deleted {
			if !first {
				result += ", "
			}
			result += fmt.Sprintf("(%q: %v)", e.key, e.value)
			first = false
		}
	}
	result += "]"
	return result
}
