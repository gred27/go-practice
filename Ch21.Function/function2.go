/*
	defer
	- 함수 종료 전에 실행을 보장
	- 함수를 미리 쓰지만 함수 종료 전에 실행됨
	- os 자원 반납할 때 주로 사용
*/

package main

import (
	"fmt"
	"os"
)

func function2() {

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("failed to create afile,", err)
		return
	}

	// 함수 종료 전으로 지연됨
	// 1. 파일 출력
	// 2. Hello Go wolrd 입력
	// 3. 반드시 실행됩니다.

	defer fmt.Println("반드시 호출됨")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다.")

	fmt.Println("파일 출력")
	fmt.Fprintln(f, "Hello, Go world!")
}
