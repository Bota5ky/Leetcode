package temp

import "strconv"

func compressString(S string) string {
	cnt := 1
	var str string
	for i := 1; i <= len(S); i++ {
		//输出前面的
		if i == len(S) || i != 0 && S[i-1] != S[i] {
			str += string(S[i-1]) + strconv.Itoa(cnt)
			cnt = 1
		} else {
			cnt++
		}
	}
	if len(str) >= len(S) {
		return S
	}
	return str
}
