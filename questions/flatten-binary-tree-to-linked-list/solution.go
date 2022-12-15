package flatten_binary_tree_to_linked_list

import "algorithm-trials/questions/utils"

func flatten(root *utils.TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)

	flatten(root.Right)

	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
	} else if root.Left != nil {
		ltail := root.Left
		for ltail.Right != nil {
			ltail = ltail.Right
		}

		ltail.Right = root.Right

		root.Right = root.Left
		root.Left = nil
	}
}
