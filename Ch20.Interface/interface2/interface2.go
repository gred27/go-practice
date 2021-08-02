package main

import (
	"go-practice/Ch20.interface/fedex"
	"go-practice/Ch20.interface/koreaPost"
)

// Sender 인터페이스 정의
// Send 메서드만 가지고 있으면 Sender로 쓰겠다.
type Sender interface {
	Send(parcel string)
}

// SendBook 함수
// Sender를 인자로 받으면 fedex, koreapost 둘다 사용 가능
// SendBook 함수는 Sender가 fedex인지 koreaPost인지 알 필요가 없다
func SendBook(name string, sender Sender) {
	sender.Send(name)
}

/*
	덕타이핑
	1. 타입 선언시 인터페이스에 정의된 메서드만 포함되어있으면 사용가능.
	ex) 자바의 경우 class A implement MyinterFace { } 와 같이 interface 구현여부 선언
	2. 서비스제공자가 인터페이스 정의 할 필요 없이 구체화된 객체만 제공, 서비스 사용자는 사용할 인터페이스만 찾아서 타입 구현
	3. 서비스 제공자가 사용자 니즈에 따라 일일히 인터페이스를 수정할 필요 없다.

*/
func interface2() {
	// sender := &fedex.FedexSender{}
	fedexSender := &fedex.FedexSender{}
	SendBook("어린왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)

	koreaPostSender := &koreaPost.PostSender{}

	SendBook("반지의제왕", koreaPostSender)
	SendBook("데미안", koreaPostSender)
}
