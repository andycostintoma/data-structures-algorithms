package main

import (
	"fmt"
)

// BSTNode represents a node in the binary search tree
type BSTNode struct {
	left  *BSTNode
	right *BSTNode
	val   *int
}

// NewBSTNode creates a new BSTNode, optionally with a value
func NewBSTNode(val *int) *BSTNode {
	return &BSTNode{val: val}
}

// Insert adds a value to the BST
func (n *BSTNode) Insert(val int) {
	if n.val == nil {
		n.val = &val
		return
	}

	// prevent inserting duplicate values
	if *n.val == val {
		return
	}

	if val < *n.val {
		if n.left != nil {
			n.left.Insert(val)
		} else {
			n.left = NewBSTNode(&val)
		}
	} else {
		if n.right != nil {
			n.right.Insert(val)
		} else {
			n.right = NewBSTNode(&val)
		}
	}
}

// Delete removes a value from the BST and returns the modified subtree
func (n *BSTNode) Delete(val int) *BSTNode {
	if n == nil || n.val == nil {
		return nil
	}

	if val < *n.val {
		if n.left != nil {
			n.left = n.left.Delete(val)
		}
		return n
	} else if val > *n.val {
		if n.right != nil {
			n.right = n.right.Delete(val)
		}
		return n
	} else {
		// Node found
		if n.right == nil {
			return n.left
		}
		if n.left == nil {
			return n.right
		}

		// Node has two children: find min in right subtree
		minNode := n.right
		for minNode.left != nil {
			minNode = minNode.left
		}
		n.val = minNode.val
		n.right = n.right.Delete(*minNode.val)
		return n
	}
}

// GetMin returns the minimum value in the tree
func (n *BSTNode) GetMin() *int {
	if n == nil {
		return nil
	}
	if n.left == nil {
		return n.val
	}
	return n.left.GetMin()
}

// GetMax returns the maximum value in the tree
func (n *BSTNode) GetMax() *int {
	if n == nil {
		return nil
	}
	if n.right == nil {
		return n.val
	}
	return n.right.GetMax()
}

// Helper to print BST (for debugging)
func (n *BSTNode) InOrder() {
	if n == nil || n.val == nil {
		return
	}
	n.left.InOrder()
	fmt.Print(*n.val, " ")
	n.right.InOrder()
}
