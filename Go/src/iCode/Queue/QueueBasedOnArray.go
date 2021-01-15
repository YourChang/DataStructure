package Queue

import "fmt"

type ArrayQueue struct {
	q []interface{}
	capacity uint
	head uint
	tail uint
}

func NewArrayQueue(n uint) *ArrayQueue {
	return &ArrayQueue{
		q:        make([]interface{}, n),
		capacity: n,
		head:     0,  //指向头元素
		tail:     0,  // 指向尾元素得下一个
	}
}

func (Q *ArrayQueue) IsEmpty() bool{
	return Q.head == Q.tail
}

func (Q *ArrayQueue) IsFull() bool {
	return Q.tail == Q.capacity
}

func (Q *ArrayQueue) EnQueue(v interface{}) bool {
	if Q.IsFull() {
		return false
	}
	Q.q[Q.tail]= v
	Q.tail++
	return true
}

func (Q *ArrayQueue) DeQueue() interface{}{
	if Q.IsEmpty(){
		return nil
	}
	v := Q.q[Q.head]
	Q.head++
	return v
}

func (Q *ArrayQueue) String()string {
	if Q.IsEmpty(){
		return "empty queue!"
	}
	result := "head"
	for i := Q.head; i < Q.tail; i++ {
		result += fmt.Sprintf("<-%+v", Q.q[i])
	}
	result += "<-tail"
	return result
}

