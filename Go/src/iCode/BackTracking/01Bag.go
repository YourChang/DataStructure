package BackTracking

/**
@Description: 使用回溯法解决01背包问题
@Date: 2/9/2021
@Author: lichang
*/

var maxW int = -1 //存储背包中物品总重量的最大值

// cw表示当前已经装进去的物品的重量和；i表示考察到哪个物品了；
//w背包可承受重量；items表示每个物品的重量；n表示物品个数
//假设背包可承受重量100，物品个数10，物品重量存储在数组a中，那可以这样调用函数：
//f(0, 0, a, 10, 100)

func Bag01(i, cw int, items []int, n, w int) {
	if cw == w || i == n {
		if cw < maxW {
			maxW = cw
		}
		return
	}
	Bag01(i+1, cw, items, n, w) // 当前物品不装
	if cw+items[i] <= w {
		Bag01(i+1, cw+items[i], items, n, w) // 当前物品装入
	}
}
