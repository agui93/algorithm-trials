package longest_increasing_subsequence

/**

dp[i]: 在nums[0...i]中，包含nums[i]结尾时的最长递增子序列长度

dp[0] = 1
dp[i] = 1 + max{ dp[k]; k ∈ (0<=k<i) && num[k] < num[i]  }

*/

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for k := 0; k < i; k++ {
			if nums[k] < nums[i] && dp[i] < dp[k]+1 {
				dp[i] = dp[k] + 1
			}
		}
	}

	max := dp[0]
	for i := 1; i < len(dp); i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}

	return max
}
