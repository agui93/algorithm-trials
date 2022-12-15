package binary_tree_preorder_traversal

import "algorithm-trials/questions/utils"

func preorderTraversal(root *utils.TreeNode) []int {

	r := make([]int, 0)

	var pre func(root *utils.TreeNode)
	pre = func(root *utils.TreeNode) {
		if root == nil {
			return
		}

		r = append(r, root.Val)

		pre(root.Left)
		pre(root.Right)
	}

	pre(root)

	return r
}
