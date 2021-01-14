package Stack

import "fmt"

type Brower struct {
	forwardStack Stack
	backStack    Stack
}

func NewBrower() *Brower {
	return &Brower{
		forwardStack: NewArrayStack(),
		backStack:    NewLinkListStack(),
	}
}

func (B *Brower) CanBack() bool {
	if B.backStack.IsEmpty() {
		return false
	}
	return true
}

func (B *Brower) Open(addr string) {
	fmt.Printf("Open new addr %+v\n", addr)
	B.forwardStack.Flush()
}

func (B *Brower) PushBack(addr string) {
	B.backStack.Push(addr)
}

func (B *Brower) Forward() {
	if B.forwardStack.IsEmpty() {
		return
	}
	top := B.forwardStack.Pop()
	B.backStack.Push(top)
	fmt.Printf("forward to %+v\n", top)
}

func (B *Brower) Back() {
	if B.backStack.IsEmpty() {
		return
	}
	top := B.backStack.Pop()
	B.forwardStack.Push(top)
	fmt.Printf("back to %+v\n", top)
}
