package problem124

import "testing"

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test #1",
			args: args{&TreeNode{Val: 42}},
			want: 42,
		},
		{
			name: "Test #2",
			args: args{&TreeNode{
				Val:   42,
				Left:  &TreeNode{Val: -1},
				Right: &TreeNode{Val: -2},
			}},
			want: 42,
		},
		{
			name: "Test #3",
			args: args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			}},
			want: 6,
		},
		{
			name: "Test #4",
			args: args{&TreeNode{
				Val:  -10,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			}},
			want: 42,
		},
		{
			name: "Test #5",
			args: args{&TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val:  8,
					Left: &TreeNode{Val: 13},
					Right: &TreeNode{
						Val:   4,
						Right: &TreeNode{Val: 1},
					},
				},
			}},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
