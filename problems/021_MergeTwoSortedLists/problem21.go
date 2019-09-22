package problem21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	var resHead *ListNode
	var resTail *ListNode
	for it1, it2 := l1, l2; it1 != nil || it2 != nil; {
		node := &ListNode{}
		if it1 != nil && (it2 == nil || it1.Val <= it2.Val) {
			node.Val = it1.Val
			it1 = it1.Next
		} else {
			node.Val = it2.Val
			it2 = it2.Next
		}
		if resHead == nil {
			resHead = node
		}
		if resTail != nil {
			resTail.Next = node
		}
		resTail = node
	}
	return resHead
}
