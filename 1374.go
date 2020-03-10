package temp

func generateTheString(n int) string {
	var tail, ret string
	if n%2 == 0 {
		n--
		tail = "b"
	}
	for n > 0 {
		ret += "a"
		n--
	}
	return ret + tail
}
