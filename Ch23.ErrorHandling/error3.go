package main

import "fmt"

type PassWordError struct {
	Len        int
	RequireLen int
}

func (err PassWordError) Error() string {
	return "암호길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PassWordError{len(password), 8}
	}

	return nil
}

func error3() {
	err := RegisterAccount("myId", "myPw")

	if err != nil {
		if errInfo, ok := err.(PassWordError); ok {
			fmt.Printf("%v Len : %d RequireLen : %d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("화원가입됐습니다.")
	}
}
