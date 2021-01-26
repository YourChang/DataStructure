package Sort_2

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	a := []int{1, 7, 3, 5, 4, 6, 8, 5, 2}
	QuickSort(a, 0, len(a)-1)
	if !IsSorted(a, 0, len(a)-1){
		t.Error("some error occured")
	}

}
