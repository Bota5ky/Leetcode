package leetcode

//https://leetcode-cn.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	ret[0] = 1
	for i := 0; i < len(nums)-1; i++ {
		ret[i+1] = ret[i] * nums[i]
	}
	temp := 1
	for i := len(nums) - 2; i >= 0; i-- {
		temp *= nums[i+1]
		ret[i] *= temp
	}
	return ret
}
