package Queue

import (
	"fmt"
	"testing"
)

func TestArrayQueue_EnQueue(t *testing.T) {
	var Q *ArrayQueue = NewArrayQueue(5)
	if !Q.IsEmpty(){
		t.Error("something was wrong about Q.IsEmpty()")
	}
	for i := 0; i < 5; i++ {
		Q.EnQueue(i)
	}
	fmt.Println(Q)
}

func TestArrayQueue_DeQueue(t *testing.T) {
	var Q *ArrayQueue = NewArrayQueue(5)
	if !Q.IsEmpty(){
		t.Error("something was wrong about Q.IsEmpty()")
	}
	for i := 0; i < 5; i++ {
		Q.EnQueue(i)
	}
	fmt.Println(Q)

	fmt.Printf("%+v\n", Q.DeQueue())
	fmt.Println(Q)
}
