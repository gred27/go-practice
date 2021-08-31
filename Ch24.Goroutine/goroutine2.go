package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지의 합계는 %d입니다.\n", a, b, sum)
	wg.Done() //작업 완료
}

func main() {
	wg.Add(10) // 작업개수 10개로 설정
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000) // 고루틴 실행
	}

	wg.Wait() // 작업 끝날때까지 대기
	fmt.Println("end")
}
