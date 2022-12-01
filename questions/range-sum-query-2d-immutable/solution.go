package range_sum_query_2d_immutable

import "fmt"

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rows := len(matrix)
	columns := len(matrix[0])

	preMatrix := make([][]int, rows+1)
	preMatrix[0] = make([]int, columns+1)

	for r := 0; r < rows; r++ {
		preMatrix[r+1] = make([]int, columns+1)
		for c := 0; c < columns; c++ {
			preMatrix[r+1][c+1] = preMatrix[r][c+1] + preMatrix[r+1][c] + matrix[r][c] - preMatrix[r][c]
		}
	}

	return NumMatrix{
		matrix: preMatrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.matrix[row2+1][col2+1] - this.matrix[row2+1][col1] - this.matrix[row1][col2+1] + this.matrix[row1][col1]
}

func (this *NumMatrix) debugLog() {
	for r := 0; r < len(this.matrix); r++ {
		for c := 0; c < len(this.matrix[0]); c++ {
			fmt.Printf("%3d", this.matrix[r][c])
		}
		fmt.Println()
	}
}
