package temp

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var a *ListNode
	c := head.Next
	for head != nil {
		head.Next = a
		a = head
		head = c
		if c != nil {
			c = c.Next
		}
	}
	return a
}
