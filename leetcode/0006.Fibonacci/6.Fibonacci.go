package leetcode

func fibonacci() func() int {
	back1, back2 := 0, 1
	return func() int {
		back1, back2 = back2%1000000007, back1%1000000007+back2%1000000007
		return back2 % 1000000007
	}
}
func fib(n int) int {
	f := fibonacci()
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	for i := 2; i < n; i++ {
		f()
	}
	return f()
}
