package main

import "fmt"

func slice3() {
	array := [5]int{1, 2, 3, 4, 5}
	// 배열 슬라이싱
	// [시작인덱스 : 끝인덱스 : 최대인덱스]
	// 배열의 2번째 요소를 가르키는 포인터 + len 2짜리 슬라이스
	slice := array[1:3]

	// append 하면 기존 배열도 값 바뀜
	slice = append(slice, 200)
	fmt.Println(array, slice)

	slice2 := []int{1, 2, 3, 4, 5}
	// 처음부터 슬라이싱
	slice3 := slice2[0:3]
	// 끝까지 슬라이싱
	slice4 := slice2[2:]
	// 전체 슬라이싱
	slice5 := slice2[:]
	// cap 조절
	slice6 := slice2[1:3:4]

	fmt.Println(slice3, slice4, slice5, slice6)

}
