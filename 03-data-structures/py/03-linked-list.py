class Node:
    def __init__(self, value):
        self.value = value
        self.next = None


class LinkedList:
    def __init__(self):
        self.head = None
        self.size = 0

    # O(1)
    def prepend(self, node):
        if node is None:
            return  # Ensure node is not None
        node.next = self.head
        self.head = node
        self.size += 1

    # O(n)
    def append(self, node):
        if node is None:
            return  # Ensure node is not None
        self.insert_at_index(self.size, node)

    # O(n)
    def insert_at_index(self, index, node):
        if node is None or index < 0 or index > self.size:
            print("Index out of bounds")
            return  # Return without modifying the list

        if index == 0:
            self.prepend(node)
            return

        current = self.head
        for _ in range(index - 1):
            current = current.next
        node.next = current.next
        current.next = node
        self.size += 1

    # O(1)
    def insert_after_node(self, prev_node, node):
        if prev_node is None or node is None:
            print("Previous node or node cannot be None")
            return
        node.next = prev_node.next
        prev_node.next = node
        self.size += 1

    # O(n)
    def delete_at_index(self, index):
        if index < 0 or index >= self.size:
            print("Index out of bounds")
            return

        if index == 0:
            to_delete = self.head
            self.head = self.head.next
        else:
            current = self.head
            for _ in range(index - 1):
                current = current.next
            to_delete = current.next
            current.next = to_delete.next

        del to_delete
        self.size -= 1

    # O(1)
    def delete_node_after(self, prev_node):
        if prev_node is None or prev_node.next is None:
            print("Previous node is None or there is no node after the previous node")
            return
        node_to_delete = prev_node.next
        prev_node.next = node_to_delete.next
        del node_to_delete
        self.size -= 1

    # O(n)
    def get_at_index(self, index):
        if index < 0 or index >= self.size:
            print("Index out of bounds")
            return None

        current = self.head
        for _ in range(index):
            current = current.next
        return current

    # O(n)
    def search(self, value):
        current = self.head
        index = 0
        while current:
            if current.value == value:
                return index
            current = current.next
            index += 1
        return -1

    # O(n)
    def print_list(self):
        current = self.head
        values = []
        while current:
            values.append(str(current.value))
            current = current.next
        print(f"[{', '.join(values)}]")
