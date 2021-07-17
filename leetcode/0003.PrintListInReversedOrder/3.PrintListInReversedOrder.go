package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	l := 0
	s := head
	for s != nil {
		l++
		s = s.Next
	}

	t := make([]int, l)
	s = head
	for s != nil {
		t[l-1] = s.Val
		l--
		s = s.Next
	}

	return t
}
