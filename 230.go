package leetcode

//https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
func kthSmallest(root *TreeNode, k int) int {
	var nums []int
	traverse(root, &nums)
	return nums[k-1]
}

func traverse(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	traverse(root.Left, nums)
	*nums = append(*nums, root.Val)
	traverse(root.Right, nums)
}
