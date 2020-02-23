package temp

import "strings"

func firstUniqChar(s string) byte {
	for i := 0; i < len(s); i++ {
		if strings.Count(s, string(s[i])) == 1 {
			return s[i]
		}
	}
	return ' '
}
