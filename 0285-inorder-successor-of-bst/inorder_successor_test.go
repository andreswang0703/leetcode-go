package _285_inorder_successor_of_bst

import (
	"github.com/andreswang0703/leetcode-go/binarytree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_inorder_successor(t *testing.T) {

	inputs := []struct {
		name     string    // test name
		root     *TreeNode // input tree root
		p        int       // p
		expected *int      // expected TreeNode's value
	}{
		{
			"case 1",
			binarytree.Parse([]*int{newInt(1), newInt(2), newInt(3), newInt(4), newInt(5), newInt(6), newInt(7)}),
			1,
			newInt(6),
		},
		{
			"case 2",
			binarytree.Parse([]*int{newInt(1), newInt(2), newInt(3), newInt(4), newInt(5), newInt(6), newInt(7)}),
			7,
			nil,
		},
		{
			"case 2",
			binarytree.Parse([]*int{newInt(1), newInt(2), newInt(3), newInt(4), newInt(5), newInt(6), newInt(7)}),
			6,
			newInt(3),
		},
	}

	for _, tt := range inputs {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t,
				getExpectedNode(tt.root, tt.expected),                         // expected node
				InorderSuccessor(tt.root, binarytree.FindNode(tt.root, tt.p)), // actual node
			)
		})
	}
}

func newInt(n int) *int {
	return &n
}

func getExpectedNode(root *TreeNode, e *int) *TreeNode {
	if e == nil {
		return nil
	} else {
		return binarytree.FindNode(root, *e)
	}
}
