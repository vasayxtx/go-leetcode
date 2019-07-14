package problem198

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return maxInt(nums[0], nums[1])
	}
	a1 := nums[0]
	a2 := nums[1]
	a3 := nums[2] + a1
	for i := 3; i < len(nums); i++ {
		tmp := nums[i] + maxInt(a1, a2)
		a2, a1 = a3, a2
		a3 = tmp
	}
	return maxInt(a3, a2)
}
