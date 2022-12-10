package minimum_falling_path_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minFallingPathSum(t *testing.T) {
	assert.Equal(t, 13, minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
}
