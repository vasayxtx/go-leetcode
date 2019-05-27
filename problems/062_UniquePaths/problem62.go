package problem62

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			left := 0
			if j > 0 {
				left = grid[i][j-1]
			}
			top := 0
			if i > 0 {
				top = grid[i-1][j]
			}
			grid[i][j] = maxInt(1, left + top)
		}
	}
	return grid[n-1][m-1]
}
