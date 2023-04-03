package main

import "fmt"

type Student struct {
	name string
	rank int
}

func (s Student) sayHello() {
	fmt.Println(s.name, " 入学年度:", (2023 - s.rank), "")
}

func (s Student) whatName() string {
	return s.name
}

func main() {
	player := Student{"佐藤", 2}
	player.sayHello()

	fmt.Println(player.whatName())
}
