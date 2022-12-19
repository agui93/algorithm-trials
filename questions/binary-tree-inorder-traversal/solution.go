package binary_tree_inorder_traversal

import (
	"algorithm-trials/questions/utils"
	"container/list"
)

// 递归的二叉树中序遍历
func inorderTraversal(root *utils.TreeNode) []int {
	r := make([]int, 0)

	var inorder func(node *utils.TreeNode)

	inorder = func(node *utils.TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		r = append(r, node.Val)
		inorder(node.Right)
	}

	inorder(root)

	return r
}

// 非递归的二叉树中序遍历
func inorderTraversalByStack(root *utils.TreeNode) []int {
	r := make([]int, 0)
	if root == nil {
		return r
	}

	stack := list.New()

	cur := root
	for cur != nil || stack.Len() > 0 {
		for cur != nil {
			stack.PushFront(cur)
			cur = cur.Left
		}

		n := stack.Remove(stack.Front()).(*utils.TreeNode)
		r = append(r, n.Val)

		if n.Right != nil {
			cur = n.Right
		}
	}

	return r
}

// 非递归的二叉树中序遍历
func inorderTraversalBySlice(root *utils.TreeNode) []int {
	r := make([]int, 0)
	if root == nil {
		return r
	}

	var stack []*utils.TreeNode

	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		n := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		r = append(r, n.Val)

		if n.Right != nil {
			cur = n.Right
		}
	}

	return r
}
