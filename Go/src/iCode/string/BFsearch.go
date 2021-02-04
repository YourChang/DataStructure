package string

/**
@Description: BFsearch, 暴力匹配字符串算法
@Date: 2/4/2021
@Author: lichang
*/

func BFSearch(main string, pattern string) int {
	if len(main) == 0 || len(pattern) == 0 || len(main) < len(pattern) {
		return -1
	}

	for i := 0; i <= len(main)-len(pattern); i++ {
		subStr := main[i : i+len(pattern)]
		if subStr == pattern {
			return i
		}
	}

	return -1
}
