package validate_binary_search_tree

import (
	"algorithm-trials/questions/utils"
	"math"
)

func isValidBST(root *utils.TreeNode) bool {
	var isValid func(root *utils.TreeNode) bool

	pre := math.MinInt64

	isValid = func(node *utils.TreeNode) bool {
		if node == nil {
			return true
		}

		if !isValid(node.Left) {
			return false
		}

		if node.Val <= pre {
			return false
		}
		pre = node.Val

		return isValid(node.Right)
	}

	return isValid(root)
}

func isValidBSTByRreOrder(root *utils.TreeNode) bool {
	var stack []*utils.TreeNode

	pre := math.MinInt64
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		n := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if n.Val <= pre {
			return false
		}
		pre = n.Val

		if n.Right != nil {
			cur = n.Right
		}
	}

	return true
}
