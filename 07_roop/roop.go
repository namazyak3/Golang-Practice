package main

import "fmt"

func main() {

	fmt.Println("---Pattern1---")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("---Pattern2---")
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("---Pattern3---")
	j := 1
	for {
		fmt.Println(j)
		if j == 10 {
			break
		}
		j++
	}
}
