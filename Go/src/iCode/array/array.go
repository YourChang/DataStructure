package main

import (
	"errors"
	"fmt"
)

type Array struct{
	data []int
	length uint
}

// 创建数组
func NewArray(capacity uint)*Array{
	if capacity == 0 {
		return nil
	}
	return &Array{
		data: make([]int, capacity, capacity),
		length: 0,
	}
}

//返回数组长度
func (array *Array)  Len() uint {
	return array.length
}

// 判断索引是否越界
func (array *Array) isIndexOutOfRange(index uint) bool{
	if index >= uint(cap(array.data)){
		return true
	}
	return false
}

// 通过索引查找数组，索引范围【0， n-1】
func (array *Array) Find(index uint) (int, error){
	if array.isIndexOutOfRange(index){
		return 0, errors.New("index out of range")
	}
	return array.data[index], nil
}

// 插入数值到索引index上
func (array *Array) Insert(index uint, v int) error{
	if array.Len() == uint(cap(array.data)){
		return errors.New("full array")
	}
	if index != array.length && array.isIndexOutOfRange(index) {
		return errors.New("index out of range")
	}
	for i := array.length; i > index; i-- {
		array.data[i] = array.data[i-1]
	}
	array.data[index] = v
	array.length ++
	return nil
}

// 插入数值到数组最末
func (array *Array) InsertToTail(v int) error {
	return array.Insert(array.length, v)
}

// 删除索引上的值
func (array *Array) Delete(index uint) (int, error){
	if index < 0 || array.isIndexOutOfRange(index){
		return 0, errors.New("error index")
	}
	v := array.data[index]
	for i := index; i < array.length; i++ {
		array.data[i] = array.data[i+1]
	}
	array.length --
	return v, nil
}

// 打印数组
func (array *Array) Print(){
	var format string
	for i := uint(0); i < array.Len(); i++ {
		format += fmt.Sprintf("|%+v", array.data[i])
	}
	fmt.Println(format)
}