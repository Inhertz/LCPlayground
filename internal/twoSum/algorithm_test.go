package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		input    []int
		target   int
		expected []int
		message  string
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}, "Test Case 1"},
		{[]int{3, 2, 4}, 6, []int{1, 2}, "Test Case 2"},
		{[]int{3, 3}, 6, []int{0, 1}, "Test Case 3"},
		{[]int{1, 2, 3, 4, 5}, 10, nil, "Test Case 4"},
	}

	for _, tc := range testCases {
		result := twoSum(tc.input, tc.target)
		assert.Equal(t, tc.expected, result, tc.message)
	}
}
