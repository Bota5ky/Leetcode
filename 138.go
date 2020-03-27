package leetcode

//https://leetcode-cn.com/problems/copy-list-with-random-pointer/
func copyRandomList(head *Node2) *Node2 {
	cur := head
	for cur != nil {
		newnode := &Node2{Val: cur.Val, Next: cur.Next}
		cur.Next = newnode
		cur = cur.Next.Next
	}
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	if head != nil {
		cur = head.Next
	}
	ret := cur
	for cur != nil && cur.Next != nil {
		cur.Next = cur.Next.Next
		cur = cur.Next
	}
	return ret
}

//Node2 Node
type Node2 struct {
	Val    int
	Next   *Node2
	Random *Node2
}
