package main

import (
	"fmt"
)

func main() {

	// 変数宣言
	x := 1
	y := 10
	z := 100

	// 計算
	result1 := x + y     // 1 + 10
	result2 := z / y     // 100 / 10
	result3 := z - x + y // 100 - 1 + 10
	y += x               // y = 10 + 1
	x++                  // x = 1 + 1
	z--                  // z = 100 -1

	// 出力
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(y)
	fmt.Println(x)
	fmt.Println(z)
}
