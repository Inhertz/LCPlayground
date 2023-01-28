package twosum

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	var dif int
	for i := range nums {
		dif = target - nums[i]
		if index, exist := hash[dif]; exist {
			var intSlice = []int{index, i}
			return intSlice
		}
		hash[nums[i]] = i
	}
	return nil
}
