package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		// 잘못 입력시 입력 스트림을 비움
		stdin.ReadString('\n')
	}

	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(100)
	var cnt int = 0

	for {
		fmt.Printf("숫자값 입력:")
		n, err := InputIntValue()

		if err != nil {
			fmt.Println("숫자만 입력하세요")
		} else {
			fmt.Println("입력하신 숫자는 ", n, " 입니다.")

			if n > r {
				fmt.Println("입력값이 큽니다.")
			} else if n < r {
				fmt.Println("입력값이 작습니다.")
			} else {
				fmt.Printf("정답입니다. 시도횟수 : %d\n", cnt)
				break
			}
		}
		cnt++

	}

}
