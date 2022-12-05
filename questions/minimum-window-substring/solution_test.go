package advantage_shuffle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minWindow(t *testing.T) {

	assert.Equal(t, "a", minWindow("a", "a"))

	assert.Equal(t, "", minWindow("a", "aa"))

	assert.Equal(t, "b", minWindow("ab", "b"))

	assert.Equal(t, "BANC", minWindow("ADOBECODEBANC", "ABC"))
}
