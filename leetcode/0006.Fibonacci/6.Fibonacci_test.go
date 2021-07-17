package leetcode

import (
	"fmt"
	"testing"
)

func Test_fib(t *testing.T) {
	for i := 0; i < 49; i++ {
		fmt.Printf("fib(%d) = %d\n", i, fib(i))
	}
}
