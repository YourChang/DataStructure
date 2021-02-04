package string

import "testing"

func TestBFSearch(t *testing.T) {
	main := "abcdacdf"
	pattern := "acd"
	if BFSearch(main, pattern) != 4 {
		t.Fatal("something error occured")
	}
}
