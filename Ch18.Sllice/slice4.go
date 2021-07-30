package main

import (
	"fmt"
	"sort"
)

func slice4() {

	// 슬라이스 복제
	slice := []int{1, 2, 3, 4, 5}
	slice1 := append([]int{}, slice...)
	slice2 := make([]int, 10)
	copy1 := copy(slice2, slice)

	slice1[2] = 500
	fmt.Println(slice, slice1, slice2, copy1)

	// 슬라이스 요소 삭제
	idx := 2
	slice2 = append(slice2[:idx], slice2[idx+1:]...)

	// 슬라이스 요소 추가
	idx = 3
	slice3 := []int{1, 2, 3, 4, 5, 6}
	slice3 = append(slice3[:idx], append([]int{100}, slice3[idx:]...)...)

	// 슬라이스 요소 추가2
	slice4 := []int{12, 35, 47, 51, 16, 10}
	slice4 = append(slice4, 0)
	copy(slice4[idx+1:], slice4[idx:])
	slice4[idx] = 200

	fmt.Println(slice3, slice4)

	// 슬라이스 요소 정렬
	sort.Ints(slice4)
	fmt.Println(slice4)

}
