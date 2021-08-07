// Stack 구현
package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func (st *Stack) Push(val interface{}) {
	st.v.PushBack(val)
}

func (st *Stack) Pop() interface{} {
	back := st.v.Back()
	if back != nil {
		return st.v.Remove(back)
	}

	return nil
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func stack1() {
	// stack := NewStack()
	var stack Stack = Stack{list.New()}

	for i := 1; i < 5; i++ {
		stack.Push(i)
	}

	v := stack.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = stack.Pop()
	}
}
