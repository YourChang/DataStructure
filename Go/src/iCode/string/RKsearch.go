package string

/**
@Description: Rabin-Karp 算法，通过对字符串求哈希值减少对比次数
@Date: 2/4/2021
@Author: lichang
*/

var p []int = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

func hash(pattern string) int {
	var res int = 0

	for i, v := range pattern {
		res += (i + 1) * p[int(v)-int('a')]
	}
	return res
}

func RKSearch(main, pattern string) int {
	if len(main) == 0 || len(main) < len(pattern) || len(pattern) == 0 {
		return -1
	}
	hashP := hash(pattern)
	for i := 0; i < len(main)-len(pattern)+1; i++ {
		hashM := hash(main[i : i+len(pattern)])
		if hashM == hashP && pattern == main[i:i+len(pattern)] {
			return i
		}
	}
	return -1
}
