package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	matrix [][]int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one bool
}

func Test_Problem240(t *testing.T) {

	qs := []question{

		{
			para{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5},
			ans{true},
		},

		{
			para{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20},
			ans{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem ------------------------\n")

	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("【input】:%v       【output】:%v\n", p, findNumberIn2DArray(p.matrix, p.target))
	}

	fmt.Printf("\n\n\n")
}
