package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructBinaryTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	l := len(preorder)

	root := &TreeNode{Val: preorder[0]}
	if l == 1 {
		return root
	}

	leftDistance := 0
	for _, v := range inorder {
		if root.Val == v {
			break
		}
		leftDistance++
	}

	if leftDistance > 0 {
		root.Left = constructBinaryTree(preorder[1:leftDistance+1], inorder[0:leftDistance])
	}

	if l-leftDistance-1 > 0 {
		root.Right = constructBinaryTree(preorder[leftDistance+1:], inorder[leftDistance+1:])
	}

	return root
}
