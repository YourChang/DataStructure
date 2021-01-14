package Stack

import "fmt"

/**
@Description: 基于数组实现的栈
@Date: 1/14/2021
@Author: lichang
*/


type ArrayStack struct {
	data []interface{}
	top int
}

func NewArrayStack() *ArrayStack{
	return &ArrayStack{
		data: make([]interface{}, 0, 5),
		top:  -1,
	}
}

func (S *ArrayStack) IsEmpty() bool {
	if S.top == -1 {
		return true
	}
	return false
}

func (S *ArrayStack) Push(v interface{}){
	if S.top < 0 {
		S.top = 0
	} else {
		S.top += 1
	}

	if S.top > len(S.data)-1 {
		S.data = append(S.data, v)
	} else {
		S.data[S.top] = v
	}
}

func (S *ArrayStack) Pop() interface{} {
	if S == nil {
		panic("the stack does note exist")
	}
	if S.IsEmpty(){
		fmt.Println("empty stack")
	}
	v := S.data[S.top]
	S.data = S.data[:len(S.data)-1]
	return v
}

func (S *ArrayStack)Top()interface{}{
	if S == nil {
		panic("the stack does note exist")
	}
	if S.IsEmpty(){
		fmt.Println("empty stack")
	}
	return S.data[S.top]
}

func (S *ArrayStack) Flush(){
	d := make([]interface{}, 0, 5)
	S.data = d
	S.top = -1
}

func (S *ArrayStack) Print() {
	if S == nil {
		fmt.Println("the stack does not exist")
	}
	var result string
	for i := 0; i < len(S.data); i++ {
		result += fmt.Sprintf("%d|", S.data[i])
	}
	fmt.Println(result)
}