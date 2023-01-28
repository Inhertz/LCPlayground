package maxarea

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums     []int
	expected int
}

func TestMinSubArrayLen(t *testing.T) {

	testCases := []testCase{
		{nums: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, expected: 49},
		{nums: []int{1, 1}, expected: 1},
	}

	for _, tc := range testCases {

		result := maxArea(tc.nums)
		assert.Equal(t, tc.expected, result, fmt.Sprintf("Expected : %d, Got: %d", tc.expected, result))
	}
}
