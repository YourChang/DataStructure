package Coin

import "math"

/**
@Description: 有几种面值的硬币（[]int），求用最少的硬币数目支付amount元的硬币数目
@Date: 2/15/2021
@Author: lichang
*/

// 状态转移方程为：dp[i] = 1 + min(dp[i]-coins...)

func coin(coins []int, amount int) int {
	max := amount + 1

	dp := make([]int, amount+1) // 初始化
	for i, _ := range dp {
		dp[i] = max
	}

	dp[0] = 0

	for i := 1; i <= amount; i++ { // 将dp数组扫描一遍
		for j := 0; j < len(coins); j++ { // 判断用哪一个硬币使得个数更少
			if coins[j] <= i {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coins[j]])+1))
			}
		}
	}

	return dp[amount]
}
