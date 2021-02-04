package string

import "testing"

func TestRKSearch(t *testing.T) {
	main := "abcdacdf"
	pattern := "acd"
	if RKSearch(main, pattern) != 4 {
		t.Fatal("something error occured")
	}
}
