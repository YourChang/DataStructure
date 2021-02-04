package string

import (
	"testing"
)

func TestBMSearch(t *testing.T) {
	main := "abcacabcbcabcabc"
	pattern := "cabcab"

	if BMSearch(main, pattern) != 9 {
		t.Fatal("errors occured")
	}

}
