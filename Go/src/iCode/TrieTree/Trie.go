package TrieTree

/**
@Description: Trie 树，也叫“字典树”。
		它是一种专门处理字符串匹配的数据结构，
		用来解决在一组字符串集合中快速查找某个字符串的问题
@Date: 2/6/2021
@Author: lichang
*/

type Trie struct {
	data      rune
	children  []Trie
	isEndChar bool
}

func NewTrie() *Trie {
	return &Trie{
		data:      '/',
		children:  make([]Trie, 26),
		isEndChar: false,
	}
}

func (trie *Trie) Insert(str string) {
	var t *Trie = trie
	for _, v := range str {
		index := v - 'a'
		if t.children[index].data == 0 {
			t.children[index].data = v
			t.children[index].children = make([]Trie, 26)
		}
		t = &t.children[index]
	}
	t.isEndChar = true
}

func (trie *Trie) Find(pattern string) bool {
	var t *Trie = trie
	for _, v := range pattern {
		index := v - 'a'
		if t.children[index].data == 0 {
			return false
		}
		t = &t.children[index]
	}
	if t.isEndChar == false {
		return false
	}
	return true
}
