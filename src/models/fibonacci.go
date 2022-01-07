package models

import (
	"sync"

	"github.com/paingha/pex/src/lib"
)

// FibonacciDataStore read and write protected data store.
type FibonacciDataStore struct {
	Counter         int
	FibonacciNumber int
	sync.RWMutex
}

// FibonacciService interface for fibonacci model.
// NB: using an interface so the FibonacciDataStore can be swapped
// to using a database and have similar method signatures.
type FibonacciService interface {
	Next() int
	Previous() int
	Current() int
}

// NewFibonacciDataStore creates a new fibonacci data store.
func NewFibonacciDataStore() *FibonacciDataStore {
	return &FibonacciDataStore{
		Counter:         0,
		FibonacciNumber: 0,
	}
}

// Next increments the fibonacci number to the next number in the sequence.
func (f *FibonacciDataStore) Next() int {
	f.Lock()
	defer f.Unlock()
	f.Counter = f.Counter + 1
	next := lib.FibonacciSequence(f.Counter)
	f.FibonacciNumber = next
	return f.FibonacciNumber
}

// Previous decrements the fibonacci number to the previous number in the sequence.
func (f *FibonacciDataStore) Previous() int {
	f.Lock()
	defer f.Unlock()
	if f.Counter >= 0 {
		f.Counter = f.Counter - 1
	}
	f.FibonacciNumber = lib.FibonacciSequence(f.Counter)
	return f.FibonacciNumber
}

// Current returns the current number in the fibonacci sequence.
func (f *FibonacciDataStore) Current() int {
	f.RLock()
	defer f.RUnlock()
	return f.FibonacciNumber
}
