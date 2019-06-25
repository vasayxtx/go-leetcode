package problem322

func coinChange(coins []int, amount int) int {
	counts := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		counts[i] = -1
	}
	counts[0] = 0
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if counts[i-coin] == -1 {
				continue
			}
			c := counts[i-coin] + 1
			if counts[i] == -1 || counts[i] > c {
				counts[i] = c
			}
		}
	}
	return counts[amount]
}
