package main

import "fmt"

func main() {

	// 変数宣言
	age := 20

	// if, else
	fmt.Println("---if, else---")
	if age >= 20 {
		fmt.Println("adult")
	} else {
		fmt.Println("child")
	}

	// swich, case
	fmt.Println("\n---swich, case---")
	switch {
	case age >= 20:
		fmt.Println("adult")
	default:
		fmt.Println("child")
	}
}
