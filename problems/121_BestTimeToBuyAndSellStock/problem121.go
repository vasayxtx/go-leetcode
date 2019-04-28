package problem121

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	p := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > p {
			maxProfit = maxInt(maxProfit, prices[i]-p)
			continue
		}
		p = prices[i]
	}
	return maxProfit
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
