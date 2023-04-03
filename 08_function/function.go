package main

import "fmt"

func textRes(text string) {
	fmt.Println(text)
}

func dubleRes(num int) {
	fmt.Println(num * 2)
}

func main() {

	fmt.Println("---Text response---")
	textRes("called res()")

	fmt.Println("\n---Duble response---")
	dubleRes(4)

	fmt.Println("\n---Function in variable---")
	x := 2
	y := 3
	num := func() int {
		return x * y
	}
	fmt.Println(num())

	fmt.Println("\n---Nameless functioon---")
	func(num int) {
		fmt.Println(num * 2)
	}(5)
}
