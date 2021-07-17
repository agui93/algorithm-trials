package leetcode

import (
	"fmt"
	"testing"
)

func Test_PrintListInReversedOrder(t *testing.T) {
	head := &ListNode{1, nil}
	s := head

	s.Next = &ListNode{3, nil}
	s = s.Next

	s.Next = &ListNode{2, nil}

	r := reversePrint(head)
	fmt.Println(r)
}
