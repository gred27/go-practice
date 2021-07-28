package main

import "fmt"

func pointer1() {
	var a int = 500
	var b int = 1000
	var p1 *int
	var p2 *int
	var p3 *int

	p1 = &a
	p2 = &b
	p3 = &a

	fmt.Printf("address: %p\n", p1)
	fmt.Printf("p value: %d\n", *p1)

	fmt.Printf("p1 == p2 : %v\n", p1 == p2)
	fmt.Printf("p1 == p3 : %v\n", p1 == p3)

	*p1 = 300

	fmt.Printf("p1 p3 : %d %d\n", *p1, *p3)
}
