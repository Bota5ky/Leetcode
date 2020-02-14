package temp

func hammingWeight(num uint32) int {
	var cnt uint32
	for num != 0 {
		cnt += num & 1
		num >>= 1
	}
	return int(cnt)
}
