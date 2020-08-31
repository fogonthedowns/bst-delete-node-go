package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) (t *TreeNode) {
	if root == nil {
		return nil
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		// the node is a leaf
		if root.Left == nil && root.Right == nil {
			root = nil
			// not a leaf - has a right child
		} else if root.Right != nil {
			root.Val = successor(root)
			root.Right = deleteNode(root.Right, root.Val)
			// not a leaf - has a left child
		} else {
			root.Val = predecessor(root)
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	return root
}

// One step right, then left as far as you can
// after node - the next node or the smallest after the current
func successor(root *TreeNode) int {
	root = root.Right
	for root.Left != nil {
		root = root.Left

	}
	return root.Val
}

// One step left, and then right as far as you can go
// before node - the previous node, or largest before the current one
func predecessor(root *TreeNode) int {
	root = root.Left
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}
