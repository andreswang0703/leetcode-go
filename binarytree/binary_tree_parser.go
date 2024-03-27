package binarytree

func Parse(input []*int) *TreeNode {
	nodes := make([]*TreeNode, len(input))

	root := TreeNode{
		*input[0], nil, nil,
	}
	nodes[0] = &root

	for i, node := range input {
		if node == nil || i == 0 {
			continue
		}
		parentIdx := (i - 1) / 2
		parentNode := nodes[parentIdx]

		newNode := TreeNode{
			*input[i], nil, nil,
		}
		nodes[i] = &newNode

		if i%2 == 1 {
			parentNode.Left = &newNode
		} else {
			parentNode.Right = &newNode
		}
	}

	return &root
}
