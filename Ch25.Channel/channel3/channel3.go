// channel은 thread safe 한 queue이다.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	// 채널 닫음
	close(ch)

	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {

	for n := range ch {
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)

	}
	wg.Done()
}

// 마지막에 데드락 발생
// main 은 Wait 중인데 square에서 n은 무한하게 데이터를 뽑아오려고 대기하기때문에
// Done이 실행 안되고 모든 고루틴이 멈춰 데드락이 발생
// 그래서 채널을 닫아줘야한다.
// 채널이 닫히지 않아서 무한대기하는 고루틴을 고루틴 leak 이라고 함.
