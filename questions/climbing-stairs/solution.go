package climbing_stairs

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	dp1, dp2 := 1, 2

	for i := 3; i <= n; i++ {
		dp1, dp2 = dp2, dp1+dp2
	}

	return dp2
}
