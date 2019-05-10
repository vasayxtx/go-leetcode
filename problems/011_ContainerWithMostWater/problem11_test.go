package problem11

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test #1",
			args: args{[]int{1, 2}},
			want: 1,
		},
		{
			name: "Test #2",
			args: args{[]int{3, 2}},
			want: 2,
		},
		{
			name: "Test #3",
			args: args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			want: 49,
		},
		{
			name: "Test #4",
			args: args{[]int{2, 3, 4, 5, 18, 17, 6}},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.heights); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
