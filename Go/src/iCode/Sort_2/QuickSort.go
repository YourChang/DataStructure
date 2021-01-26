package Sort_2

import "fmt"

/**
@Description: 快排算法
@Date: 1/26/2021
@Author: lichang
*/

func QuickSort(a []int , lo, hi int){
	if lo < hi {
		pivotPos := Partition(a, lo, hi)
		QuickSort(a, lo, pivotPos-1)
		QuickSort(a, pivotPos+1, hi)
	}
}

func Partition(a []int, lo, hi int) int{
	pivot := a[lo]
	for lo < hi {
		for lo < hi && a[hi] >= pivot {  //从后往前找，直至a[hi]小于当前枢轴值
			hi--
		}
		a[lo] = a[hi]  // 将找到的较小值放在lo
		for lo < hi && a[lo] <= pivot {  // 从前往后找，直至a[lo]大于当前枢轴值
			lo++
		}
		a[hi] = a[lo]  // 将找到的较大值放在hi
	}
	a[lo] = pivot
	return lo
}

func IsSorted(a []int, lo, hi int) bool {
	for lo < hi{
		if a[lo] > a[lo+1]{
			return false
		}
		lo++
	}
	return true
}

func Print(a []int, n int) {
	res := ""
	for i := 0; i < n; i++ {
		res+=fmt.Sprintf("%d|", a[i])
	}
	fmt.Println(res)
}

