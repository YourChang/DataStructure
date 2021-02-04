package string

import "math"

/**
@Description: BM（Boyer-Moore）算法，实现高效字符串匹配
@Date: 2/4/2021
@Author: lichang
*/

func generateBC(pattern string) []int {
	bc := make([]int, 256) // 记录每个字符最后出现的位置

	for index := range bc {
		bc[index] = -1
	}

	for index, char := range pattern {
		bc[int(char)] = index
	}

	return bc
}

func generateGS(pattern string) ([]int, []bool) {
	m := len(pattern)

	// suffix的下标 k，表示后缀子串的长度，下标对应的数组值存储的是，
	// 在模式串中跟好后缀{u}相匹配的子串{u*}的起始下标值
	suffix := make([]int, m)

	// 模式串的后缀子串是否能匹配模式串的前缀子串
	prefix := make([]bool, m)

	// init _fix
	for i := 0; i < m; i++ {
		suffix[i] = -1
		prefix[i] = false
	}

	for i := 0; i < m-1; i++ {
		j := i // j是公共后缀字串的起始下标
		k := 0 // k是公共后缀子串的长度
		for j >= 0 && pattern[j] == pattern[m-1-k] {
			j--               // 公共后缀字串起始位置往前移一位
			k++               // 公共后缀字串变长
			suffix[k] = j + 1 // 更新公共后缀字串起始位置
		}

		if j == -1 {
			// 模式串的后缀子串可以能匹配模式串的前缀子串
			prefix[k] = true
		}
	}

	return suffix, prefix

}

func moveByGS(patternLength int, badCharStartIndex int, suffix []int, prefix []bool) int {
	// badCharStartIndex表示坏字符对应的模式串中的字符下标

	k := patternLength - badCharStartIndex - 1 // 好后缀长度

	if suffix[k] != -1 { // 模式中存在坏字符
		return badCharStartIndex + 1 - suffix[k]
	}

	for t := patternLength - 1; t > badCharStartIndex; t-- {
		if prefix[t] { // 存在可匹配的前缀子串
			return t
		}
	}

	return patternLength
}

func findBadChar(subStr string, pattern string, bc []int) (int, int) {
	j := -1
	k := -1
	badChar := rune(0)

	for index := len(subStr) - 1; index >= 0; index-- {
		if subStr[index] != pattern[index] {
			j = index
			badChar = rune(subStr[index])
			break
		}
	}

	if j > 0 { // 坏字符存在
		k = bc[int(badChar)]
	}

	return k, j
}

func BMSearch(main string, pattern string) int {
	if len(main) == 0 || len(pattern) == 0 || len(pattern) > len(main) {
		return -1
	}

	bc := generateBC(pattern)
	suffix, prefix := generateGS(pattern)

	n := len(main)
	m := len(pattern)

	step := 1
	for i := 0; i <= n-m; i = i + step {
		subStr := main[i : i+m]
		k, j := findBadChar(subStr, pattern, bc)
		// j是坏字符的位置（主串中）
		// k是子串中坏字符位置

		stepForBC := j - k // 根据坏字符规则应移动的步幅
		if j == -1 {
			return i
		}

		stepForGS := -1 // 根据好后缀规则应移动的步幅
		if j < m-1 {
			stepForGS = moveByGS(m, j, suffix, prefix)
		}

		step = int(math.Max(float64(stepForBC), float64(stepForGS)))
	}

	return -1
}
