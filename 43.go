package leetcode

//https://leetcode-cn.com/problems/multiply-strings/
func multiply(num1 string, num2 string) string {
	//2位*5位 不超过2+5位
	num := make([]int, len(num1)+len(num2))
	for i := 0; i < len(num2); i++ {
		for j := 0; j < len(num1); j++ {
			num[i+j] += int(num2[len(num2)-1-i]-'0') * int(num1[len(num1)-1-j]-'0')
		}
	}
	pre := 0
	for i := 0; i < len(num); i++ {
		num[i] += pre
		pre = num[i] / 10
		num[i] %= 10
	}
	i := 0
	for i = len(num) - 1; i > 0; i-- {
		if num[i] != 0 {
			break
		}
	}
	var ret string
	for ; i >= 0; i-- {
		ret += string(num[i] + '0')
	}
	return ret
}
