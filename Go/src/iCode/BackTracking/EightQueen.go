package BackTracking

import "fmt"

/**
@Description: 八皇后问题，使用回溯算法求解
@Date: 2/9/2021
@Author: lichang
*/

var res []int = make([]int, 8) // 下标表示行，值表示queen存储在哪一列

func cal8queen(row int) {
	if row == 8 {
		printQueen(res)
		return
	}
	for column := 0; column < 8; column++ {
		if isOk(row, column) {
			res[row] = column
			cal8queen(row + 1)
		}
	}
}

func printQueen([]int) {
	var p string
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if res[row] == column {
				p += fmt.Sprintf("Q ")
			} else {
				p += fmt.Sprint("* ")
			}
		}
		p += fmt.Sprintf("\n")
	}
	fmt.Println(p)
}

func isOk(row, column int) bool {
	leftup := column - 1
	rightup := column + 1
	for i := row - 1; i >= 0; i-- {
		if res[i] == column {
			return false
		}
		if leftup >= 0 {
			if res[i] == leftup {
				return false
			}
		}
		if rightup < 8 {
			if res[i] == rightup {
				return false
			}
		}
		leftup--
		rightup++
	}
	return true
}
