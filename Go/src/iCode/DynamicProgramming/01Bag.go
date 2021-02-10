package DynamicProgramming

/**
@Description: 01背包问题动态规划解法
@Date: 2/10/2021
@Author: lichang
*/

func knapsack(weight []int, n, w int) int {
	states := make([][]bool, n) // 初始化状态数组
	for i, _ := range states {
		states[i] = make([]bool, w+1)
	}

	states[0][0] = true
	if weight[0] <= w {
		states[0][weight[0]] = true
	}

	for i := 1; i < n; i++ { // 动态规划状态转移
		for j := 0; j <= w; j++ {
			if states[i-1][j] == true { // 不装入
				states[i][j] = true
			}
		}
		for j := 0; j <= w; j++ { // 装入
			if states[i-1][j] == true {
				states[i][j+weight[i]] = true
			}
		}
	}

	for i := w; i >= 0; i-- {
		if states[n-1][i] == true {
			return i
		}
	}

	return 0
}

func knapsack2(items []int, n, w int) int {
	states := make([]bool, w+1)

	states[0] = true
	if items[0] <= w {
		states[items[0]] = true
	}

	for i := 1; i < n; i++ {
		for j := w - items[i]; j >= 0; j-- {
			if states[j] == true {
				states[j+items[i]] = true
			}
		}
	}

	for i := w; i >= 0; i-- {
		if states[i] == true {
			return i
		}
	}

	return 0
}
