package Stack

import "fmt"

type node struct {
	v interface{}
	next *node
}

type LinkListStack struct {
	top *node
}

func NewLinkListStack() *LinkListStack{
	// 链栈有头节点
	return &LinkListStack{
		top: &node{
			v:    "head",
			next: nil,
		},
	}
}

func (S *LinkListStack) IsEmpty() bool {
	if S.top.next == nil {
		return true
	}
	return false
}

func (S *LinkListStack) Push(v interface{}){
	var newnode node = node{
		v:    v,
		next: nil,
	}
	newnode.next = S.top.next
	S.top.next = &newnode
}

func (S *LinkListStack) Pop() interface{} {
	v := S.top.next.v
	S.top.next = S.top.next.next
	return v
}

func (S *LinkListStack) Top() interface{}{
	return S.top.next.v
}

func (S *LinkListStack) Flush() {
	S.top.next = nil
}

func (S *LinkListStack) Print() {
	var p *node = S.top.next
	var result string
	for p.next != nil {
		result += fmt.Sprintf("%d|", p.v)
		p = p.next
	}
	fmt.Println(result)
}