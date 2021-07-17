package leetcode

import (
	"fmt"
	"testing"
)

func Test_ConstructBinaryTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	r := constructBinaryTree(preorder, inorder)
	fmt.Println(r)
}
