package minsizearraysum

import "math"

func minSubArrayLen(target int, nums []int) int {

	var head int
	var tail int
	windowSize := math.MaxInt
	var sum int

	for head = 0; head < len(nums); head++ {
		sum += nums[head]
		for sum >= target {
			windowSize = int(math.Min(float64(windowSize), float64(1+head-tail)))
			sum -= nums[tail]
			tail++
		}
	}

	if windowSize == math.MaxInt {
		windowSize = 0
	}

	return windowSize
}
