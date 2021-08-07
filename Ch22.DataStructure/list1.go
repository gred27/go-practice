// Linked List
package main

import (
	"container/list"
	"fmt"
)

func list1() {

	v := list.New()
	e4 := v.PushBack(4)
	e1 := v.PushFront(1)
	v.InsertAfter(3, e4)
	v.InsertBefore(2, e1)

	fmt.Println(v)

	// 순회
	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	// 역방향 순회
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
}
