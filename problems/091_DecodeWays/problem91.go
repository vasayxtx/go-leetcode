package problem91

func computeDecodingsNumber(s string, cache map[string]int) int {
	if s == "" {
		return 1
	}
	if s[0] == '0' {
		return 0
	}

	n, ok := cache[s]
	if ok {
		return n
	}

	res := computeDecodingsNumber(s[1:], cache)
	if len(s) > 1 && (s[0] == '1' || (s[0] == '2' && s[1] <= '6')) {
		res += computeDecodingsNumber(s[2:], cache)
		delete(cache, s[2:])
	}
	cache[s] = res
	return res
}

func numDecodings(s string) int {
	if s == "" {
		return 0
	}
	cache := make(map[string]int)
	return computeDecodingsNumber(s, cache)
}
