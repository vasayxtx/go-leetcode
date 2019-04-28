package problem122

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	k := -1
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			if k == -1 {
				k = i - 1
			}
			continue
		}
		if k == -1 {
			continue
		}
		profit += prices[i-1] - prices[k]
		k = -1
	}
	if k != -1 {
		profit += prices[len(prices)-1] - prices[k]
	}
	return profit
}
