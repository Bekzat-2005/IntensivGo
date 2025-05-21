package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) SayHello() {
	fmt.Println("Hello my name is ", p.name, p.age)
}

func (p *Person) birthday() {
	p.age += 1
}

func main() {

	p := Person{"Bekzat", 19}
	p.SayHello()
	p.birthday()
	p.SayHello()
}
