package main

import "errors"

type Queue struct {
	list *LinkedList
	tail *Node
}

func NewQueue() *Queue {
	return &Queue{
		list: NewLinkedList(),
		tail: nil,
	}
}

// Peek returns the node at the front without removing it. O(1)
func (q *Queue) Peek() (*Node, error) {
	if q.list.Head == nil {
		return nil, errors.New("queue is empty")
	}
	return q.list.Head, nil
}

// Enqueue adds a node to the end of the queue. O(1)
func (q *Queue) Enqueue(node *Node) error {
	if node == nil {
		return errors.New("node cannot be nil")
	}
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
	return nil
}

// Dequeue removes and returns the node at the front of the queue. O(1)
func (q *Queue) Dequeue() (*Node, error) {
	if q.list.Head == nil {
		return nil, errors.New("queue is empty")
	}
	node := q.list.Head
	q.list.Head = node.Next
	node.Next = nil
	q.list.Size--

	if q.list.Head == nil {
		// Queue is now empty, reset tail
		q.tail = nil
	}
	return node, nil
}

// Size returns the current size of the queue.
func (q *Queue) Size() int {
	return q.list.Size
}
