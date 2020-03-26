package leetcode

//https://leetcode-cn.com/problems/basic-calculator-ii/
func calculate(s string) int {
	var nums []int
	var sign []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			sign = append(sign, s[i])
			continue
		}
		temp := 0
		for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
			temp = temp*10 + int(s[i]-'0')
		}
		i--
		nums = append(nums, temp)
		if len(sign) > 0 && len(nums) > 1 {
			if sign[len(sign)-1] == '*' {
				nums[len(nums)-2] *= nums[len(nums)-1]
				sign = sign[:len(sign)-1]
				nums = nums[:len(nums)-1]
			} else if sign[len(sign)-1] == '/' {
				nums[len(nums)-2] /= nums[len(nums)-1]
				sign = sign[:len(sign)-1]
				nums = nums[:len(nums)-1]
			}
		}
	}
	for len(sign) > 0 {
		if sign[0] == '+' {
			nums[1] = nums[0] + nums[1]
			sign = sign[1:]
			nums = nums[1:]
		} else if sign[0] == '-' {
			nums[1] = nums[0] - nums[1]
			sign = sign[1:]
			nums = nums[1:]
		}
	}
	return nums[0]
}
