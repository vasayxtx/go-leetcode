package problem98

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test #1",
			args: args{},
			want: true,
		},
		{
			name: "Test #2",
			args: args{&TreeNode{Val: 1}},
			want: true,
		},
		{
			name: "Test #3",
			args: args{&TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			}},
			want: true,
		},
		{
			name: "Test #4",
			args: args{&TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 1},
			}},
			want: false,
		},
		{
			name: "Test #5",
			args: args{&TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			}},
			want: false,
		},
		{
			name: "Test #6",
			args: args{&TreeNode{
				Val: 34,
				Left: &TreeNode{
					Val:  -6,
					Left: &TreeNode{Val: -21},
				},
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.node); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
