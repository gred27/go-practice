// panic
// 문제 발생 시점에 프로그램을 바로 종료하고 에러 메세지출력

package main

import "fmt"

func devide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다.")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func error5() {
	devide(9, 3)
	devide(9, 0)
}
