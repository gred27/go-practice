package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("Hi, I'm %d old, %s.", s.Age, s.Name)
}

func main() {
	student := Student{"Hong", 28}
	var stringer Stringer

	/*
		string을 반환하는 String() 메서드를 포함한 모든 타입은 Stringer interface로 사용 가능
		Student 구조체가 Stringer 인터페이스를 구현
	*/

	stringer = student
	fmt.Printf("%s\n", stringer.String())
}