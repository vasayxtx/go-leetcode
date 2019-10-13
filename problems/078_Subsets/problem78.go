package problem78

func findAllSubsets(res [][]int, nums []int, start int, neededLen int, subset []int, index int) [][]int {
	if index == neededLen {
		copiedSubset := make([]int, neededLen)
		copy(copiedSubset, subset)
		return append(res, copiedSubset)
	}
	for i := start; i < len(nums); i++ {
		subset[index] = nums[i]
		res = findAllSubsets(res, nums, i+1, neededLen, subset, index+1)
	}
	return res
}

func subsets(nums []int) [][]int {
	var res [][]int
	for i := 0; i <= len(nums); i++ {
		subset := make([]int, i)
		res = findAllSubsets(res, nums, 0, i, subset, 0)
	}
	return res
}
