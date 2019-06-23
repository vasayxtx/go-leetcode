package problem15

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sNums := sort.IntSlice(nums)
	sNums.Sort()
	var res [][]int
	for i := 0; i < len(sNums)-2; i++ {
		if sNums[i] > 0 || (i > 0 && sNums[i] == sNums[i-1]) {
			continue
		}
		for l, r := i+1, len(sNums)-1; l < r; {
			s := sNums[i] + sNums[l] + sNums[r]
			if s > 0 {
				r -= 1
			} else if s < 0 {
				l += 1
			} else {
				res = append(res, []int{sNums[i], sNums[l], sNums[r]})
				for l < r && sNums[l+1] == sNums[l] {
					l += 1
				}
				for l < r && sNums[r-1] == sNums[r] {
					r -= 1
				}
				l += 1
				r -= 1
			}
		}
	}
	return res
}
