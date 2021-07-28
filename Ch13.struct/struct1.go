package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func struct1() {
	// 구조체 초기화
	var house House
	house.Address = "Address"
	house.Size = 28
	house.Price = 9.8
	house.Type = "Apart"

	fmt.Println("Address: ", house.Address)
	fmt.Printf("Size: %d\n", house.Size)

	// 구조체 초기화2
	var house2 House = House{"주소", 90, 10.5, "단독"}
	fmt.Print(house2)

	// 구조체 일부만 초기화
	var house3 House = House{Size: 30, Price: 5.5}
	fmt.Print(house3)

}
