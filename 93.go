package leetcode

import "strconv"

//https://leetcode-cn.com/problems/restore-ip-addresses/
func restoreIPAddresses(s string) []string {
	var pre string
	var ret []string
	addr(s, pre, 4, &ret)
	return ret
}

func addr(s string, pre string, n int, ret *[]string) { //pre表示前面累计的xxx.xxx.xx.
	if len(s) == 0 {
		return
	}
	if n == 1 {
		num, err := strconv.Atoi(s)
		if err == nil && num < 256 {
			if len(s) == 2 && num < 10 || len(s) == 3 && num < 100 || len(s) > 3 {
				return
			}
			pre += s
			*ret = append(*ret, pre)
		}
		return
	}
	for i := 1; i <= 3 && i <= len(s); i++ {
		num, err := strconv.Atoi(s[:i])
		if err == nil && num < 256 {
			if i == 2 && num < 10 || i == 3 && num < 100 {
				break
			}
			temp := pre + s[:i] + "."
			addr(s[i:], temp, n-1, ret)
		}
	}
}
