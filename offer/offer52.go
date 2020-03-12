package temp

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB
	for a != nil || b != nil {
		if a == nil {
			a = headB
			b = b.Next
			continue
		}
		if b == nil {
			b = headA
			a = a.Next
			continue
		}
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next
	}
	return nil
}