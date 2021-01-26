package Sort_1

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{1, 7, 3, 5, 4, 6, 8, 5, 2}
	n := 9
	Print(a, n)
	BubbleSort(a, n)
	Print(a, n)

}

func TestInsertSort(t *testing.T) {
	a := []int{1, 7, 3, 5, 4, 6, 8, 5, 2}
	n := 9
	Print(a, n)
	InsertSort(a, n)
	Print(a, n)
}

func TestSelectSort(t *testing.T) {
	a := []int{1, 7, 3, 5, 4, 6, 8, 5, 2}
	n := 9
	Print(a, n)
	SelectSort(a, n)
	Print(a, n)
}
