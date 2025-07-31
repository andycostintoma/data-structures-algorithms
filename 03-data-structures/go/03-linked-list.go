package main

import (
	"errors"
	"fmt"
)

// Node represents a single element in the linked list.
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// NewNode creates a new node with the given value. (O(1))
func NewNode[T comparable](value T) *Node[T] {
	return &Node[T]{Value: value}
}

// LinkedList represents the linked list structure.
type LinkedList[T comparable] struct {
	Head *Node[T]
	Size int
}

// NewLinkedList creates and returns a new, empty linked list. (O(1))
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Free frees all nodes in the linked list.
func (list *LinkedList[T]) Free() {
	list.Head = nil
	list.Size = 0
}

// GetAtIndex returns the node at the given index. (O(n))
func (list *LinkedList[T]) GetAtIndex(index int) (*Node[T], error) {
	if index < 0 || index >= list.Size {
		return nil, errors.New("index out of bounds")
	}
	current := list.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current, nil
}

// Search searches for a value and returns its index or -1 if not found. (O(n))
func (list *LinkedList[T]) Search(value T) int {
	current := list.Head
	index := 0
	for current != nil {
		if current.Value == value {
			return index
		}
		current = current.Next
		index++
	}
	return -1
}

// Prepend inserts a value at the head of the list. (O(1))
func (list *LinkedList[T]) Prepend(value T) {
	node := NewNode(value)
	node.Next = list.Head
	list.Head = node
	list.Size++
}

// Append inserts a value at the tail of the list. (O(n))
func (list *LinkedList[T]) Append(value T) error {
	return list.InsertAtIndex(list.Size, value)
}

// InsertAtIndex inserts a value at the specified index. (O(n))
func (list *LinkedList[T]) InsertAtIndex(index int, value T) error {
	if index < 0 || index > list.Size {
		return errors.New("index out of bounds")
	}
	node := NewNode(value)
	if index == 0 {
		list.Prepend(value)
		return nil
	}
	current := list.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	node.Next = current.Next
	current.Next = node
	list.Size++
	return nil
}

// InsertAfterNode inserts a value after the given node. (O(1))
func (list *LinkedList[T]) InsertAfterNode(prevNode *Node[T], value T) error {
	if prevNode == nil {
		return errors.New("previous node cannot be nil")
	}
	node := NewNode(value)
	node.Next = prevNode.Next
	prevNode.Next = node
	list.Size++
	return nil
}

// DeleteAtIndex deletes the node at the given index. (O(n))
func (list *LinkedList[T]) DeleteAtIndex(index int) error {
	if index < 0 || index >= list.Size {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		list.Head = list.Head.Next
		list.Size--
		return nil
	}
	current := list.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	toDelete := current.Next
	current.Next = toDelete.Next
	toDelete.Next = nil
	list.Size--
	return nil
}

// DeleteNodeAfter deletes the node after the given node. (O(1))
func (list *LinkedList[T]) DeleteNodeAfter(prevNode *Node[T]) error {
	if prevNode == nil || prevNode.Next == nil {
		return errors.New("previous node is nil or there is no node after it")
	}
	nodeToDelete := prevNode.Next
	prevNode.Next = nodeToDelete.Next
	nodeToDelete.Next = nil
	list.Size--
	return nil
}

// Print prints the linked list. (O(n))
func (list *LinkedList[T]) Print() {
	if list.Head == nil {
		fmt.Println("[]")
		return
	}
	current := list.Head
	fmt.Print("[")
	for current != nil {
		fmt.Print(current.Value)
		if current.Next != nil {
			fmt.Print(", ")
		}
		current = current.Next
	}
	fmt.Println("]")
}
