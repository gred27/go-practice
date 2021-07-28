import "fmt"

func ex12_4() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
	}
	// 이중배열 출력
	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
	}
}
