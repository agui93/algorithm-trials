package longest_common_subsequence

/**
text1 和 text2 仅由小写英文字符组成

dp[i][j]含义: t1[0...i-1] 和 t2[0...j-1] 的最长公共子序列长度, 其中i,j不包含


dp[0][0] = 0
dp[0][1] = 0
dp[1][0] = 0
dp[i][j] = if t1[i] == t2[j]   dp[i-1][j-1]+1
		   else
				dp[i][j] = max{dp[i-1][j], dp[i][j-1] }
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	t1len := len(text1)
	t2len := len(text2)
	if t1len == 0 || t2len == 0 {
		return 0
	}

	dp := make([][]int, t1len+1)
	dp[0] = make([]int, t2len+1)

	for i := 1; i < t1len+1; i++ {
		dp[i] = make([]int, t2len+1)
		for j := 1; j < t2len+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i][j] < dp[i-1][j] {
					dp[i][j] = dp[i-1][j]
				}
				if dp[i][j] < dp[i][j-1] {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[t1len][t2len]
}
