package Sort_3

import "math"

/**
@Description: 计数排序
@Date: 1/26/2021
@Author: lichang
*/

func CountSort(a []int, n int) {
	if n <= 1{
		return
	}

	var max int = math.MinInt32
	for i := range a {  // 获取最大值
		if a[i] > max {
			max = a[i]
		}
	}

	c := make([]int, max+1)
	for i := range a {   // 类似桶排序中的分配
		c[a[i]]++
	}
	for i := 1; i <= max; i++ {  // 计数
		c[i] += c[i-1]
	}

	res := make([]int, n)
	for i := n-1; i >= 0; i-- {  // 收集
		index := c[a[i]] -1
		res[index] = a[i]
		c[a[i]]--
	}
	copy(a, res)
}
