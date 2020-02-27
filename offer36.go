package temp

//未测试
var head *TreeNode     //定义链表当前结点
var realHead *TreeNode //定义链表头部的结点
//中序递归遍历修改链表指针即可实现
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	node := treeToDoublyList(root.Left) //左
	if head == nil {                    //根
		head = node
		realHead = node
	} else {
		head.Right = node
		node.Left = head
		head = node
	}
	node = treeToDoublyList(root.Right) //右
	return realHead
}
