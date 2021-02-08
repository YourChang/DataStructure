package Aho_Corasick

import "testing"

func TestAc_Match(t *testing.T) {
	var ac *Ac = NewAc()
	ac.Insert("abcd")
	ac.Insert("bcd")
	ac.Insert("c")

	ac.BuildFailPointer()
	ac.Match("abcaabcda")
}
