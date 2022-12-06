package coin_change

/**
判断是动态规划问题

dp[n]定义: 当amount=n时, 需要至少dp[n]个硬币

状态转移方程
dp[n] = -1 , n<0
dp[n] = 0, n=0
dp[n] = min{ dp[n-coin]+1 | coin ∈ coins }, n>0
*/

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}

	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i-coin >= 0 {
				if 1+dp[i-coin] < dp[i] {
					dp[i] = 1 + dp[i-coin]
				}
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
