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

	var rest int = 1000

	for {
		fmt.Printf("1~5 사이의 숫자 입력 : ")
		input, err := InputIntValue()

		if err != nil {
			fmt.Println("잘못 입력했습니다. 다시 입력하세요.")
			continue
		} else if input < 1 || input > 5 {
			fmt.Println("1~5사이의 값을 입력하세요.")
			continue
		} else {
			r := rand.Intn(5) + 1

			if r == input {
				rest += 500
				fmt.Printf("축하합니다. 현재 잔액 : %d\n", rest)
			} else {
				rest -= 100
				fmt.Printf("아쉽네요. 현재 잔액 : %d\n", rest)
			}
		}

		if rest >= 5000 || rest <= 0 {
			fmt.Println("Game over")
			break
		}
	}

}
