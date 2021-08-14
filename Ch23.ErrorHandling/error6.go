// 패닉 전파 및 복구
package main

import "fmt"

func f() {
	fmt.Println("f() 함수 시작")

	defer func() {

		// recover은 호출 순서를 역순으로 전파하기떄문에 주의해서 사용해야함
		// ex 파일에 데이터 기록하고 recover 발생하면 파일에 이상한 데이터가 저장된 상태로 남음.
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r)
		}
	}()

	g()
	fmt.Println("f() 함수 종료")
}

func g() {
	fmt.Printf("9/3 = %d\n", h(9, 3))
	fmt.Printf("3/0 = %d\n", h(3, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("b는 0일 수 없습니다.")
	}

	return a / b
}

func main() {
	f()
	fmt.Println("프로그램 실행 됨")
}
