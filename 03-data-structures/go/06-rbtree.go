package main

import (
	"fmt"
	"strings"
)

// Define the RBNode struct
type RBNode struct {
	val    int
	red    bool
	parent *RBNode
	left   *RBNode
	right  *RBNode
}

/*
	Define the RBTree struct

In addition to all the rules of a Binary Search Tree, a red-black tree must follow some additional ones:
1. Each node is either red or black
2. The root is black. This rule is sometimes omitted. Since the root can always be changed from red to black, but not necessarily vice versa, this rule has little effect on analysis.
3. All Nil leaf nodes are black.
4. If a node is red, then both its children are black.
5. All paths from a single node go through the same number of black nodes to reach any of its descendant NIL nodes.
*/
type RBTree struct {
	nilNode *RBNode
	root    *RBNode
}

// Constructor for creating a new Red-Black Tree
func NewRBTree() *RBTree {
	nilNode := &RBNode{red: false} // Nil node is always black
	return &RBTree{
		nilNode: nilNode,
		root:    nilNode,
	}
}

// Insert inserts a value into the Red-Black Tree and fixes any violations
func (tree *RBTree) Insert(val int) {
	newNode := &RBNode{
		val:    val,
		red:    true,
		left:   tree.nilNode,
		right:  tree.nilNode,
		parent: nil,
	}

	var parent *RBNode
	current := tree.root

	// Find the correct place to insert
	for current != tree.nilNode {
		parent = current
		if newNode.val < current.val {
			current = current.left
		} else if newNode.val > current.val {
			current = current.right
		} else {
			// Ignore duplicates
			return
		}
	}

	// Set the parent and insert the new node
	newNode.parent = parent
	if parent == nil {
		tree.root = newNode
	} else if newNode.val < parent.val {
		parent.left = newNode
	} else {
		parent.right = newNode
	}

	// Fix any violations of Red-Black properties
	tree.fixInsert(newNode)
}

// FixInsert fixes any violations after an insert operation
func (tree *RBTree) fixInsert(newNode *RBNode) {
	for newNode != tree.root && newNode.parent.red {
		parent := newNode.parent
		grandparent := parent.parent
		var uncle *RBNode

		if parent == grandparent.right {
			uncle = grandparent.left
			if uncle.red {
				// Case 1: Uncle is red
				uncle.red = false
				parent.red = false
				grandparent.red = true
				newNode = grandparent
			} else {
				// Case 2: Uncle is black and newNode is on the left
				if newNode == parent.left {
					newNode = parent
					tree.rotateRight(newNode)
					parent = newNode.parent
				}
				parent.red = false
				grandparent.red = true
				tree.rotateLeft(grandparent)
			}
		} else {
			uncle = grandparent.right
			if uncle.red {
				// Case 1: Uncle is red
				uncle.red = false
				parent.red = false
				grandparent.red = true
				newNode = grandparent
			} else {
				// Case 2: Uncle is black and newNode is on the right
				if newNode == parent.right {
					newNode = parent
					tree.rotateLeft(newNode)
					parent = newNode.parent
				}
				parent.red = false
				grandparent.red = true
				tree.rotateRight(grandparent)
			}
		}
	}
	tree.root.red = false // The root is always black
}

// RotateLeft performs a left rotation on the given node
func (tree *RBTree) rotateLeft(pivotParent *RBNode) {
	if pivotParent == tree.nilNode || pivotParent.right == tree.nilNode {
		return
	}
	pivot := pivotParent.right
	pivotParent.right = pivot.left
	if pivot.left != tree.nilNode {
		pivot.left.parent = pivotParent
	}

	pivot.parent = pivotParent.parent
	if pivotParent.parent == nil {
		tree.root = pivot
	} else if pivotParent == pivotParent.parent.left {
		pivotParent.parent.left = pivot
	} else {
		pivotParent.parent.right = pivot
	}
	pivot.left = pivotParent
	pivotParent.parent = pivot
}

// RotateRight performs a right rotation on the given node
func (tree *RBTree) rotateRight(pivotParent *RBNode) {
	if pivotParent == tree.nilNode || pivotParent.left == tree.nilNode {
		return
	}
	pivot := pivotParent.left
	pivotParent.left = pivot.right
	if pivot.right != tree.nilNode {
		pivot.right.parent = pivotParent
	}

	pivot.parent = pivotParent.parent
	if pivotParent.parent == nil {
		tree.root = pivot
	} else if pivotParent == pivotParent.parent.right {
		pivotParent.parent.right = pivot
	} else {
		pivotParent.parent.left = pivot
	}
	pivot.right = pivotParent
	pivotParent.parent = pivot
}

// PrettyPrint prints the Red-Black Tree in a visually structured way
func (tree *RBTree) PrettyPrint(node *RBNode, level int) {
	if node == nil {
		node = tree.root
	}
	if node != tree.nilNode {
		// Recursively print the right subtree (to create top-down effect)
		if node.right != tree.nilNode {
			tree.PrettyPrint(node.right, level+1)
		}

		// Print the current node with proper indentation
		fmt.Printf("%s%d [%s]\n", strings.Repeat("        ", level), node.val, map[bool]string{true: "R", false: "B"}[node.red])

		// Recursively print the left subtree
		if node.left != tree.nilNode {
			tree.PrettyPrint(node.left, level+1)
		}
	}
}

func main() {
	rbt := NewRBTree()

	// Inserting values into the Red-Black Tree
	rbt.Insert(10)
	rbt.Insert(20)
	rbt.Insert(15)
	rbt.Insert(30)
	rbt.Insert(25)
	rbt.Insert(5)

	// Pretty print the tree
	fmt.Println("Red-Black Tree Structure:")
	rbt.PrettyPrint(rbt.root, 0)
}
