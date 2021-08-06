package main

import "fmt"

// sum 가변인수 nums를 인수로 받는다
// nums의 타입은 slice []int 타입
func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입 :%T\n", nums)

	for _, v := range nums {
		sum += v
	}

	return sum
}

func function1() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())

	// Println은 인자로 빈 인터페이스 (arg ...interface{}) 을 받아서 각각의 타입별로 출력
	fmt.Println(1, 2, 3, 4, "hello", 3.14)
}
