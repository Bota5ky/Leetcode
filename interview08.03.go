package temp

func findMagicIndex(nums []int) int {
	for c, v := range nums {
		if c == v {
			return c
		}
	}
	return -1
}
