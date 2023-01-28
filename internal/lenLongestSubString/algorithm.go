package lenlongestsubstring

import "math"

func lengthOfLongestSubstring(s string) int {

	var head int
	var tail int
	windowElements := make(map[byte]int)
	windowSize := 0

	for head = 0; head < len(s); head++ {
		if _, exists := windowElements[s[head]]; exists {
			tail = int(math.Max(float64(tail), float64(windowElements[s[head]]+1)))
		}
		windowElements[s[head]] = head
		windowSize = int(math.Max(float64(windowSize), float64(1+head-tail)))
	}

	return windowSize
}
