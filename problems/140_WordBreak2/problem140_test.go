package problem140

import (
	"testing"
)

func checkSlicesEqualWithoutOrder(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[string]int)
	for i := 0; i < len(s1); i++ {
		m[s1[i]] += 1
		m[s2[i]] -= 1
	}
	for _, c := range m {
		if c != 0 {
			return false
		}
	}
	return true
}

func TestWordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test #1",
			args: args{
				s:        "catsanddog",
				wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			},
			want: []string{
				"cats and dog",
				"cat sand dog",
			},
		},
		{
			name: "Test #2",
			args: args{
				s:        "pineapplepenapple",
				wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			},
			want: []string{
				"pine apple pen apple",
				"pineapple pen apple",
				"pine applepen apple",
			},
		},
		{
			name: "Test #3",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: []string{},
		},
		{
			name: "Test #4",
			args: args{
				s:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaba",
				wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreak(tt.args.s, tt.args.wordDict)
			if !checkSlicesEqualWithoutOrder(got, tt.want) {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
