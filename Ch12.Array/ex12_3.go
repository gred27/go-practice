package main

import "fmt"

func ex12_3() {
	nums := [...]int{10, 20, 30, 40, 50}

	nums[2] = 300

	// for문 배열 순회
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	// range 배열 순회
	for i, v := range nums {
		fmt.Println(i, v)
	}

	// index 필요 없을 때
	for _, v := range nums {
		fmt.Println(v)
	}
}
