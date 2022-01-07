package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciSequence(t *testing.T) {
	sequence := FibonacciSequence(4)
	assert.Equal(t, sequence, 3)
}
