package main

import "testing"

func TestArray_Insert(t *testing.T) {
	capacity := 10
	array := NewArray(uint(capacity))
	for i := 0; i < capacity-2; i++{
		err := array.Insert(uint(i), i+1)
		if err != nil{
			t.Fatal(err.Error())
		}
	}

	array.Print()

	array.Insert(uint(6), 999)
	array.Print()

	array.InsertToTail(666)
	array.Print()
}

func TestArray_Find(t *testing.T) {
	capacity := 10
	array := NewArray(uint(capacity))
	for i := 0; i < capacity-2; i++{
		err := array.Insert(uint(i), i+1)
		if err != nil{
			t.Fatal(err.Error())
		}
	}
	array.Print()

	t.Log(array.Find(0))
	t.Log(array.Find(9))
	t.Log(array.Find(11))
}

func TestArray_Delete(t *testing.T) {
	capacity := 10
	array := NewArray(uint(capacity))
	for i := 0; i < capacity-2; i++{
		err := array.Insert(uint(i), i+1)
		if err != nil{
			t.Fatal(err.Error())
		}
	}
	array.Print()

	array.Delete(0)
	array.Print()

	array.Delete(9)
	array.Print()

	array.Delete(11)
	array.Print()
}