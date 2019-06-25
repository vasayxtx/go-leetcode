package problem322

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test #1",
			args: args{coins: nil, amount: 0},
			want: 0,
		},
		{
			name: "Test #2",
			args: args{coins: []int{2}, amount: 7},
			want: -1,
		},
		{
			name: "Test #3",
			args: args{coins: []int{2, 4, 6, 8}, amount: 7},
			want: -1,
		},
		{
			name: "Test #4",
			args: args{coins: []int{1}, amount: 7},
			want: 7,
		},
		{
			name: "Test #5",
			args: args{coins: []int{1, 2, 5}, amount: 11},
			want: 3,
		},
		{
			name: "Test #6",
			args: args{coins: []int{10, 5, 7, 1}, amount: 23},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
