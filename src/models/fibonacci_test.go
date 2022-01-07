package models_test

import (
	"testing"

	"github.com/paingha/pex/src/models"
	"github.com/stretchr/testify/assert"
)

func TestNewFibonacciDataStore(t *testing.T) {
	newFibonacciDataStore := models.NewFibonacciDataStore()
	assert.NotEmpty(t, newFibonacciDataStore)
	assert.Equal(t, newFibonacciDataStore.FibonacciNumber, 1)
}

func TestNext(t *testing.T) {
	newFibonacciDataStore := models.NewFibonacciDataStore()
	assert.Equal(t, newFibonacciDataStore.FibonacciNumber, 1)
	nextFibonacciNumber := newFibonacciDataStore.Next()
	assert.Equal(t, nextFibonacciNumber, 1)
}

func TestPrevious(t *testing.T) {
	newFibonacciDataStore := models.NewFibonacciDataStore()
	assert.Equal(t, newFibonacciDataStore.FibonacciNumber, 1)
	nextFibonacciNumber := newFibonacciDataStore.Next()
	assert.Equal(t, nextFibonacciNumber, 1)
	previousFibonacciNumber := newFibonacciDataStore.Previous()
	assert.Equal(t, previousFibonacciNumber, 1)
}
