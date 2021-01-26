package Sort_2

func MergeSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}

	mergeSort(arr, 0, arrLen-1)
}


func mergeSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}

	mid := (lo + hi) / 2
	mergeSort(a, lo, mid)
	mergeSort(a, mid+1, hi)
	merge(a, lo, mid, hi)
}

func merge(a []int, lo, mid, hi int) {
	tmpArr := make([]int, hi-lo+1)
	i := lo
	j := mid+1
	k := 0
	for ; i <= mid && j <= hi; k++{
		if a[i] <= a[j]{
			tmpArr[k] = a[i]
			i++
		}else{
			tmpArr[k] = a[j]
			j++
		}
	}

	for ;i <= mid ;i++{
		tmpArr[k] = a[i]
		k++
	}

	for ;j <= hi; j++ {
		tmpArr[k] = a[j]
		k++
	}

	copy(a[lo:hi+1], tmpArr)
}

