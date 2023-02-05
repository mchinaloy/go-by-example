package main

import "fmt"

type Person interface {
	speak()
}

type Employee struct {
}

func main() {
	emp := Employee{}
	for_all_those_who_can_speak(emp)
}

func for_all_those_who_can_speak(p Person) {
	p.speak()
}

func (p Employee) speak() {
	fmt.Println("I am a Person speaking!")
}
