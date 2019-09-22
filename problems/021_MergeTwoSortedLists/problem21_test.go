package problem21

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
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
			args: args{nil, nil},
			want: nil,
		},
		{
			name: "Test #2",
			args: args{&ListNode{1, nil}, nil},
			want: &ListNode{1, nil},
		},
		{
			name: "Test #3",
			args: args{nil, &ListNode{1, nil}},
			want: &ListNode{1, nil},
		},
		{
			name: "Test #4",
			args: args{&ListNode{1, nil}, &ListNode{1, nil}},
			want: &ListNode{1, &ListNode{1, nil}},
		},
		{
			name: "Test #5",
			args: args{
				&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
				&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			},
			want: &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
