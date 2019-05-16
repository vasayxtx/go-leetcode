package problem142

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	p1 := head
	p2 := head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 { // We have a cycle
			begin := head
			for begin != p1 {
				begin = begin.Next
				p1 = p1.Next
			}
			return begin
		}

	}
	return nil
}
