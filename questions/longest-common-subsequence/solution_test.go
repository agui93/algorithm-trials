package longest_common_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonSubsequence(t *testing.T) {
	assert.Equal(t, 3, longestCommonSubsequence("abcde", "ace"))

	assert.Equal(t, 3, longestCommonSubsequence("abc", "abc"))

	assert.Equal(t, 0, longestCommonSubsequence("abc", "def"))
}
