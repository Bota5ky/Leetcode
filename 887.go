package leetcode

//https://leetcode-cn.com/problems/super-egg-drop/
func superEggDrop(K int, N int) int {
	dp := make([]int, K+1)
	ans := 0 // 操作的次数
	for dp[K] < N {
		for i := K; i > 0; i-- { // 从后往前计算
			dp[i] = dp[i] + dp[i-1] + 1
		}
		ans++
	}
	return ans
}

/**
鸡蛋掉落，鹰蛋（Leetcode 887）：（经典dp）
有 K 个鸡蛋，有 N 层楼，用最少的操作次数 F 检查出鸡蛋的质量。
思路：
本题应该逆向思维，若你有 K 个鸡蛋，你最多操作 F 次，求 N 最大值。
dp[k][f] = dp[k][f-1] + dp[k-1][f-1] + 1;
解释：
0.dp[k][f]：如果你还剩 k 个蛋，且只能操作 f 次了，所能确定的楼层。
1.dp[k][f-1]：蛋没碎，因此该部分决定了所操作楼层的上面所能容纳的楼层最大值
2.dp[k-1][f-1]：蛋碎了，因此该部分决定了所操作楼层的下面所能容纳的楼层最大值
又因为第 f 次操作结果只和第 f-1 次操作结果相关，因此可以只用一维数组。
时复：O(K*根号(N))
**/
