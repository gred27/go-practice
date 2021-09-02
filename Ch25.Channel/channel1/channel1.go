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
	ch <- 9 // 채널에 데이터 넣음
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch // 채널에서 데이터 빼옴

	time.Sleep(time.Second)
	fmt.Printf("Square: %d\n", n*n)
	wg.Done()
}
