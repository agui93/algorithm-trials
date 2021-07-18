package leetcode

import (
	"container/list"
	"fmt"
)

type node struct {
	r int
	c int
}

func sum(r1, c1 int) int {
	s := 0
	for r1 != 0 {
		s = s + r1%10
		r1 = r1 / 10
	}
	for c1 != 0 {
		s = s + c1%10
		c1 = c1 / 10
	}

	return s
}

func robotMovingCount(m int, n int, k int) int {
	if m <= 0 || n <= 0 || k < 0 {
		return 0
	}

	count := 0

	checked := make([][]bool, m)
	for i := 0; i < m; i++ {
		checked[i] = make([]bool, n)
	}

	stack := list.New()

	stack.PushBack(&node{0, 0})
	checked[0][0] = true
	for stack.Len() > 0 {
		count++
		e := stack.Remove(stack.Back()).(*node)

		fmt.Printf("%d %d %d\n", e.r, e.c, sum(e.r, e.c))

		if e.r-1 >= 0 && !checked[e.r-1][e.c] && sum(e.r-1, e.c) <= k {
			stack.PushBack(&node{e.r - 1, e.c})
			checked[e.r-1][e.c] = true
		}

		if e.r+1 < m && !checked[e.r+1][e.c] && sum(e.r+1, e.c) <= k {
			stack.PushBack(&node{e.r + 1, e.c})
			checked[e.r+1][e.c] = true
		}

		if e.c-1 >= 0 && !checked[e.r][e.c-1] && sum(e.r, e.c-1) <= k {
			stack.PushBack(&node{e.r, e.c - 1})
			checked[e.r][e.c-1] = true
		}

		if e.c+1 < n && !checked[e.r][e.c+1] && sum(e.r, e.c+1) <= k {
			stack.PushBack(&node{e.r, e.c + 1})
			checked[e.r][e.c+1] = true
		}
	}

	return count
}
