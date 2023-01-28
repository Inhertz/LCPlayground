package fib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	mem = make(map[int]int)
	// Test case with input 0
	result := fib(0)
	assert.Equal(t, 0, result, "Expected output for input 0 to be 0")

	// Test case with input 1
	result = fib(1)
	assert.Equal(t, 1, result, "Expected output for input 1 to be 1")

	// Test case with input 10
	result = fib(10)
	assert.Equal(t, 55, result, "Expected output for input 10 to be 55")

	// Test case with input 20
	result = fib(20)
	assert.Equal(t, 6765, result, "Expected output for input 20 to be 6765")
}

func TestFibTwo(t *testing.T) {
	mem = make(map[int]int)
	// Test case with input 0
	result := fibTwo(0)
	assert.Equal(t, 0, result, "Expected output for input 0 to be 0")

	// Test case with input 1
	result = fibTwo(1)
	assert.Equal(t, 1, result, "Expected output for input 1 to be 1")

	// Test case with input 10
	result = fibTwo(10)
	assert.Equal(t, 55, result, "Expected output for input 10 to be 55")

	// Test case with input 20
	result = fibTwo(20)
	assert.Equal(t, 6765, result, "Expected output for input 20 to be 6765")
}
