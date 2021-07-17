package leetcode

import (
	"fmt"
	"testing"
)

func Test_Constructor(t *testing.T) {
	cq := Constructor()
	cq.AppendTail(3)
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
}
