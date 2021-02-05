package string

/**
@Description: KMP算法
@Date: 2/5/2021
@Author: lichang
*/

// nexts数组中存储的值实际上是模式串中以每个位置为结束的
// 最长可匹配后缀子串对应的最长可匹配前缀子串的结束位置
func getNexts(pattern string) []int {
	m := len(pattern)
	nexts := make([]int, m)
	for index := range nexts {
		nexts[index] = -1
	}

	for i := 1; i < m-1; i++ {
		j := nexts[i-1] // 通过当前最长可匹配前缀子串寻找下一个最长可匹配前缀子串的位置

		for pattern[j+1] != pattern[i] && j >= 0 {
			j = nexts[j] // 最长可匹配后缀子串失配，转向次长可匹配后缀子串
		}

		if pattern[j+1] == pattern[i] { // 如果当前最长可匹配后缀子串后一位依然适配
			j += 1 // 最长可匹配后缀子串增1
		}

		nexts[i] = j
	}

	return nexts
}

func KMP(main string, pattern string) int {
	n := len(main)
	m := len(pattern)
	if n < m {
		return -1
	}

	nexts := getNexts(pattern)

	j := 0                   // j标识模式串中的位置
	for i := 0; i < n; i++ { // i标识主串中的位置
		for j > 0 && main[i] != pattern[j] { // 失配
			j = nexts[j-1] + 1 // 借助nexts数组更新模式串位置
		}

		if main[i] == pattern[j] {
			if j == m-1 { // 匹配完毕
				return i - m + 1
			}
			j += 1 // 继续向后匹配
		}
	}
	return -1
}
