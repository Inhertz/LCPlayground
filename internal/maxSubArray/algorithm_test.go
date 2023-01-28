package maxsubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in       []int
	expected int
	message  string
}

func TestMaxSubArray(t *testing.T) {
	testCases := []TestCase{
		{in: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expected: 6, message: "error on first test case"},
		{in: []int{1}, expected: 1, message: "error on second test case"},
		{in: []int{5, 4, -1, 7, 8}, expected: 23, message: "error on third test case"},
	}

	for _, tc := range testCases {
		result := maxSubArray(tc.in)
		assert.Equal(t, tc.expected, result, tc.message)
	}
}
