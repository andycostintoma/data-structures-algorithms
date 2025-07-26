from __future__ import annotations
from typing import Optional


class BSTNode:
    def __init__(self, val: Optional[int] = None) -> None:
        self.left: Optional[BSTNode] = None
        self.right: Optional[BSTNode] = None
        self.val: Optional[int] = val

    # -------------------- Access --------------------

    # Time Complexity: O(log n)
    def get_min(self) -> Optional[int]:
        if self.left is None:
            return self.val
        return self.left.get_min()

    # Time Complexity: O(log n)
    def get_max(self) -> Optional[int]:
        if self.right is None:
            return self.val
        return self.right.get_max()

    # -------------------- Search --------------------

    # Time Complexity: O(log n)
    def exists(self, val: int) -> bool:
        if val == self.val:
            return True
        if val < self.val:
            if self.left is None:
                return False
            return self.left.exists(val)
        if self.right is None:
            return False
        return self.right.exists(val)

    # -------------------- Insertion --------------------

    # Time Complexity: O(log n) average,
    def insert(self, val: int) -> None:
        if self.val is None:
            self.val = val
            return

        if self.val == val:
            return

        if val < self.val:
            if self.left:
                self.left.insert(val)
            else:
                self.left = BSTNode(val)
        else:
            if self.right:
                self.right.insert(val)
            else:
                self.right = BSTNode(val)

    # -------------------- Deletion --------------------

    # Time Complexity: O(log n) average
    def delete(self, val: int) -> Optional[BSTNode]:
        if self.val is None:
            return None

        if val < self.val:
            if self.left:
                self.left = self.left.delete(val)
            return self
        elif val > self.val:
            if self.right:
                self.right = self.right.delete(val)
            return self
        else:
            if self.right is None:
                return self.left
            if self.left is None:
                return self.right

            min_larger_node = self.right
            while min_larger_node.left:
                min_larger_node = min_larger_node.left
            self.val = min_larger_node.val
            self.right = self.right.delete(min_larger_node.val)
            return self

    # -------------------- Other operations --------------------

    # Time Complexity: O(n)
    def height(self) -> int:
        if self.val is None:
            return 0
        lh = self.left.height() if self.left else 0
        rh = self.right.height() if self.right else 0
        return max(lh, rh) + 1

    # Time Complexity: O(n)
    # Preorder = current node visited before its children
    def preorder_traversal(self, visited: list[int]) -> list[int]:
        visited.append(self.val)
        if self.left:
            self.left.preorder_traversal(visited)
        if self.right:
            self.right.preorder_traversal(visited)
        return visited
    
    # Time Complexity: O(n)
    # Postorder = current node visited after its children
    def postorder_traversal(self, visited: list[int]) -> list[int]:
        if self.left:
            self.left.postorder_traversal(visited)
        if self.right:
            self.right.postorder_traversal(visited)
        visited.append(self.val)
        return visited


    # Time Complexity: O(n)
    def inorder_taversal(self, visited: list[int]) -> list[int]:
        if self.left:
            self.left.inorder_taversal(visited)
        visited.append(self.val)
        if self.right:
            self.right.inorder_taversal(visited)
        return visited


    # Time Complexity: O(n)
    def print(self, level: int = 0) -> None:
        if self.right:
            self.right.print(level + 1)
        print(" " * (level * 4) + str(self.val))
        if self.left:
            self.left.print(level + 1)


def main():
    # Create the root of the BST
    root = BSTNode()

    # Insert elements
    for val in [50, 30, 70, 20, 40, 60, 80]:
        root.insert(val)

    print("Tree structure:")
    root.print()

    print("\nPreorder traversal:")
    print(root.preorder_traversal([]))

    print("\nPostorder traversal:")
    print(root.postorder_traversal([]))

    print("\nInorder traversal:")
    print(root.inorder_taversal([]))  # Should be sorted

    print("\nMin value:", root.get_min())
    print("Max value:", root.get_max())
    print("Height of tree:", root.height())
    print("Exists 40?", root.exists(40))
    print("Exists 100?", root.exists(100))

    # Delete an element
    print("\nDeleting 70...")
    root.delete(70)
    print("Tree after deletion:")
    root.print()


if __name__ == "__main__":
    main()
