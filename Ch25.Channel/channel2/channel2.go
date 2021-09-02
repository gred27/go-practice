// channel은 thread safe 한 queue이다.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// ch := make(chan int)
	ch := make(chan int, 2) // 버퍼를 주면 데이터 한개 놓고 종료

	wg.Add(1)
	go square(&wg, ch)
	ch <- 9 // 채널에 데이터 넣음

	// 채널에서 데이터를 빼가는 게 없으면 무한 대기
	fmt.Println("Never print")
}

func square(wg *sync.WaitGroup, ch chan int) {
	for {
		time.Sleep(time.Second)
		fmt.Println("sleep")
	}
}
