package leetcode

//https://leetcode-cn.com/problems/max-area-of-island/
func maxAreaOfIsland(grid [][]int) int {
	visit := make(map[[2]int]bool)
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			cnt := area(grid, visit, i, j)
			if max < cnt {
				max = cnt
			}
		}
	}
	return max
}

func area(grid [][]int, visit map[[2]int]bool, i int, j int) int {
	if grid[i][j] == 0 || visit[[2]int{i, j}] {
		return 0
	}
	cnt := 1
	visit[[2]int{i, j}] = true
	if i > 0 {
		cnt += area(grid, visit, i-1, j)
	}
	if j > 0 {
		cnt += area(grid, visit, i, j-1)
	}
	if i < len(grid)-1 {
		cnt += area(grid, visit, i+1, j)
	}
	if j < len(grid[i])-1 {
		cnt += area(grid, visit, i, j+1)
	}
	return cnt
}
