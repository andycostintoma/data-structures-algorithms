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

// Free frees all nodes in the linked list. (O(n))
func (list *LinkedList) Free() {
	current := list.Head
	for current != nil {
		next := current.Next
		current = next
	}
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

// DeleteNode deletes a specific node from the list. (O(n))
func (list *LinkedList) DeleteNode(nodeToDelete *Node) error {
	if nodeToDelete == nil {
		return errors.New("node to delete cannot be nil")
	}
	if nodeToDelete == list.Head {
		list.Head = list.Head.Next
		list.Size--
		return nil
	}
	current := list.Head
	for current != nil && current.Next != nodeToDelete {
		current = current.Next
	}
	if current == nil {
		return errors.New("node not found in the list")
	}
	current.Next = nodeToDelete.Next
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

func main() {
	// Example usage
	list := NewLinkedList()

	// Prepend nodes
	list.Prepend(NewNode(10))
	list.Prepend(NewNode(20))
	list.Prepend(NewNode(30))

	// Print the list
	list.Print() // Expected: [30, 20, 10]

	// Append a node
	if err := list.Append(NewNode(5)); err != nil {
		fmt.Println("Error:", err)
	}

	// Print the list
	list.Print() // Expected: [30, 20, 10, 5]

	// Search for a value
	index := list.Search(20)
	fmt.Printf("Index of 20: %d\n", index) // Expected: Index of 20: 1

	// Delete a node
	if err := list.DeleteAtIndex(1); err != nil {
		fmt.Println("Error:", err)
	}

	// Print the list after deletion
	list.Print() // Expected: [30, 10, 5]
}
