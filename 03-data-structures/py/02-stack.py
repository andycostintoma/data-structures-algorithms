class Stack:
    def __init__(self):
        self.items = []

    # O(1) - Simply returns the length of the list, which is tracked by the list object
    def size(self):
        return len(self.items)

    # O(1) - Accessing the last element (constant time)
    def peek(self):
        if len(self.items) == 0:
            return None
        return self.items[-1]

    # O(n) - Search operation: We need to iterate through the entire stack to find the item
    def search(self, item):
        for index, element in enumerate(self.items):
            if element == item:
                return index
        return -1  # Return -1 if item is not found

    # O(1) - Insertion at the end of the list (amortized constant time for dynamic arrays)
    def push(self, item):
        self.items.append(item)

    # O(1) - Deletion from the end of the list (constant time)
    def pop(self):
        if len(self.items) == 0:
            return None
        item = self.items[-1]
        del self.items[-1]
        return item


def main():
    stack = Stack()

    # Insertion (push items to stack)
    stack.push(1)  # O(1)
    stack.push(2)  # O(1)
    stack.push(3)  # O(1)

    # Size of the stack
    print(f"Stack size: {stack.size()}")  # O(1)

    # Peek operation (access the top item)
    print(f"Top of stack: {stack.peek()}")  # O(1)

    # Search for an item (return index where found)
    index = stack.search(2)  # O(n)
    if index != -1:
        print(f"Item 2 found at index {index}")
    else:
        print("Item 2 not found")

    # Pop operation (remove the top item)
    print(f"Popped item: {stack.pop()}")  # O(1)
    print(f"Stack size after pop: {stack.size()}")  # O(1)


if __name__ == "__main__":
    main()
