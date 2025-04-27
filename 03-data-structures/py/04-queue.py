class Node:
    """Represents a single element in the linked list."""
    def __init__(self, value):
        self.value = value
        self.next = None


class Queue:
    """Queue implementation using a linked list."""
    def __init__(self):
        self.head = None
        self.tail = None
        self.size = 0

    def is_empty(self):
        """Returns True if the queue is empty, otherwise False."""
        return self.size == 0

    def enqueue(self, value):
        """Adds an element to the end of the queue (O(1))."""
        new_node = Node(value)
        if self.tail is None:
            # If the queue is empty, set both head and tail to the new node
            self.head = self.tail = new_node
        else:
            # Otherwise, add the new node to the tail and update the tail
            self.tail.next = new_node
            self.tail = new_node
        self.size += 1

    def dequeue(self):
        """Removes and returns the element at the front of the queue (O(1))."""
        if self.is_empty():
            raise IndexError("Dequeue from empty queue")
        node_to_return = self.head
        self.head = self.head.next
        if self.head is None:
            # If the queue is now empty, reset the tail as well
            self.tail = None
        self.size -= 1
        node_to_return.next = None  # Disconnect the node
        return node_to_return.value

    def peek(self):
        """Returns the element at the front of the queue without removing it (O(1))."""
        if self.is_empty():
            raise IndexError("Peek from empty queue")
        return self.head.value

    def get_size(self):
        """Returns the number of elements in the queue (O(1))."""
        return self.size

    def print_queue(self):
        """Prints the queue elements (O(n))."""
        if self.is_empty():
            print("[]")
            return
        current = self.head
        values = []
        while current:
            values.append(str(current.value))
            current = current.next
        print(f"[{', '.join(values)}]")


def main():
    # Example usage:
    queue = Queue()

    # Enqueue elements
    queue.enqueue(10)
    queue.enqueue(20)
    queue.enqueue(30)

    # Print the queue
    queue.print_queue()  # Expected: [10, 20, 30]

    # Peek the front element
    print(f"Front element: {queue.peek()}")  # Expected: 10

    # Dequeue an element
    print(f"Dequeued element: {queue.dequeue()}")  # Expected: 10

    # Print the queue after dequeue
    queue.print_queue()  # Expected: [20, 30]


if __name__ == "__main__":
    main()
