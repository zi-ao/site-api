package utils

// Max 获取参数类表中最大的数字
func Max(nums ...int) int {
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	return maxNum
}

// Min 获取参数类表中最小的数字
func Min(nums ...int) int {
	minNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if minNum > nums[i] {
			minNum = nums[i]
		}
	}
	return minNum
}
