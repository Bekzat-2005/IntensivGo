package main

import "fmt"

type Speaker interface {
	Speak()
}

type Human struct {
	Name string
}
type Animal struct {
	Name string
}

func (h Human) Speak() {
	fmt.Println(h.Name, "сказал: Привет!")
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "Gav Gav")
}

func MakeItSpeak(s Speaker) {
	s.Speak()
}

func main() {
	h := Human{Name: "Бекзат"}
	a := Animal{Name: "Rex"}
	MakeItSpeak(h)
	MakeItSpeak(a)

}
