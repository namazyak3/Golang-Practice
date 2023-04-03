package main

import (
	"fmt"
)

func main() {

	// 1次元配列
	array1 := [...]string{"apple", "grape", "orange"}

	// 多次元配列
	array2 := [2][3][4]int{
		{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}},
		{{12, 13, 14, 15}, {16, 17, 18, 19}, {20, 21, 22, 23}},
	}

	fmt.Println(array1)
	fmt.Println(array1[1])
	fmt.Println(array2[0][2])
	fmt.Println(array2[1][2][1])
}
