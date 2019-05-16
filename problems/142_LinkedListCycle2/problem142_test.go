package problem142

import (
	"reflect"
	"testing"
)

func makeLinkedList(length int, cycleIndex int) (*ListNode, *ListNode) {
	if length == 0 {
		return nil, nil
	}
	var nodes []*ListNode
	for i := 0; i < length; i++ {
		nodes = append(nodes, &ListNode{Val: i})
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	var cycleBegin *ListNode
	if cycleIndex != -1 {
		nodes[len(nodes)-1].Next = nodes[cycleIndex]
		cycleBegin = nodes[cycleIndex]
	}
	return nodes[0], cycleBegin
}

type args struct {
	head *ListNode
}

type test struct {
	name string
	args args
	want *ListNode
}

func makeTest(name string, listLen int, cycleIndex int) test {
	head, beginCycle := makeLinkedList(listLen, cycleIndex)
	return test{
		name: name,
		args: args{head},
		want: beginCycle,
	}
}

func Test_detectCycle(t *testing.T) {
	tests := []test{
		makeTest("Test #1", 0, -1),
		makeTest("Test #2", 1, -1),
		makeTest("Test #3", 10, -1),
		makeTest("Test #4", 1, 0),
		makeTest("Test #5", 2, 0),
		makeTest("Test #6", 100, 0),
		makeTest("Test #7", 10000, 5000),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
