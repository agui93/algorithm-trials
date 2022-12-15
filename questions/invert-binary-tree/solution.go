package invert_binary_tree

import "algorithm-trials/questions/utils"

func invertTree(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	//left
	invertTree(root.Left)

	//right
	invertTree(root.Right)

	root.Left, root.Right = root.Right, root.Left

	return root
}
