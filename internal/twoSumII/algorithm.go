package twosumii

func twoSum(numbers []int, target int) []int {
	result := []int{0, len(numbers) - 1}

	for numbers[result[0]]+numbers[result[1]] != target {
		if numbers[result[0]]+numbers[result[1]] < target {
			result[0]++
		} else {
			result[1]--
		}
	}

	result[0]++
	result[1]++

	return result
}
