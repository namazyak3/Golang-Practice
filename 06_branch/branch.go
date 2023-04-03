package main

import "fmt"

func main() {

	age := 20

	fmt.Println("---if, else---")
	if age >= 20 {
		fmt.Println("adult")
	} else {
		fmt.Println("child")
	}

	fmt.Println("\n---swich, case---")
	switch {
	case age >= 20:
		fmt.Println("adult")
	default:
		fmt.Println("child")
	}
}
