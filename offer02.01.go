package temp

func removeDuplicateNodes(head *ListNode) *ListNode {
	slow := &ListNode{Next: head}
	ret := slow.Next
	cnt := make(map[int]int)
	for head != nil {
		cnt[head.Val]++
		if cnt[head.Val] > 1 {
			slow.Next = head.Next
			head = head.Next
		} else {
			slow = head
			head = head.Next
		}
	}
	return ret
}
