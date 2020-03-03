package temp

import "strings"

func isUnique(astr string) bool {
	for i := 0; i < len(astr); i++ {
		if strings.Contains(astr[i+1:], string(astr[i])) {
			return false
		}
	}
	return true
}
