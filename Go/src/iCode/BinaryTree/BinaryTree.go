package BinaryTree

import (
	"fmt"
	Stack "iCode/Stack"
)

/**
@Description: 二叉树
@Date: 1/29/2021
@Author: lichang
*/

type TreeNode struct {
	value int
	right *TreeNode
	left  *TreeNode
}

// 通过前序和中序序列生成二叉树
func PreIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("the length of pre and in is unequal")
	}

	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		value: pre[0],
		right: nil,
		left:  nil,
	}
	if len(pre) == 1 {
		return res
	}
	index := getIndex(in, res.value)
	if index == -1 {
		panic(fmt.Sprintf("can't find %d in inOrder", res.value))
	}

	res.right = PreIn2Tree(pre[index+1:], in[index+1:])
	res.left = PreIn2Tree(pre[1:index+1], in[:index])

	return res

}

func getIndex(arr []int, f int) int {
	for i, v := range arr {
		if v == f {
			return i
		}
	}
	return -1
}

func PreOrder(root *TreeNode, visit func(interface{})) {
	if root == nil {
		return
	}

	visit(root.value)
	PreOrder(root.left, visit)
	PreOrder(root.right, visit)
}

// 先序遍历非递归
func PreOrderNotR(root *TreeNode, visit func(interface{})) {
	if root == nil {
		return
	}
	var s *Stack.LinkListStack = Stack.NewLinkListStack()
	p := root
	for !s.IsEmpty() || p != nil {
		if p != nil {
			visit(p.value)
			s.Push(p)
			p = p.left
		} else {
			p = (s.Pop()).(*TreeNode)
			p = p.right
		}
	}
}

func InOrder(root *TreeNode, visit func(interface{})) {
	if root == nil {
		return
	}

	InOrder(root.left, visit)
	visit(root.value)
	InOrder(root.right, visit)
}

// 中序遍历非递归
func InOrderNotR(root *TreeNode, visit func(interface{})) {
	if root == nil {
		return
	}
	var s *Stack.ArrayStack = Stack.NewArrayStack()
	p := root
	for !s.IsEmpty() {
		if p != nil {
			s.Push(p)
			p = p.left
		} else {
			p = (s.Pop()).(*TreeNode)
			visit(p.value)
			p = p.right
		}
	}
}

func PostOrder(root *TreeNode, visit func(interface{})) {
	if root == nil {
		return
	}

	PostOrder(root.left, visit)
	PostOrder(root.right, visit)
	visit(root.value)
}
