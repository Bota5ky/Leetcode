package temp

func printNumbers(n int) []int {
	max := 1
	for i := 0; i < n; i++ {
		max *= 10
	}
	ret := make([]int, max-1)
	for i := 0; i < max-1; i++ {
		ret[i] = i + 1
	}
	return ret
}
