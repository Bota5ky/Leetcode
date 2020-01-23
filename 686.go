package temp

import "strings"

//repeated-string-match
func repeatedStringMatch(A string, B string) int {
	maxlen := len(A)*2 + len(B)
	cnt := 1
	C := A
	for len(C) <= maxlen {
		if strings.Contains(C, B) {
			return cnt
		} else {
			C += A
			cnt++
		}
	}
	return -1
}
