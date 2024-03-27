package binarytree

// TreeNode /**
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return IsEqual(a.Left, b.Left) && IsEqual(a.Right, b.Right)
}

// FindNode is for searching the node in the tree given the val with DFS order and return the first matched node.
func FindNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	leftRes := FindNode(root.Left, val)
	rightRes := FindNode(root.Right, val)

	if leftRes != nil {
		return leftRes
	}
	if rightRes != nil {
		return rightRes
	}
	return nil
}
