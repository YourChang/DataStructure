package Stack

import "testing"

func TestLinkListStack_Push(t *testing.T) {
	var S *LinkListStack = NewLinkListStack()
	for i := 0; i < 9; i++{
		S.Push(i)
	}
	S.Print()

	S.Push(999)
	S.Print()
}

func TestLinkListStack_Pop(t *testing.T) {
	var S *LinkListStack = NewLinkListStack()
	for i := 0; i < 9; i++{
		S.Push(i)
	}
	S.Print()

	S.Pop()
	S.Print()
}

func TestLinkListStack_IsEmpty(t *testing.T) {
	var S *LinkListStack = NewLinkListStack()
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

func TestLinkListStack_Flush(t *testing.T) {
	var S *LinkListStack = NewLinkListStack()
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
