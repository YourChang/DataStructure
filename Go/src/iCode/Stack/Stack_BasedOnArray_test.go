package Stack

import (
	"testing"
)

func TestArrayStack_Push(t *testing.T) {
	var S *ArrayStack = NewArrayStack()
	for i := 0; i < 9; i++{
		S.Push(i)
	}
	S.Print()

	S.Push(999)
	S.Print()
}

func TestArrayStack_Pop(t *testing.T) {
	var S *ArrayStack = NewArrayStack()
	for i := 0; i < 9; i++{
		S.Push(i)
	}
	S.Print()

	S.Pop()
	S.Print()
}

func TestArrayStack_IsEmpty(t *testing.T) {
	var S *ArrayStack = NewArrayStack()
	if !S.IsEmpty(){
		t.Error("something was wrong about function IsEmpty() ")
	}

	for i := 0; i < 9; i++{
		S.Push(i)
	}
	if S.IsEmpty(){
	t.Error("something was wrong about function IsEmpty() ")
	}

}

func TestArrayStack_Flush(t *testing.T) {
	var S *ArrayStack = NewArrayStack()
	if !S.IsEmpty(){
		t.Error("something was wrong about function IsEmpty() ")
	}

	for i := 0; i < 9; i++{
		S.Push(i)
	}

	S.Flush()
	if !S.IsEmpty(){
		t.Error("something was wrong about function Flush() ")
	}
}
