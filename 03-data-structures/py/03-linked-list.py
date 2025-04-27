class Node:
    def __init__(self, val):
        self.val = val
        self.next = None

    def set_next(self, node):
        self.next = node

    def __repr__(self):
        return str(self.val)

class LinkedList:

    def __init__(self):
        self.head = None
        self.size = 0

    def __iter__(self):
        node = self.head
        while node is not None:
            yield node
            node = node.next

    def __repr__(self):
        nodes = []
        for node in self:
            nodes.append(str(node.val))
        return " -> ".join(nodes)

    # O(n) - Linear time complexity, as we need to traverse the list to the index
    def get_at_index(self, index):
        """Returns the value at a given index."""
        if index < 0 or index >= self.size:
            raise IndexError("Index out of bounds")
        
        current_node = self.head
        for _ in range(index):
            current_node = current_node.next
        return current_node.val
    
    # O(n) - Linear time complexity, as we need to search through the entire list
    def search(self, value):
        """Searches for a value and returns its index, or -1 if not found."""
        current_node = self.head
        index = 0
        while current_node is not None:
            if current_node.val == value:
                return index
            current_node = current_node.next
            index += 1
        return -1
    
    # O(n) - Linear time complexity, as we need to traverse to the given index
    def insert_at_index(self, index, node):
        """Inserts a node at a given index."""
        if index < 0 or index > self.size:
            raise IndexError("Index out of bounds")
        
        if index == 0:
            self.prepend(node)
            return
        
        current_node = self.head
        for _ in range(index - 1):
            current_node = current_node.next
        node.set_next(current_node.next)
        current_node.set_next(node)
        self.size += 1

    # O(1) - Constant time complexity
    def prepend(self, node):
        """Prepends a node to the head of the list."""
        node.next = self.head
        self.head = node
        self.size += 1

    # O(n) - Linear time complexity, as we need to traverse to the given index
    def append(self, node):
        """Appends a node to the tail of the list using insert_at_index."""
        self.insert_at_index(self.size, node)

    # O(n) - Linear time complexity, as we need to traverse to the given index
    def delete_at_index(self, index):
        """Deletes a node at the given index and returns its value."""
        if index < 0 or index >= self.size:
            raise IndexError("Index out of bounds")
        
        if index == 0:
            deleted_node = self.head
            self.head = self.head.next
        else:
            current_node = self.head
            for _ in range(index - 1):
                current_node = current_node.next
            deleted_node = current_node.next
            current_node.set_next(deleted_node.next)
        
        self.size -= 1
        return deleted_node.val

    # O(1) - Constant time complexity
    def free(self):
        """Frees the entire list."""
        self.head = None
        self.size = 0


# Main function to test LinkedList
def main():
    linked_list = LinkedList()

    # Prepend nodes
    linked_list.prepend(Node(10))
    linked_list.prepend(Node(20))
    linked_list.append(Node(30))

    print("List after adding elements:")
    print(linked_list)  # Expected Output: 20 -> 10 -> 30

    # Insert at index 1
    linked_list.insert_at_index(1, Node(25))
    print("List after inserting at index 1:")
    print(linked_list)  # Expected Output: 20 -> 25 -> 10 -> 30

    # Delete at index 2
    linked_list.delete_at_index(2)
    print("List after deleting at index 2:")
    print(linked_list)  # Expected Output: 20 -> 25 -> 30

    # Access value at index 1
    print("Value at index 1:", linked_list.get_at_index(1))  # Expected Output: 25

    # Search for value 30
    print("Index of value 30:", linked_list.search(30))  # Expected Output: 2

    # Free the list
    linked_list.free()
    print("List after freeing:")
    print(linked_list)  # Expected Output: (empty list)
    

if __name__ == "__main__":
    main()
