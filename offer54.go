package temp

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
	var nums []int
	list(root, &nums)
	return nums[len(nums)-k]
}

//中序遍历
func list(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	list(root.Left, nums)
	*nums = append(*nums, root.Val)
	list(root.Right, nums)
}