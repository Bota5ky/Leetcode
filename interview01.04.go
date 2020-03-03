package temp

func canPermutePalindrome(s string) bool {
	numOfOdd := 0
	m := make(map[rune]bool)
	for _, val := range s {
		if m[val] == false {
			m[val] = true
			numOfOdd++
		} else {
			m[val] = false
			numOfOdd--
		}
	}
	return numOfOdd <= 1
}
