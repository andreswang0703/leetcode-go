package binarytree

import (
	"testing"
)

func TestParse(t *testing.T) {
	a, b, c := 1, 2, 3
	input := []*int{&a, &b, &c}

	root := Parse(input)

	leftChild := &TreeNode{2, nil, nil}
	rightChild := &TreeNode{3, nil, nil}
	expectedRoot := &TreeNode{1, leftChild, rightChild}

	if !IsEqual(root, expectedRoot) {
		t.Error("Failed parsing, expected:", expectedRoot, " actual:", root)
	}
}

func TestParseNilNode(t *testing.T) {
	a, b := 1, 3
	input := []*int{&a, nil, &b}

	root := Parse(input)

	rightChild := &TreeNode{3, nil, nil}
	expectedRoot := &TreeNode{1, nil, rightChild}

	if !IsEqual(root, expectedRoot) {
		t.Error("Failed parsing, expected:", expectedRoot, " actual:", root)
	}
}
