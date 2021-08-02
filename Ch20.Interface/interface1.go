package main

import "fmt"

/*
	interface도 타입
	구현이 빠진 메소드를 정의
	메서드명이 중복되면 안된다.(overriding X)
*/
type Stringer1 interface {
	String() string
}

type Student1 struct {
	Name string
	Age  int
}

func (s Student1) String() string {
	return fmt.Sprintf("Hi, I'm %d old, %s.", s.Age, s.Name)
}

func interface1() {
	student := Student1{"Hong", 28}
	var stringer Stringer1

	/*
		string을 반환하는 String() 메서드를 포함한 모든 타입은 Stringer interface로 사용 가능
		Student 구조체가 Stringer 인터페이스를 구현
	*/

	stringer = student
	fmt.Printf("%s\n", stringer.String())
}
