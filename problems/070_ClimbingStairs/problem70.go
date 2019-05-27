package problem70

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	p, res := 1, 2
	for i := 2; i < n; i++ {
		res, p = res+p, res
	}
	return res
}
