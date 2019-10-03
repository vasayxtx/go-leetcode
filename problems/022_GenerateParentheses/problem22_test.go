package problem22

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

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test #1",
			args: args{n: 1},
			want: []string{"()"},
		},
		{
			name: "Test #2",
			args: args{n: 2},
			want: []string{"()()", "(())"},
		},
		{
			name: "Test #3",
			args: args{n: 3},
			want: []string{"((()))", "(())()", "()(())", "()()()", "(()())"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !checkSlicesEqualWithoutOrder(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
