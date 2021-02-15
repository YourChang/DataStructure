package minDist

import "math"

/**
@Description: 使用动态规划算法求矩阵从左上角到右下角的最短距离
@Date: 2/15/2021
@Author: lichang
*/

func minDistDp(matrix [][]int, n int) int {
	states := make([][]int, n)
	for i, _ := range states {
		states[i] = make([]int, n)
	}

	sum := 0
	for j := 0; j < n; j++ {
		sum += matrix[0][j]
		states[0][j] = sum
	}

	sum = 0
	for i := 0; i < n; i++ {
		sum += matrix[i][0]
		states[i][0] = sum
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			states[i][j] = matrix[i][j] +
				int(math.Min(float64(states[i-1][j]), float64(states[i][j-1])))
		}
	}
	return states[n-1][n-1]
}
