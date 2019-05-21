package problem91

func numDecodings(s string) int {
	if s == "" || s[0] == '0' {
		return 0
	}
	nums := make([]int, len(s))
	nums[0] = 1
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
			nums[i] = nums[i-1]
		}
		if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
			if i >= 2 {
				nums[i] += nums[i-2]
			} else {
				nums[i] += 1
			}
		}
	}
	return nums[len(s)-1]
}
