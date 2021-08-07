package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	// map[key]value
	m := make(map[string]string)
	m["AA"] = "AB"
	m["BB"] = "BC"
	m["CC"] = "CD"
	m["DD"] = "DE"

	m["AA"] = "ZZ"

	fmt.Printf("AA -> %s\n", m["AA"])
	fmt.Printf("BB -> %s\n", m["BB"])
	fmt.Println()
	// map 순회
	prd := make(map[int]Product)
	prd[16] = Product{"Pen", 1000}
	prd[200] = Product{"Coffee", 4100}
	prd[42] = Product{"Pizza", 12800}
	prd[302] = Product{"Phone", 999000}

	for k, v := range prd {
		fmt.Println(k, v)
	}

	// map 요소삭제
	delete(prd, 200)
	fmt.Println()
	for k, v := range prd {
		fmt.Println(k, v)
	}

	// map 요소 존재여부 확인
	v, ok := prd[200]
	fmt.Println(v, ok)
}
