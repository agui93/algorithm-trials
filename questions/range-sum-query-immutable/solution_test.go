package range_sum_query_immutable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumArray_SumRange(t *testing.T) {
	na := Constructor([]int{-2, 0, 3, -5, 2, -1})
	assert.Equal(t, 1, na.SumRange(0, 2))
	assert.Equal(t, -1, na.SumRange(2, 5))
	assert.Equal(t, -3, na.SumRange(0, 5))
}
