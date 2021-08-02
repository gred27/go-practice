package main

import "fmt"

/*
	인터페이스를 포함한 인터페이스
	1. Read(), Write(), Close() 메서드를 가진 타입 : Reader, Writer, ReadWriter 인터페이스 사용가능
	2. Read(), Close() : Reader만 사용가능
	3. Write(), Close() : Writer만 사용가능
	4. Read(), Write() : 모두 사용 불가능
*/

type Reader interface {
	Read() (n int, err error)
	Close() error
}

type Writer interface {
	Read() (n int, err error)
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}

// 빈 인터페이스를 인수로 사용, 모든 타입을 인수로 받을수 있음
func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type: %T:%v\n", t, t)

	}
}

type Student struct {
	Age int
}

type IStudent interface {
	getAge()
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(Student{15})

	var s IStudent // 기본값 nil
	s.getAge()     // s = nil 이므로 런타임에러
}
