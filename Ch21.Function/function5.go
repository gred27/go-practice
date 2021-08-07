package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writeHello(writer Writer) {
	writer("Hello, Go wolrd!")
}

func main() {
	f, err := os.Create("test1.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer f.Close()

	// msg를 파일에 출력하는 함수리터럴 을 writeHello의 인수로 넘김
	// writeHello는 함수리터럴을 "Hello, go world"를 msg 인수로 호출, 파일에 출력
	// writeHello 외부에서 파일을 출력하는 로직의 함수를 주입 => 외부에서 의존성을 주입
	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})

	// Writer 타입의 함수리터럴은 다 writeHello에서 실행 가능
	writeHello(func(msg string) {
		fmt.Println(msg)
	})

}
