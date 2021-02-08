package TrieTree

import "testing"

func TestNewTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("str")
	t.Log(trie.Find("str"))

	trie.Insert("fuck")
	t.Log(trie.Find("fuck"))
	t.Log(trie.Find("jerk"))

}
