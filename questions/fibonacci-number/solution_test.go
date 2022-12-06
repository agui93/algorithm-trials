package fibonacci_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fib(t *testing.T) {
	assert.Equal(t, 3, fib(4))
}
