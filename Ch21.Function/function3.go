/* 함수타입 변수 (함수 포인터)
- 함수를 값으로 가진 변수
- 함수 시작지점 메모리 주소를 가리킴
*/

package main

import "fmt"

type opFunc func(int, int) int

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

// 함수 시그니쳐 : 함수 타입을 func (int, int) int로 정함
// getOperator : string을 인수로 받아 func (int, int) int 타입의 함수를 반환
func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else if op == "-" {
		// 함수 리터럴 (람다) : 이름 없는 함수, 함수 타입 변수로만 호출 가능
		return func(a, b int) int {
			return a - b
		}
	} else {
		return nil
	}
}

func function3() {
	// 함수 시그니쳐 :  함수이름과 구현을 제외
	// operator : int 타입 인수 2개를 받아 int 타입을 반환하는 함수타입 변수
	var operator func(int, int) int
	operator = getOperator("*")

	var result = operator(3, 4)
	fmt.Println(result)

	// 함수 타입쓰면 코드가 길어지므로 별칭으로 함수 정의를 줄여쓴다
	var operator1 opFunc
	operator1 = getOperator("+")

	result1 := operator1(3, 4)
	fmt.Println(result1)

	fn := getOperator("-")
	result2 := fn(4, 2)
	fmt.Println(result2)

	// 함수 리터럴사용
	fn1 := func(a, b int) {
		fmt.Println(a, b)
	}
	fn1(5, 6)

	// 또는
	func(a, b int) {
		fmt.Println(a, b)
	}(3, 4)

}
