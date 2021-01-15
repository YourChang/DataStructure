package Queue

import "fmt"

type node struct {
	v interface{}
	next *node
}

type LinkListQueue struct {
	head *node
	tail *node
	length uint
	capacity uint
}

func NewLinkListQueue() *LinkListQueue{
	return &LinkListQueue{
		head:   nil,
		tail:   nil,
		length: 0,
		capacity: 0,
	}
}

func (Q *LinkListQueue) IsEmpty() bool {
	return Q.length == 0
}

func (Q *LinkListQueue) IsFull() bool {
	return Q.length == Q.capacity
}

func (Q *LinkListQueue) EnQueue(v interface{}) bool {
	if Q.IsFull() {
		return false
	}
	newNode := &node{
		v:    v,
		next: nil,
	}
	if Q.tail == nil {
		Q.tail = newNode
		Q.head = newNode
	} else {
		Q.tail.next = newNode
		Q.tail = newNode
	}
	Q.length++
	return true
}

func (Q *LinkListQueue) DeQueue() interface{}{
	if Q.IsEmpty(){
		return nil
	}
	v := Q.head.v
	Q.head = Q.head.next
	Q.length--
	return v
}

func (Q *LinkListQueue) String() string {
	if Q.head == nil {
		return "empty queue"
	}
	result := "head<-"
	for cur := Q.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.v)
	}
	result += "<-tail"
	return result
}