package Sort_3

import (
	s "iCode/Sort_2"
	"testing"
)

func TestCountSort(t *testing.T) {
	a := []int{1, 7, 3, 5, 4, 6, 8, 5, 2}
	t.Log(a)
	CountSort(a, len(a))
	t.Log(a)
	if !s.IsSorted(a, 0, len(a)-1){
		t.Error("some error occured")
	}
}
