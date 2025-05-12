class RBNode:
    def __init__(self, val: int) -> None:
        self.red: bool = False  # Node color, False for black, True for red
        self.parent: RBNode = None  # Parent node
        self.val: int = val  # Node value
        self.left: RBNode = None  # Left child
        self.right: RBNode = None  # Right child

class RBTree:
    def __init__(self) -> None:
        self.nil: RBNode = RBNode(None)  # Sentinel node (nil) with None value
        self.nil.red = False  # Nil is always black
        self.nil.left = None
        self.nil.right = None
        self.root: RBNode = self.nil  # Initially the tree is empty, root is nil

    def insert(self, val: int) -> None:
        """
        Inserts a value into the Red-Black Tree and fixes any violations of Red-Black properties.
        
        Time Complexity: O(log n) for searching and inserting the node. Fixing insert takes O(log n) due to tree height.
        """
        new_node: RBNode = RBNode(val)
        new_node.parent = None
        new_node.left = self.nil
        new_node.right = self.nil
        new_node.red = True

        parent: RBNode = None
        current: RBNode = self.root
        while current != self.nil:
            parent = current
            if new_node.val < current.val:
                current = current.left
            elif new_node.val > current.val:
                current = current.right
            else:
                # duplicate, just ignore
                return

        new_node.parent = parent
        if parent is None:
            self.root = new_node
        elif new_node.val < parent.val:
            parent.left = new_node
        else:
            parent.right = new_node

        self.fix_insert(new_node)

    def fix_insert(self, new_node: RBNode) -> None:
        """
        Fixes any violations of Red-Black Tree properties after an insertion.
        
        Time Complexity: O(log n) due to the tree height.
        """
        while new_node != self.root and new_node.parent.red:
            parent: RBNode = new_node.parent
            grandparent: RBNode = parent.parent
            if parent == grandparent.right:
                uncle: RBNode = grandparent.left
                if uncle.red:
                    # Case 1: Uncle is red
                    uncle.red = False
                    parent.red = False
                    grandparent.red = True
                    new_node = grandparent
                else:
                    # Case 2: Uncle is black and new_node is on the left
                    if new_node == parent.left:
                        new_node = parent
                        self.rotate_right(new_node)
                        parent = new_node.parent
                    parent.red = False
                    grandparent.red = True
                    self.rotate_left(grandparent)
            else:
                uncle = grandparent.right
                if uncle.red:
                    # Case 1: Uncle is red
                    uncle.red = False
                    parent.red = False
                    grandparent.red = True
                    new_node = grandparent
                else:
                    # Case 2: Uncle is black and new_node is on the right
                    if new_node == parent.right:
                        new_node = parent
                        self.rotate_left(new_node)
                        parent = new_node.parent
                    parent.red = False
                    grandparent.red = True
                    self.rotate_right(grandparent)
        self.root.red = False  # The root is always black

    def exists(self, val: int) -> RBNode:
        """
        Searches for a node with the given value in the tree.
        
        Time Complexity: O(log n) for searching through the tree.
        """
        curr: RBNode = self.root
        while curr != self.nil and val != curr.val:
            if val < curr.val:
                curr = curr.left
            else:
                curr = curr.right
        return curr  # Returns nil if the node doesn't exist

    def rotate_left(self, pivot_parent: RBNode) -> None:
        """
        Performs a left rotation around the pivot_parent node.
        
        Time Complexity: O(1) since it involves constant-time pointer manipulations.
        """
        if pivot_parent == self.nil or pivot_parent.right == self.nil:
            return
        pivot: RBNode = pivot_parent.right
        pivot_parent.right = pivot.left
        if pivot.left != self.nil:
            pivot.left.parent = pivot_parent

        pivot.parent = pivot_parent.parent
        if pivot_parent.parent is None:
            self.root = pivot
        elif pivot_parent == pivot_parent.parent.left:
            pivot_parent.parent.left = pivot
        else:
            pivot_parent.parent.right = pivot
        pivot.left = pivot_parent
        pivot_parent.parent = pivot

    def rotate_right(self, pivot_parent: RBNode) -> None:
        """
        Performs a right rotation around the pivot_parent node.
        
        Time Complexity: O(1) since it involves constant-time pointer manipulations.
        """
        if pivot_parent == self.nil or pivot_parent.left == self.nil:
            return
        pivot: RBNode = pivot_parent.left
        pivot_parent.left = pivot.right
        if pivot.right != self.nil:
            pivot.right.parent = pivot_parent

        pivot.parent = pivot_parent.parent
        if pivot_parent.parent is None:
            self.root = pivot
        elif pivot_parent == pivot_parent.parent.right:
            pivot_parent.parent.right = pivot
        else:
            pivot_parent.parent.left = pivot
        pivot.right = pivot_parent
        pivot_parent.parent = pivot

    def pretty_print(self, node: RBNode = None, level: int = 0) -> None:
        """
        Pretty prints the Red-Black Tree with more indentation for each level.
        
        Time Complexity: O(n), where n is the number of nodes, as each node is visited once.
        """
        if node is None:
            node = self.root
        if node != self.nil:
            # Recursively print right subtree first (to create the top-down effect)
            if node.right != self.nil:
                self.pretty_print(node.right, level + 1)
                
            # Print current node with more indentation
            print("        " * level + f"{node.val} [{ 'R' if node.red else 'B' }]")

            # Recursively print left subtree
            if node.left != self.nil:
                self.pretty_print(node.left, level + 1)

def main() -> None:
    """
    Main function to test the Red-Black Tree implementation.
    """
    rbt = RBTree()

    # Inserting nodes
    rbt.insert(10)
    rbt.insert(20)
    rbt.insert(15)
    rbt.insert(30)
    rbt.insert(25)
    rbt.insert(5)

    # Pretty print the tree
    rbt.pretty_print()

if __name__ == "__main__":
    main()
