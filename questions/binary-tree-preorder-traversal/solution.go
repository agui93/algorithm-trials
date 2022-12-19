package binary_tree_preorder_traversal

import (
	"algorithm-trials/questions/utils"
	"container/list"
)

/**

二叉树的遍历
https://leetcode.cn/problems/binary-tree-preorder-traversal/solutions/87526/leetcodesuan-fa-xiu-lian-dong-hua-yan-shi-xbian-2/

*/

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

// 非递归
func preorderTraversalByStack(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := list.New()
	stack.PushFront(root)

	r := make([]int, 0)

	for stack.Len() > 0 {
		c := stack.Remove(stack.Front()).(*utils.TreeNode)

		r = append(r, c.Val)
		if c.Right != nil {
			stack.PushFront(c.Right)
		}
		if c.Left != nil {
			stack.PushFront(c.Left)
		}
	}

	return r
}
