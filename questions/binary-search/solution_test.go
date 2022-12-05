package advantage_shuffle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_search(t *testing.T) {
	assert.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9))

	assert.Equal(t, 3, search([]int{-1, 3, 5, 9, 12}, 9))

	assert.Equal(t, 1, search([]int{5, 9, 12}, 9))

	assert.Equal(t, 0, search([]int{9, 12}, 9))

	assert.Equal(t, 0, search([]int{9}, 9))

	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 6))

	assert.Equal(t, -1, search([]int{-1, 3, 5, 9, 12}, 6))

}
