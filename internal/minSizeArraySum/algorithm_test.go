package minsizearraysum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	target   int
	nums     []int
	expected int
}

func TestMinSubArrayLen(t *testing.T) {

	testCases := []testCase{
		{target: 7, nums: []int{2, 3, 1, 2, 4, 3}, expected: 2},
		{target: 4, nums: []int{1, 4, 4}, expected: 1},
		{target: 11, nums: []int{1, 1, 1, 1, 1, 1, 1, 1}, expected: 0},
	}

	for _, tc := range testCases {

		result := minSubArrayLen(tc.target, tc.nums)
		assert.Equal(t, tc.expected, result, fmt.Sprintf("Expected : %d, Got: %d", tc.expected, result))
	}
}
