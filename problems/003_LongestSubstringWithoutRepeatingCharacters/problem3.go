package problem3

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	positions := make([]int, 255)
	for i := 0; i < len(positions); i++ {
		positions[i] = -1
	}

	i := 0
	j := 0
	res := 0
	for i < len(s) && j < len(s) {
		char := s[j]
		oldPos := positions[char]
		positions[char] = j
		j += 1
		if i > oldPos {
			res = maxInt(res, j-i)
		} else {
			i = oldPos + 1
		}
	}
	return res
}
