const Y int = 3

func main() {
	x := 5
	// 변수로 길이 선언시 에러 발생
	a := [x]int{1, 2, 3, 4, 5}

	b := [Y]int{1, 2, 3}
	var c [10]int
}