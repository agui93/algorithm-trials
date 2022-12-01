package range_sum_query_2d_immutable

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumMatrix_SumRegion(t *testing.T) {

	t.Run("debug", func(t *testing.T) {
		matrix := [][]int{
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
		}

		for r := 0; r < len(matrix); r++ {
			for c := 0; c < len(matrix[0]); c++ {
				fmt.Printf("%3d", matrix[r][c])
			}
			fmt.Println()
		}

		c := Constructor(matrix)
		fmt.Println()
		c.debugLog()
	})

	t.Run("test1", func(t *testing.T) {
		matrix := [][]int{
			{3, 0, 1, 4, 2},
			{5, 6, 3, 2, 1},
			{1, 2, 0, 1, 5},
			{4, 1, 0, 1, 7},
			{1, 0, 3, 0, 5},
		}

		c := Constructor(matrix)

		assert.Equal(t, 3, c.SumRegion(0, 0, 0, 0))

		assert.Equal(t, 8, c.SumRegion(2, 1, 4, 3))
		assert.Equal(t, 11, c.SumRegion(1, 1, 2, 2))
		assert.Equal(t, 12, c.SumRegion(1, 2, 2, 4))
	})

}
