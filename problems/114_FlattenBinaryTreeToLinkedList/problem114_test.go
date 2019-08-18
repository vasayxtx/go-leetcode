package problem114

import (
	"reflect"
	"testing"
)

func checkTreeHasAnyLeftNode(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left != nil {
		return true
	}
	return checkTreeHasAnyLeftNode(root.Right)
}

func convertListToSlice(root *TreeNode) []int {
	var res []int
	for node := root; node != nil; node = node.Right {
		res = append(res, node.Val)
	}
	return res
}

func Test_flatten(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			/*
				1
			*/
			name: "Test #1",
			args: args{&TreeNode{Val: 1}},
			want: []int{1},
		},
		{
			/*
			   1
			  /
			 2
			*/
			name: "Test #2",
			args: args{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}},
			want: []int{1, 2},
		},
		{
			/*
				    1
				   /
				  2
				 /
				3
			*/
			name: "Test #3",
			args: args{&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}},
			want: []int{1, 2, 3},
		},
		{
			/*
			   1
			  / \
			 2   3
			*/
			name: "Test #4",
			args: args{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}},
			want: []int{1, 2, 3},
		},
		{
			/*
				    1
				   / \
				  2   5
				 / \   \
				3   4   6
			*/
			name: "Test #5",
			args: args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}},
			}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.args.root)
			if checkTreeHasAnyLeftNode(tt.args.root) {
				t.Errorf("linked list has left node")
			}
			got := convertListToSlice(tt.args.root)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
