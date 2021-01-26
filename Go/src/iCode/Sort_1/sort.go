package Sort_1

import "fmt"

/**
@Description: 三种常见的排序算法：冒泡、插入、选择
@Date: 1/25/2021
@Author: lichang
*/

type Sorter interface{
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Unsorted struct{ a []int }

func (s Unsorted)Len()int{
	return len(s.a)
}

func (s Unsorted)Less(i, j int) bool {
	if s.a[i] < s.a[j] {
		return true
	}
	return false
}

func (s Unsorted)Swap(i, j int) {
	s.a[i], s.a[j] = s.a[j], s.a[i]
}



// 冒泡排序
func BubbleSort(a []int , n int) {
	if n <= 1{
		return
	}
	for i := 0; i < n; i++ {
		flag := false
		for j := n-1; j > i; j-- {
			if a[j-1] > a[j]{
				a[j-1], a[j] = a[j], a[j-1]
				flag = true
			}
		}
		if flag == false{
			break
		}
	}
	return
}

// 插入排序
func InsertSort(a []int, n int){
	if n <= 1{
		return
	}
	var i, j int;
	for i = 1; i < n; i++ {
		if a[i] < a[i-1]{
			guard := a[i] // 哨兵
			for j = i-1; a[j] > guard; j-- {
				a[j], a[j+1] = a[j+1], a[j]
			}
			a[j+1] = guard
		}
	}
}

// 选择排序
func SelectSort(a []int, n int){
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		minflag := i
		for j := i+1; j < n; j++ {
			if a[minflag] > a[j]{
				minflag = j
			}
		}
		a[i], a[minflag] = a[minflag], a[i]
	}
}

func Print(a []int, n int) {
	res := ""
	for i := 0; i < n; i++ {
		res+=fmt.Sprintf("%d|", a[i])
	}
	fmt.Println(res)
}


