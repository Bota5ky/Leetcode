package temp

func isSymmetric(root *TreeNode) bool {
	return isSym(root, root)
}

func isSym(a *TreeNode, b *TreeNode) bool {
	if a == nil {
		return b == nil
	}
	if b == nil {
		return a == nil
	}
	if a.Val != b.Val {
		return false
	}
	return isSym(a.Left, b.Right) && isSym(a.Right, b.Left)
}
