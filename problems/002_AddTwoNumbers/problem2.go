package problem2

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeNode(val int) (*ListNode, int) {
	return &ListNode{Val: val % 10}, val / 10
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res, nextDigit := makeNode(l1.Val + l2.Val)
	node := res
	l1 = l1.Next
	l2 = l2.Next
	for l1 != nil || l2 != nil {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		node.Next, nextDigit = makeNode(nextDigit + v1 + v2)
		node = node.Next
	}
	if nextDigit != 0 {
		node.Next = &ListNode{Val: nextDigit}
	}
	return res
}
