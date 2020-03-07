package temp

//摩尔投票法
func majorityElement2(nums []int) int {
	cnt := 0
	cmp := 0
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			cnt++
			cmp = nums[i]
		} else if cmp == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}
	if cnt > 0 {
		return cmp
	}
	return -1
}

// func majorityElement(nums []int) int {
//     m:=make(map[int]int)
//     for i:=0;i<len(nums);i++ {
//         m[nums[i]]++
//         if m[nums[i]]>len(nums)/2 {
//             return nums[i]
//         }
//     }
//     return -1
// }
