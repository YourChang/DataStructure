package minDist

import "testing"

func TestMinDistDp(t *testing.T) {
	var matrix [][]int = [][]int{{1, 3, 5, 9}, {2, 1, 3, 4}, {5, 2, 6, 7}, {6, 8, 4, 3}}
	if minDistDp(matrix, 4) != 19 {
		t.Fatal()
	}
}
