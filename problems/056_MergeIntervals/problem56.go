package problem56

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > right {
			res = append(res, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
			continue
		}
		if intervals[i][1] > right {
			right = intervals[i][1]
		}
	}
	res = append(res, []int{left, right})
	return res
}
