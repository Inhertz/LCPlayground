package lenlongestsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in       string
	expected int
	msg      string
}

func TestLengthOfLongestSubstring(t *testing.T) {

	testCases := []TestCase{
		{in: "abba", expected: 2, msg: "The answer is ba, with the length of 2."},
		{in: "abcabcbb", expected: 3, msg: "Case 1 abcabcbb"},
		{in: "bbbbb", expected: 1, msg: "Case 2 bbbbb"},
		{in: "pwwkew", expected: 3, msg: "The answer is wke, with the length of 3."},
	}

	for _, tc := range testCases {
		result := lengthOfLongestSubstring(tc.in)
		assert.Equal(t, tc.expected, result, tc.msg)
	}

}
