package main

import (
	"fmt"
	"testing"
)

func TestLinkList_InsertToHead(t *testing.T) {
	var L *LinkList = NewLinkList()
	var a [6]int = [6]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 6; i++ {
		err := L.InsertToHead(a[i])
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	L.Print()
}

func TestLinkList_InsertToTail(t *testing.T) {
	var L *LinkList = NewLinkList()
	var a [6]int = [6]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 6; i++ {
		err := L.InsertToTail(a[i])
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	L.Print()
}

func TestLinkList_InsertAfter(t *testing.T) {
	var L *LinkList = NewLinkList()
	var a [6]int = [6]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 6; i++ {
		err := L.InsertToHead(a[i])
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	L.Print()

	p := L.head.next.next
	err := L.InsertAfter(p, 555)
	if err != nil{
		fmt.Println(err)
	}
	L.Print()
}

func TestLinkList_InsertBefore(t *testing.T) {
	var L *LinkList = NewLinkList()
	var a [6]int = [6]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 6; i++ {
		err := L.InsertToHead(a[i])
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	L.Print()

	p := L.head.next.next
	err := L.InsertBefore(p, 555)
	if err != nil{
		fmt.Println(err)
	}
	L.Print()
}

func TestLinkList_FindByIndex(t *testing.T) {
	var L *LinkList = NewLinkList()
	var a [6]int = [6]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 6; i++ {
		err := L.InsertToHead(a[i])
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	L.Print()

	err, p := L.FindByIndex(3)
	if err != nil{
		fmt.Println(err)
	}
	if p.value != 4 {
		t.Error("false")
	}
}

func TestLinkList_DeleteNode(t *testing.T) {
	var L *LinkList = NewLinkList()
	var a [6]int = [6]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 6; i++ {
		err := L.InsertToHead(a[i])
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	L.Print()

	err, p := L.FindByIndex(3)
	if err != nil{
		fmt.Println(err)
	}

	err = L.DeleteNode(p)
	if err != nil{
		fmt.Println(err)
	}
	L.Print()
}
