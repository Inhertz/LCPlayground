package twosumii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	target   int
	nums     []int
	expected []int
}

func TestTwoSum(t *testing.T) {
	testCases := []testCase{
		{target: 9, nums: []int{2, 7, 11, 15}, expected: []int{1, 2}},
		{target: 6, nums: []int{2, 3, 4}, expected: []int{1, 3}},
		{target: -1, nums: []int{-1, 0}, expected: []int{1, 2}},
	}

	for _, tc := range testCases {

		result := twoSum(tc.nums, tc.target)
		assert.Equal(t, tc.expected, result, fmt.Sprintf("Expected : %d, Got: %d", tc.expected, result))
	}
}
