package main

import "fmt"

func slice2() {
	// 빈 슬라이스
	var slice1 []int
	fmt.Println(slice1)

	// 슬라이스 초기화
	// 배열선언은 var array = [...]int{1,2,3}
	var slice2 = []int{1, 2, 3}
	var slice3 = []int{1, 5: 1, 10: 1}
	fmt.Println(slice2)
	fmt.Println(slice3)

	// 슬라이스 초기화2
	// make(slice_pointer, slice_length, slice_cap)
	var slice4 = make([]int, 3)
	fmt.Println(slice4, slice4[2])

	// 슬라이스 순회
	var slice5 = []int{1, 2, 3, 4}
	for i, v := range slice5 {
		slice5[i] = v + 10
	}
	fmt.Println(slice5)

	// 슬라이스 값 추가
	// 값를 추가한 새로운 슬라이스 반환
	slice6 := append(slice5, 15)
	fmt.Println(slice6)
	// 값 여러개 추가
	slice7 := append(slice6, 16, 17, 18, 19, 20)
	fmt.Println(slice7)
}
