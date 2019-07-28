package problem19

import (
	"reflect"
	"testing"
)

func makeLinkedListFromSlice(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	head := &ListNode{Val: s[0]}
	node := head
	for i := 1; i < len(s); i++ {
		node.Next = &ListNode{Val: s[i]}
		node = node.Next
	}
	return head
}

func makeSliceFromLinkedList(head *ListNode) []int {
	if head == nil {
		return nil
	}
	res := []int{head.Val}
	for node := head; node.Next != nil; node = node.Next {
		res = append(res, node.Next.Val)
	}
	return res
}

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test #1",
			args: args{head: makeLinkedListFromSlice([]int{1}), n: 1},
			want: nil,
		},
		{
			name: "Test #2",
			args: args{head: makeLinkedListFromSlice([]int{1, 2}), n: 1},
			want: []int{1},
		},
		{
			name: "Test #3",
			args: args{head: makeLinkedListFromSlice([]int{1, 2}), n: 2},
			want: []int{2},
		},
		{
			name: "Test #4",
			args: args{head: makeLinkedListFromSlice([]int{1, 2, 3, 4, 5}), n: 2},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "Test #5",
			args: args{head: makeLinkedListFromSlice([]int{1, 2, 3}), n: 1},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNthFromEnd(tt.args.head, tt.args.n)
			if !reflect.DeepEqual(makeSliceFromLinkedList(got), tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
