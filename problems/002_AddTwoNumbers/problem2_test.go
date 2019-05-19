package problem2

import (
	"reflect"
	"testing"
)

func makeList(d1 int, digits ...int) *ListNode {
	res := &ListNode{Val: d1}
	node := res
	for _, d := range digits {
		node.Next = &ListNode{Val: d}
		node = node.Next
	}
	return res
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test #1",
			args: args{
				l1: makeList(1),
				l2: makeList(2),
			},
			want: makeList(3),
		},
		{
			name: "Test #2",
			args: args{
				l1: makeList(1, 2, 3),
				l2: makeList(9),
			},
			want: makeList(0, 3, 3),
		},
		{
			name: "Test #3",
			args: args{
				l1: makeList(9),
				l2: makeList(1, 2, 3),
			},
			want: makeList(0, 3, 3),
		},
		{
			name: "Test #4",
			args: args{
				l1: makeList(2, 4, 3),
				l2: makeList(5, 6, 4),
			},
			want: makeList(7, 0, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
