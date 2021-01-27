package BinarySearch

/**
@Description: 二分查找
@Date: 1/27/2021
@Author: lichang
*/

// 基本二分查找
func BinarySearch(a []int , v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	lo := 0
	hi := n-1

	for lo <= hi {
		mid := (lo + hi) / 2
		if a[mid] == v {
			return mid
		} else if a[mid] < v {
			lo = mid +1
		} else  {
			hi = mid -1
		}
	}

	return -1
}

// 递归版迭代查找
func BinarySearchRecursive(a []int, v int) int {
	n := len(a)
	if n== 0{
		return -1
	}
	return bs(a, v, 0, n-1)
}

func bs(a []int, v int, low, high int ) int {
	if low > high {
		return -1
	}

	mid := (low + high ) / 2
	if a[mid] == v {
		return mid
	} else if a[mid] > v {
		return bs(a, v, low, mid-1)
	} else {
		return bs(a, v, mid+1, high)
	}
}


// 找到第一个符合条件的元素
func BinarySearchFirst(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	lo := 0
	hi := n - 1
	var mid int
	for lo <= hi{
		mid = (lo + hi) / 2
		if a[mid] > v {
			lo = mid + 1
		} else if a[mid] < v {
			hi = mid - 1
		} else {
			if mid == 0 || a[mid-1] != v {
				return mid
			} else {
				hi = mid -1
			}
		}
	}
	return -1
}

// 找到最后一个满足条件的元素
func BinarySearchLast(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	lo := 0
	hi := n - 1
	var mid int
	for lo <= hi {
		mid = (lo + hi) / 2
		if a[mid] > v {
			hi = mid - 1
		} else if a[mid] < v {
			lo = mid + 1
		} else {
			if mid == n-1 || a[mid+1] != v{
				return mid
			} else {
				lo = mid + 1
			}
		}
	}
	return -1
}

// 查找第一个大于等于给定值的元素
func BinarySearchFirstGT(a []int, v int) int {
	n := len(a)
	if n == 0 {return -1}

	lo := 0
	hi := n-1
	var mid int
	for lo <= hi {
		mid = (lo + hi) / 2
		if a[mid] > v {
			hi = mid - 1
		} else if a[mid] < v {
			lo = mid + 1
		}else{
			if mid != n-1 && a[mid+1] > v {
				return mid + 1
			}else{
				lo = hi + 1
			}
		}
	}
	return -1
}


