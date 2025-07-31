package main

import "errors"

// Queue represents a FIFO queue using a linked list.
type Queue[T comparable] struct {
	list *LinkedList[T]
	tail *Node[T]
}

// NewQueue creates and returns a new empty queue.
func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		list: NewLinkedList[T](),
		tail: nil,
	}
}

// Peek returns the value at the front without removing it. O(1)
func (q *Queue[T]) Peek() (T, error) {
	var zero T
	if q.list.Head == nil {
		return zero, errors.New("queue is empty")
	}
	return q.list.Head.Value, nil
}

// Enqueue adds a value to the end of the queue. O(1)
func (q *Queue[T]) Enqueue(value T) {
	node := NewNode(value)
	node.Next = nil

	if q.tail == nil {
		// Queue is empty
		q.list.Head = node
		q.tail = node
	} else {
		q.tail.Next = node
		q.tail = node
	}
	q.list.Size++
}

// Dequeue removes and returns the value at the front of the queue. O(1)
func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.list.Head == nil {
		return zero, errors.New("queue is empty")
	}
	node := q.list.Head
	q.list.Head = node.Next
	node.Next = nil
	q.list.Size--

	if q.list.Head == nil {
		// Queue is now empty, reset tail
		q.tail = nil
	}
	return node.Value, nil
}

// Size returns the current size of the queue.
func (q *Queue[T]) Size() int {
	return q.list.Size
}
