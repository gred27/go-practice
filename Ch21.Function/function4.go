// 함수 리터럴
package main

import "fmt"

func main() {
	/*
		- 함수 리터럴은 내부 상태를 가진다
		- 함수리터럴 내부에서 사용되는 외부변수는 함수 내부 상태로 저장된다.

	*/

	i := 0
	fmt.Println("i 선언", i)
	f := func() {
		fmt.Println("함수 시작", i)
		// 여기서 i 의 값은 함수 외부에서 정의한 0
		// i 는 함수 호출 시점의 값으로 사용
		i += 10
		fmt.Println("함수 종료", i)
	}

	i++
	fmt.Println("i+1", i)

	f()
	fmt.Println("최종 결과", i)
}
