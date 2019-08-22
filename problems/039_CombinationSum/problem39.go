package problem39

import (
	"sort"
)

type Combination struct {
	Sum       int
	Numbers   []int
	LastIndex int
}

func combinationSum(candidates []int, target int) [][]int {
	sortedCandidates := sort.IntSlice(candidates)
	sortedCandidates.Sort()

	queue := []Combination{{}}
	var res [][]int
	var comb Combination
	for len(queue) > 0 {
		comb, queue = queue[0], queue[1:]
		for i := comb.LastIndex; i < len(sortedCandidates); i++ {
			c := sortedCandidates[i]
			sum := comb.Sum + c

			if sum > target {
				break
			}

			numbers := make([]int, len(comb.Numbers)+1)
			copy(numbers, comb.Numbers)
			numbers[len(numbers)-1] = c

			if sum == target {
				res = append(res, numbers)
				continue
			}

			queue = append(queue, Combination{Sum: sum, Numbers: numbers, LastIndex: i})
		}
	}
	return res
}
