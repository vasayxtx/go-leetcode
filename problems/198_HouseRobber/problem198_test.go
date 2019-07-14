package problem198

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test #1",
			args: args{[]int{}},
			want: 0,
		},
		{
			name: "Test #2",
			args: args{[]int{7}},
			want: 7,
		},
		{
			name: "Test #3",
			args: args{[]int{5, 7}},
			want: 7,
		},
		{
			name: "Test #4",
			args: args{[]int{5, 11, 5}},
			want: 11,
		},
		{
			name: "Test #5",
			args: args{[]int{7, 11, 5}},
			want: 12,
		},
		{
			name: "Test #6",
			args: args{[]int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "Test #7",
			args: args{[]int{2, 7, 9, 3, 1}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
