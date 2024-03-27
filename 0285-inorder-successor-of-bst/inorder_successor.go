package _285_inorder_successor_of_bst

import (
	"github.com/andreswang0703/leetcode-go/binarytree"
)

type TreeNode = binarytree.TreeNode

var prev *TreeNode
var res *TreeNode

// InorderSuccessor #285 Inorder Successor of BST
// This solution is generic for binary tree.
func InorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// two scenarios, one is p has a right child
	if p.Right != nil {
		successor := p.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		return successor
	}

	// second scenario, it doesn't have a right child
	dfs(root, p)
	return res
}

// keeps track of the previous node while doing inorder traversal
func dfs(root *TreeNode, p *TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Left, p)

	if prev == p {
		res = root
	}
	prev = root

	dfs(root.Right, p)
}
