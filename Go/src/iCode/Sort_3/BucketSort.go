package Sort_3

/**
@Description: 桶排序
@Date: 1/26/2021
@Author: lichang
*/

func BucketSort(a []int){
	max := getMax(a)
	bucket := make([]int, max+1)
	for _, v := range a{  //分配到桶中
		bucket[v] ++
	}
	tmpArr := make([]int, 0)
	for i, v := range bucket {
		if v <= 0{
			continue
		} else {
			for j := v; j > 0; j-- {
				tmpArr = append(tmpArr, i)
			}
		}

	}
	copy(a, tmpArr)
}

func getMax(a []int) int {
	max := a[0]
	for _, v := range a{
		if max < v {
			max = v
		}
	}
	return max
}
