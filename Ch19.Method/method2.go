package main

import "fmt"

type account1 struct {
	balance   int
	firstName string
	lastName  string
}

// 포인터 메서드
func (a1 *account1) withdrawPointer(amount int) {
	a1.balance -= amount
}

// 값타입 메서드
func (a2 account1) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a3 account1) withdrawReturnValue(amount int) account1 {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account1 = &account1{100, "ji", "Park"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(20)
	fmt.Println(mainA.balance)

	mainB := mainA.withdrawReturnValue(20)
	mainB.withdrawReturnValue(30)
	fmt.Println(mainB.balance)

}
