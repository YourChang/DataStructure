package Coin

import "testing"

func Test(t *testing.T) {
	coins := []int{1, 3, 5}
	if coin(coins, 9) != 3 {
		t.Fatal()
	}
}
