package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() { // Attaching methods to structs
	fmt.Println("sounding", d.Sound)
}

func main() {
	poodle := Dog{"Poodle", 10, "Woof"}
	// fmt.Println(poodle)
	poodle.Speak()
	poodle.Sound = "arf"
	poodle.Speak()
}
