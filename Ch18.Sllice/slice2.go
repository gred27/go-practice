package main

import "fmt"

func ChangeArray(array2 [5]int) {
	array2[2] = 200
}

func ChangeSlice(slice []int) {
	slice[2] = 200
}
func slice2() {
	// array와 slice 비교
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	ChangeArray(array)
	ChangeSlice(slice)
	/*
		Array
		1. array를 ChangeArray의 인수로 넘길 때 array의 모든 값이 복사되어 함수로 전달
		1-1. array의 타입은 [5]int, 넘긴 인수는 40byte
		2. 새로운 메모리에 복사된 array2[2]의 값을 바꾸기 때문에 기존 array의 값은 그대로

		Slice
		1. slice는 구조체, ChangeSlice에 인수로 넘기면 구조체값이 복사 되어 전달
		1-1. slice 타입은 []int, {pointer: p, len: int, cap:int} 24byte
		2. slice2의 포인터가 가리키는 실제 배열은 slice의 포인터가 가리키는 배열과 동일하므로 slice의 값이 바뀜
	*/
	fmt.Println(array)
	fmt.Println(slice)

	// slice 사용시 주의 할 점
	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5)
	/*
		1. slice1 선언
		2. slice2에 4, 5추가
		3. slice2[2] 값을 바꾸면 slice1[2]의 값도 바뀜
		4. append(slice1,200)을 하면 slice2[4] 값도 바뀜
	*/
	fmt.Println(slice1)
	fmt.Println(slice2)
	slice2[2] = 100
	fmt.Println(slice1, slice2)
	slice1 = append(slice1, 500)
	fmt.Println(slice1, slice2)

	/*
		1. slice의 남은 빈공간이 있는지 계산. 빈공간 = cap - len
		2-1. 있으면 배열 빈공간에 값 추가 후 len 값 증가
		2-2. 없으면 cap을 두배로 늘리고 새로운 배열을 가리키는 슬라이스 구조체를 반환 값 추가
		3. 새로 만들어진 slice와 기존 slice는 다른 배열을 가르키는 포인터를 가지고 있게 된다.
		4. append 하던 배열값을 바꾸던 서로 영향을 끼치지 않음
	*/
	slice3 := []int{1, 2, 3}
	slice4 := append(slice3, 4, 5)
	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))
	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))

	slice3[1] = 200
	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))
	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))

	slice3 = append(slice3, 500)
	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))
	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))
}
