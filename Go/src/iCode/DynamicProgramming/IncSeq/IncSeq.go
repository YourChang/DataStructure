package IncSeq

import "math"

/**
@Description: 有一个数字序列包含 n 个不同的数字，
			如何求出这个序列中的最长递增子序列长度？
			比如 2, 9, 3, 6, 5, 1, 7 这样一组数字序列，
			它的最长递增子序列就是 2, 3, 5, 7，所以最长递增子序列的长度是 4
@Date: 2/15/2021
@Author: lichang
*/

// 状态转移方程 dp[i]=max(dp[j])+1,其中0≤j<i且num[j]<num[i]

func IncSeq(nums []int) int {
	dp := make([]int, len(nums))

	dp[0] = 1
	maxans := 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
		maxans = int(math.Max(float64(maxans), float64(dp[i])))
	}
	return maxans
}
