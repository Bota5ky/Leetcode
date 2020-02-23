package temp

func kthToLast(head *ListNode, k int) int {
	fast := head
	for k > 0 && fast != nil {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		head = head.Next
	}
	return head.Val
}
