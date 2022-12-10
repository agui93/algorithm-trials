package longest_increasing_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	assert.Equal(t, 4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
