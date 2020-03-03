package temp

func deleteNode2(node *ListNode) {
	*node = *node.Next
}

// func deleteNode(node *ListNode) {
//     node.Val=node.Next.Val
//     node.Next=node.Next.Next
// }
