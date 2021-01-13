package main

import (
	"errors"
	"fmt"
)

/**
@Description: 单链表的基本操作
@Date: 1/13/2021
@Author: lichang
*/

type ListNode struct {
	next *ListNode
	value interface{}
}

type LinkList struct {
	head *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{
		next:  nil,
		value: v,
	}
}

func (L *ListNode) GetNext() *ListNode {
	return L.next
}

func (L *ListNode) GetValue() interface{} {
	return L.value
}

func NewLinkList() *LinkList {
	node := NewListNode(-1)
	return &LinkList{
		head:   node,
		length: 0,
	}
}

func (L *LinkList) InsertAfter(p *ListNode, v interface{}) error {
	if L == nil {
		return errors.New("empty list")
	}
	if p == nil {
		return errors.New("error node")
	}
	node := NewListNode(v)
	node.next = p.next
	p.next = node
	L.length ++
	return nil
}

func (L *LinkList) InsertBefore(p *ListNode, v interface{}) error {
	if p == nil {
		return errors.New("empty list")
	}
	var pre *ListNode = L.head
	for pre.next != p{
		pre = pre.next
	}
	node := NewListNode(v)
	pre.next = node
	node.next =  p
	L.length++
	return nil
}

func (L *LinkList) InsertToHead(v interface{}) error {
	if L == nil{
		return errors.New("empty list")
	}
	newNode := NewListNode(v)
	newNode.next = L.head.next
	L.head.next = newNode
	L.length ++
	return nil
}

func (L *LinkList) InsertToTail(v interface{}) error {
	if L == nil {
		return errors.New("empty list")
	}
	var p *ListNode = L.head
	for p.next != nil {
		p = p.next
	}
	p.next = NewListNode(v)
	return nil
}

func (L *LinkList) FindByIndex(index uint) (error, *ListNode) {
	if L == nil {
		return errors.New("empty list"), nil
	}
	if index > L.length {
		return errors.New("error index"), nil
	}
	var p *ListNode = L.head
	for i := 0; uint(i) < index; i++ {
		p = p.next
	}
	return nil, p
}

func (L *LinkList) DeleteNode(p *ListNode) error{
	if L == nil {
		return errors.New("empty list")
	}
	var pre, cur *ListNode
	pre = L.head
	cur = pre.next
	for cur != p{
		pre = cur
		cur = cur.next
	}
	pre.next = cur.next
	return nil
}

func (L *LinkList) Print() {
	var result string
	var p *ListNode = L.head.next
	for p != nil {
		result += fmt.Sprintf("%d|", p.value)
		p = p.next
	}
	fmt.Println(result)
}

