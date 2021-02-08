package TrieTree

/**
@Description: Trie 树，也叫“字典树”。
		它是一种专门处理字符串匹配的数据结构，
		用来解决在一组字符串集合中快速查找某个字符串的问题
@Date: 2/6/2021
@Author: lichang
*/

type Trie struct {
	Data      rune
	Children  []*Trie
	IsEndChar bool
}

func NewTrie() *Trie {
	return &Trie{
		Data:      '/',
		Children:  make([]*Trie, 26),
		IsEndChar: false,
	}
}

func (trie *Trie) Insert(str string) {
	var t *Trie = trie
	for _, v := range str {
		index := v - 'a'
		if t.Children[index] == nil {
			t.Children[index] = NewTrie()
			t.Children[index].Data = v
		}
		t = t.Children[index]
	}
	t.IsEndChar = true
}

func (trie *Trie) Find(pattern string) bool {
	var t *Trie = trie
	for _, v := range pattern {
		index := v - 'a'
		if t.Children[index] == nil {
			return false
		}
		t = t.Children[index]
	}
	if t.IsEndChar == false {
		return false
	}
	return true
}
