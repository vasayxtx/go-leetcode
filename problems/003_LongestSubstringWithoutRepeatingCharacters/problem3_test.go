package problem3

import "testing"

func test(t *testing.T, s string, expRes int) {
	actRes := lengthOfLongestSubstring(s)
	if actRes != expRes {
		t.Errorf("Fail for %v, %d != %d", s, actRes, expRes)
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	test(t, "", 0)
	test(t, " ", 1)
	test(t, "a", 1)
	test(t, "ab", 2)
	test(t, "aa", 1)
	test(t, "abcabcbb", 3)
	test(t, "pwwkew", 3)
	test(t, "abaab", 2)
	test(t, "aab", 2)

}
