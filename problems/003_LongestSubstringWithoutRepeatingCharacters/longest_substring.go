package longestsubstring

func lengthOfLongestSubstring(s string) int {
	positions := make([]int, 255)
	for i := 0; i < len(positions); i++ {
		positions[i] = -1
	}

	res := 0
	curLen := 0
	for i := 0; i < len(s); i++ {
		char := s[i]
		oldPos := positions[char]
		if oldPos != -1 {
			if curLen > res {
				res = curLen
			}
			for j := 0; j < len(positions); j++ {
				if positions[j] >= 0 && positions[j] <= oldPos {
					positions[j] = -1
					curLen -= 1
				}
			}
		}

		positions[char] = i
		curLen += 1
	}

	if curLen > res {
		res = curLen
	}

	return res
}
