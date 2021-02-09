package Aho_Corasick

/**
@Description: AC自动机，实现敏感词过滤
@Date: 2/8/2021
@Author: lichang
*/

import (
	"fmt"
	"iCode/Queue"
)

type Ac struct {
	Data      rune
	Children  []*Ac
	IsEndChar bool

	fail   *Ac
	length int
}

//type Ac struct {
//	TrieTree.Trie
//	fail *Ac
//	length int
//}

// 类似KMP，通过上一层的fail得出当前层的fail
func (ac *Ac) BuildFailPointer() {
	var queue *Queue.ArrayQueue = Queue.NewArrayQueue(10)
	var root *Ac = ac
	root.fail = nil
	queue.EnQueue(root)

	for !queue.IsEmpty() {
		var ip = queue.DeQueue()
		p := ip.(*Ac) // p为当前节点父节点
		for i := 0; i < 26; i++ {
			var pc *Ac = p.Children[i] // pc为当前节点
			if pc == nil {
				continue
			} // pc为空
			if p == root { // 父节点为根节点
				pc.fail = root // 失败指针指向父节点（即根）
			} else {
				q := p.fail    // q为p的失败指针指向节点
				for q != nil { // p不为空
					qc := q.Children[pc.Data-'a'] // qc是q的子节点
					if qc != nil {                // q的子节点不为空
						pc.fail = qc
						break
					}
					q = q.fail // 如果q的子节点为空，则使q为q的失败指针指向节点
				}
				if q == nil {
					pc.fail = root
				}
			}
			queue.EnQueue(pc)
		}
	}
}

func (ac *Ac) Match(main string) {
	n := len(main)
	var p *Ac = ac

	for i := 0; i < n; i++ {
		idx := main[i] - 'a'
		for p.Children[idx] == nil && p != ac { // 要寻找的字符不存在，使用fail指针
			p = p.fail
		}
		p = p.Children[idx]
		if p == nil {
			p = ac
		}
		var tmp *Ac = p
		for tmp != ac {
			if tmp.IsEndChar == true {
				fmt.Printf("匹配起始下标为%d, 长度为%d\n", i-tmp.length+1, tmp.length)
			}
			tmp = tmp.fail // 查看已匹配字符串的后缀是否是过滤词
		}
	}
}

func NewAc() *Ac {
	return &Ac{
		Data:      '/',
		Children:  make([]*Ac, 26),
		IsEndChar: false,
		fail:      nil,
		length:    -1,
	}
}

func (trie *Ac) Insert(str string) {
	var t *Ac = trie
	for _, v := range str {
		index := v - 'a'
		if t.Children[index] == nil {
			t.Children[index] = NewAc()
			t.Children[index].Data = v
			//t.Children[index].Children = make([]*Ac, 26)
		}
		t = t.Children[index]
	}
	t.IsEndChar = true
	t.length = len(str)
}

func (trie *Ac) Find(pattern string) bool {
	var t *Ac = trie
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
