package coin_change_ii

/**

dp[i][j]含义: 当coins[0...i]参与，可以凑成(总金额==j)的硬币组合数

dp[i][j] = dp[i][j-coins[i]]  //当第i个硬币参与
			+ dp[i-1][j]	  //当第i个硬币不参与时

*/

func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	if len(coins) == 0 {
		return 0
	}

	dp := make([][]int, len(coins)+1)
	dp[0] = make([]int, amount+1)
	dp[0][0] = 1

	for i := 1; i <= len(coins); i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1

		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				dp[i][j] = dp[i][j-coins[i-1]] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}

	}

	return dp[len(coins)][amount]
}
