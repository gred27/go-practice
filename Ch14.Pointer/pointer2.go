package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func pointer2() {
	// 포인터 사용 예제
	var data Data
	ChangeData(&data)
	fmt.Println(data)

	// 구조체 포인터변수 초기화
	// 1. 구조체 변수 생성
	// 2. 구조체 포인터 변수 생성 후 대입
	var data1 Data
	var p1 *Data = &data1

	// 1. 구조체 포인터 변수 생성하고 바로 주소 대입
	var p2 *Data = &Data{}

	fmt.Println(p1, p2)
}
