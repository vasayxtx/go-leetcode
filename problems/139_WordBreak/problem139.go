package problem139

import (
	"strings"
)

func wordBreakWithCaching(s string, wordDict []string, cache map[string]bool) bool {
	if s == "" {
		return true
	}
	res, ok := cache[s]
	if ok {
		return res
	}
	for _, word := range wordDict {
		if !strings.HasPrefix(s, word) {
			continue
		}
		res = wordBreakWithCaching(s[len(word):], wordDict, cache)
		cache[s] = res
		if res {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	return wordBreakWithCaching(s, wordDict, map[string]bool{})
}
