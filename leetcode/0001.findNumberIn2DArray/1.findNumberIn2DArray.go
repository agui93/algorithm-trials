package leetcode

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	row := 0
	column := len(matrix[0]) - 1
	for row < len(matrix) && column >= 0 {
		if matrix[row][column] > target {
			column--
		} else if matrix[row][column] < target {
			row++
		} else {
			return true
		}
	}

	return false
}
