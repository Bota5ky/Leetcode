package temp

//654321后序遍历
func flatten(root *TreeNode) {
	var pre *TreeNode
	connect(root, &pre)
}
func connect(root *TreeNode, pre **TreeNode) {
	if root == nil {
		return
	}
	connect(root.Right, pre)
	connect(root.Left, pre)
	root.Left = nil
	root.Right = (*pre)
	*pre = root
}
