package main

import "fmt"

type account struct {
	balance int
}

type money int

// function
func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// method
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

// 별칭 타입을 리시버로 갖는 method
func (a money) add(b int) int {
	return int(a) + b
}

func method1() {
	a := &account{100}

	withdrawFunc(a, 30)

	a.withdrawMethod(30)

	// 동작은 똑같지만 Method는 struct에 종속
	fmt.Printf("%d \n", a.balance)

	// 별칭타입
	var myMoney money = 10
	fmt.Println(myMoney.add(30))
	var b int = 20
	// 타입변환
	fmt.Println(money(b).add(50))

}
