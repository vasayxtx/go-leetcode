package problem103

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test #1",
			args: args{nil},
			want: nil,
		},
		{
			name: "Test #2",
			args: args{&TreeNode{Val: 1}},
			want: [][]int{{1}},
		},
		{
			name: "Test #3",
			args: args{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}},
			want: [][]int{{1}, {3, 2}},
		},
		{
			name: "Test #4",
			args: args{&TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
			}},
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			name: "Test #5",
			args: args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}},
			}},
			want: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
