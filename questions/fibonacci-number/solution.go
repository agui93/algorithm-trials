package fibonacci_number

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	pre, curr := 0, 1
	tmp := 0

	for i := 2; i <= n; i++ {
		tmp = curr
		curr = pre + curr
		pre = tmp
	}

	return curr
}
