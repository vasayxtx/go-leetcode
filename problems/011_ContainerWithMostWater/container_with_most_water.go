package containerwithmostwater

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	res := 0
	for left < right {
		res = max(res, (right-left)*min(heights[left], heights[right]))
		if heights[left] < heights[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return res
}
