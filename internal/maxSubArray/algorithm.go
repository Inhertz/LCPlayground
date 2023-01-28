package maxsubarray

import "math"

func maxSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	result := math.MinInt

	currSum := 0

	for i := 0; i < len(nums); i++ {
		currSum = int(math.Max(float64(nums[i]), float64(currSum+nums[i])))
		result = int(math.Max(float64(result), float64(currSum)))
	}

	return result
}
