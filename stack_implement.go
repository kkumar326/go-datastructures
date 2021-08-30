package main

import (
	"fmt"

	stack "github.com/kkumar326/go-dsa/pkgs/stack"
)

func stackRun() {
	var element stack.ItemType = "First element"
	stack := stack.NewStack(&element)

	stack.Push(1)
	stack.Push("Second")
	stack.Push(true)
	//stack.Pop()
	//stack.Clear()

	fmt.Println(stack.Peek())
}
