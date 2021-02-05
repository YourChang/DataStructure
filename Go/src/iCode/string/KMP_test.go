package string

import "testing"

func TestKMP(t *testing.T) {
	var kmpTestData = []struct {
		main    string
		pattern string
		want    int
	}{{main: "abc abcdab abcdabcdabde",
		pattern: "bcdabd",
		want:    16,
	}, {
		main:    "aabbbbaaabbababbabbbabaaabb",
		pattern: "abab",
		want:    11,
	}, {
		main:    "aabbbbaaabbababbabbbabaaabb",
		pattern: "ababacd",
		want:    -1,
	}}
	for _, v := range kmpTestData {
		if v.want != KMP(v.main, v.pattern) {
			t.Fatalf("error occured when search %s in %s", v.pattern, v.main)
		}
	}
}
