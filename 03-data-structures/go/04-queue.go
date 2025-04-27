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
	Tail *Node
	Size int // This is the actual size of the queue
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
	list.Head = nil
	list.Tail = nil
	list.Size = 0
}

// Enqueue adds an element to the end of the queue (O(1))
func (list *LinkedList) Enqueue(value int) {
	newNode := NewNode(value)
	if list.Tail == nil {
		// If the queue is empty, set both head and tail to the new node
		list.Head = newNode
		list.Tail = newNode
	} else {
		// Add the new node at the end and update the tail
		list.Tail.Next = newNode
		list.Tail = newNode
	}
	list.Size++
}

// Dequeue removes and returns the element at the front of the queue. (O(1))
func (list *LinkedList) Dequeue() (*Node, error) {
	if list.Size == 0 {
		return nil, errors.New("queue is empty")
	}
	nodeToReturn := list.Head
	list.Head = list.Head.Next
	if list.Head == nil {
		// If the queue is now empty, reset the tail as well
		list.Tail = nil
	}
	list.Size--
	nodeToReturn.Next = nil // Disconnect the node from the rest of the list
	return nodeToReturn, nil
}

// Peek returns the element at the front of the queue without removing it. (O(1))
func (list *LinkedList) Peek() (*Node, error) {
	if list.Size == 0 {
		return nil, errors.New("queue is empty")
	}
	return list.Head, nil
}

// IsEmpty checks if the queue is empty. (O(1))
func (list *LinkedList) IsEmpty() bool {
	return list.Size == 0
}

// GetSize returns the number of elements in the queue. (O(1))
func (list *LinkedList) GetSize() int {
	return list.Size
}

// Print prints the queue's elements. (O(n))
func (list *LinkedList) Print() {
	if list.Size == 0 {
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
	// Example usage of the Queue
	queue := NewLinkedList()

	// Enqueue some elements
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	// Print the queue
	queue.Print() // Expected: [10, 20, 30]

	// Peek the front element
	front, _ := queue.Peek()
	fmt.Printf("Front element: %d\n", front.Value) // Expected: Front element: 10

	// Dequeue an element
	dequeued, _ := queue.Dequeue()
	fmt.Printf("Dequeued element: %d\n", dequeued.Value) // Expected: Dequeued element: 10

	// Print the queue after dequeue
	queue.Print() // Expected: [20, 30]
}
