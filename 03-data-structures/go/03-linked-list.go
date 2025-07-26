package main

import (
	"errors"
	"fmt"
)

// Node represents a single element in the linked list.
type Node struct {
	Value int
	Next  *Node
}

// LinkedList represents the linked list structure.
type LinkedList struct {
	Head *Node
	Size int
}

// NewNode creates a new node with the given value. (O(1))
func NewNode(value int) *Node {
	return &Node{Value: value}
}

// NewLinkedList creates and returns a new, empty linked list. (O(1))
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Free frees all nodes in the linked list.
func (list *LinkedList) Free() {
	list.Head = nil
	list.Size = 0
}

// GetAtIndex returns the node at the given index. (O(n))
func (list *LinkedList) GetAtIndex(index int) (*Node, error) {
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
func (list *LinkedList) Search(value int) int {
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

// Prepend inserts a node at the head of the list. (O(1))
func (list *LinkedList) Prepend(node *Node) {
	if node == nil {
		return
	}
	node.Next = list.Head
	list.Head = node
	list.Size++
}

// Append inserts a node at the tail of the list. (O(n))
func (list *LinkedList) Append(node *Node) error {
	if node == nil {
		return errors.New("node cannot be nil")
	}
	return list.InsertAtIndex(list.Size, node)
}

// InsertAtIndex inserts a node at the specified index. (O(n))
func (list *LinkedList) InsertAtIndex(index int, node *Node) error {
	if index < 0 || index > list.Size {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		list.Prepend(node)
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

// InsertAfterNode inserts a node after the given node. (O(1))
func (list *LinkedList) InsertAfterNode(prevNode *Node, node *Node) error {
	if prevNode == nil || node == nil {
		return errors.New("previous node or node cannot be nil")
	}
	node.Next = prevNode.Next
	prevNode.Next = node
	list.Size++
	return nil
}

// DeleteAtIndex deletes the node at the given index. (O(n))
func (list *LinkedList) DeleteAtIndex(index int) error {
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
	list.Size--
	return nil
}

// DeleteNodeAfter deletes the node after the given node. (O(1))
func (list *LinkedList) DeleteNodeAfter(prevNode *Node) error {
	if prevNode == nil || prevNode.Next == nil {
		return errors.New("previous node is nil or there is no node after the previous node")
	}

	// The node to delete is the one after prevNode
	nodeToDelete := prevNode.Next

	// Link the previous node to the next node
	prevNode.Next = nodeToDelete.Next

	// Clear the node's next to help garbage collection
	nodeToDelete.Next = nil

	// Decrement the size of the list
	list.Size--

	return nil
}

// Print prints the linked list. (O(n))
func (list *LinkedList) Print() {
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
