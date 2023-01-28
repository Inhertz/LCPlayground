package maxarea

import "math"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	result := math.MinInt

	for right > left {

		result = int(math.Max(float64(result), math.Min(float64(height[left]), float64(height[right]))*float64(right-left)))

		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}

	return result
}
