package main

import "fmt"

type User struct {
	Name string
	Id   string
	Age  int
}

// 내장 타입 사용 구조체
type VipUser struct {
	UserInfo User
	VipLever int
	Price    int
}

// 포함된 필드
type VVipUser struct {
	User
	VipLevel int
	VVip     bool
}

func struct2() {
	user := User{"Moon", "moon", 24}
	vip := VipUser{
		User{"Kim", "kim", 26},
		2,
		10,
	}
	vvip := VVipUser{
		User{"Seol", "sera", 28},
		1,
		true,
	}
	fmt.Println(user, vip)
	fmt.Println(vvip.Name, vvip.Id, vvip.VipLevel)
}
