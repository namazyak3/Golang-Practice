package main

import "fmt"

type Student struct {
	lastname, firstname string
	age                 int
	gender              string
}

func main() {

	fmt.Println("---Pattern1---")
	var player1 Student
	player1.lastname = "佐藤"
	player1.firstname = "太郎"
	player1.age = 25
	player1.gender = "男性"
	fmt.Println(player1)

	fmt.Println("\n---Pattern2---")
	player2 := Student{lastname: "ロアンヌ", age: 12, gender: "不明"}
	fmt.Println(player2)

	fmt.Println("\n---Pattern3---")
	player3 := Student{"山田", "守", 18, "男性"}
	fmt.Println(player3)
}
