package main

import "fmt"

// 배열 기본
func ex12_1() {
	var t [5]float64 = [5]float64{24.0, 22.1, 3.2, 5.7, 98.6}

	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}

}
