package problem139

import "testing"

func test(t *testing.T, s string, wordDict []string, expRes bool) {
	actRes := wordBreak(s, wordDict)
	if actRes != expRes {
		t.Errorf("fail for %q, expected %v, actual %v", s, expRes, actRes)
	}
}

func TestWordBreak(t *testing.T) {
	test(t, "leetcode", []string{"leet", "code"}, true)
	test(t, "applepenapple", []string{"apple", "pen"}, true)
	test(t, "catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false)
	test(
		t,
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
		[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
		false,
	)
}
