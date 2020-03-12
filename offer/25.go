package temp

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	node := ret
	for l1 != nil || l2 != nil {
		if l2 == nil || (l1 != nil && l1.Val < l2.Val) {
			node.Next = l1
			l1 = l1.Next
			node = node.Next
		} else {
			node.Next = l2
			l2 = l2.Next
			node = node.Next
		}
	}
	return ret.Next
}
