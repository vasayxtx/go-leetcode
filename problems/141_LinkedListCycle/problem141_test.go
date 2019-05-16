package problem141

import "testing"

func makeLinkedList(length int, cycleIndex int) *ListNode {
	if length == 0 {
		return nil
	}
	var nodes []*ListNode
	for i := 0; i < length; i++ {
		nodes = append(nodes, &ListNode{Val: i})
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	if cycleIndex != -1 {
		nodes[len(nodes)-1].Next = nodes[cycleIndex]
	}
	return nodes[0]
}

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test #1",
			args: args{nil},
			want: false,
		},
		{
			name: "Test #2",
			args: args{&ListNode{}},
			want: false,
		},
		{
			name: "Test #3",
			args: args{makeLinkedList(10, -1)},
			want: false,
		},
		{
			name: "Test #4",
			args: args{makeLinkedList(1, 0)},
			want: true,
		},
		{
			name: "Test #5",
			args: args{makeLinkedList(2, 0)},
			want: true,
		},
		{
			name: "Test #6",
			args: args{makeLinkedList(100, 0)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
