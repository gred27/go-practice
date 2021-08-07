// 함수 리터럴
package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		// 함수 리터럴은 실행 시의 i 값을 참조로 캡쳐하기때문에
		// f[i] 호출시 i 값인 3을 출력
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		// for문 반복시 매번 새로운 v 변수를 생성
		// 함수리터럴은 서로다른 v변수를 캡쳐
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}
	for i := 0; i < 3; i++ {
		// 함수 실행시 각각의 v 값을 출력
		f[i]()
	}
}

func function4() {
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

	// 함수리터럴 내부 상태 주의점
	CaptureLoop()  // 3 3 3
	CaptureLoop2() // 1 2 3

}
