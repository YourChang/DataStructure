package Sort_3

import (
	"testing"
	s "iCode/Sort_2"
)

func TestBucketSort(t *testing.T) {
	a := []int{1, 7, 3, 5, 4, 6, 8, 5, 2}
	t.Log(a)
	BucketSort(a)
	t.Log(a)
	if !s.IsSorted(a, 0, len(a)-1){
		t.Error("some error occured")
	}
}
