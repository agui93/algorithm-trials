package minimum_falling_path_sum

import "math"

/**
dp[r][c]定义: 第r行第c列的元素被选中时, 下降路径的最小和是dp[r][c]

dp[0][j] = 0; n <= j <= n
dp[i][j] = matrix[i][j] + min{ dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1] } ;
*/

func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	dp := make([][]int, len(matrix)+1)
	dp[0] = make([]int, len(matrix)+1)

	for r := 1; r <= len(matrix); r++ {
		dp[r] = make([]int, len(matrix)+1)
		for c := 1; c < len(matrix)+1; c++ {
			dp[r][c] = math.MaxInt64
		}
		for c := 1; c <= len(matrix); c++ {

			if dp[r][c] > dp[r-1][c]+matrix[r-1][c-1] {
				dp[r][c] = dp[r-1][c] + matrix[r-1][c-1]
			}

			if 1 < c {
				if dp[r][c] > dp[r-1][c-1]+matrix[r-1][c-1] {
					dp[r][c] = dp[r-1][c-1] + matrix[r-1][c-1]
				}
			}

			if c < len(matrix) {
				if dp[r][c] > dp[r-1][c+1]+matrix[r-1][c-1] {
					dp[r][c] = dp[r-1][c+1] + matrix[r-1][c-1]
				}
			}
		}
	}

	res := math.MaxInt64
	for c := 1; c < len(matrix)+1; c++ {

		if res > dp[len(matrix)][c] {
			res = dp[len(matrix)][c]
		}
	}

	if res == math.MaxInt64 {
		return -1
	}

	return res
}
