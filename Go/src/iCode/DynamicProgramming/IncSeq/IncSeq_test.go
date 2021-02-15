package IncSeq

import "testing"

func TestIncSeq(t *testing.T) {
	var test = []struct {
		nums []int
		want int
	}{
		{nums: []int{4, 10, 4, 3, 8, 9},
			want: 3},
		{nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4},
	}
	for _, v := range test {
		if IncSeq(v.nums) != v.want {
			t.Fatal()
		}
	}

}
