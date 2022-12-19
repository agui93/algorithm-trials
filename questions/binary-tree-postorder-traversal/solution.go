package binary_tree_postorder_traversal

import (
	"algorithm-trials/questions/utils"
	"container/list"
)

// 递归版二叉树后序遍历
func postorderTraversal(root *utils.TreeNode) []int {
	r := make([]int, 0)

	var postorder func(node *utils.TreeNode)

	postorder = func(node *utils.TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			postorder(node.Left)
		}
		if node.Right != nil {
			postorder(node.Right)
		}
		r = append(r, node.Val)
	}

	postorder(root)

	return r
}

// 非递归版二叉树后序遍历
func postorderTraversalByStack(root *utils.TreeNode) []int {
	r := make([]int, 0)
	if root == nil {
		return r
	}

	stack1 := list.New()
	stack2 := list.New()

	stack1.PushFront(root)

	for stack1.Len() > 0 {
		c := stack1.Remove(stack1.Front()).(*utils.TreeNode)
		stack2.PushFront(c)

		if c.Left != nil {
			stack1.PushFront(c.Left)
		}
		if c.Right != nil {
			stack1.PushFront(c.Right)
		}
	}

	for stack2.Len() > 0 {
		c := stack2.Remove(stack2.Front()).(*utils.TreeNode)
		r = append(r, c.Val)
	}

	return r
}
