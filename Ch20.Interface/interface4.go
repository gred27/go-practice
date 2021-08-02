package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
}

func (s *Student) String() string {
	return "Student"
}

type Actor struct {
}

func (a *Actor) String() string {
	return "Actor"
}

func ConvertType(stringer Stringer) {
	// 런타임 에러 :  *Student 타입 Stringer 인터페이스로 쓸 수 있지만 stringer타입이 *Actor 라 에러 발생
	// student := stringer.(*Student)

	// 런타임 에러를 안내기 위해서 변환 되는지 확인후 메서드 실행
	// 타입 변환한 결과, 변환성공여부
	if student, ok := stringer.(*Student); ok {
		fmt.Println(student, ok)
	} else {
		fmt.Println("You aren't Student")
	}

}

func main() {
	actor := &Actor{}
	// *Actor 구조체 값을 인수로 넘김
	ConvertType(actor)

	// *Student 구조체 값을 인수로 넘김
	student := &Student{}
	ConvertType(student)

}
