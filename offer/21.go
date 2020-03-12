package temp

func exchange(nums []int) []int {
	i := 0
	j := len(nums) - 1
	for i < j {
		for nums[i]%2 == 1 && i < j {
			i++
		}
		for nums[j]%2 == 0 && i < j {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
