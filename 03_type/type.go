package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 数値
	num1 := 1
	var num2 int = 2
	num3 := 3.4
	var num4 float64 = 5.6

	// 文字列
	string1 := "文字列1"
	var string2 string = "文字列2"

	// ブール値型
	bool1 := true
	var bool2 bool = false

	fmt.Println("---数値---")
	fmt.Println(reflect.TypeOf(num1))
	fmt.Println(reflect.TypeOf(num2))
	fmt.Println(reflect.TypeOf(num3))
	fmt.Println(reflect.TypeOf(num4))
	fmt.Println("\n---文字列---")
	fmt.Println(reflect.TypeOf(string1))
	fmt.Println(reflect.TypeOf(string2))
	fmt.Println("\n---ブール値列---")
	fmt.Println(reflect.TypeOf(bool1))
	fmt.Println(reflect.TypeOf(bool2))
}
