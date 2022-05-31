package main

import "fmt"

func main() {
	value := 12
	pointer := &value
	fmt.Println("The pointer is ", *pointer)
	*pointer = *pointer / 12
	fmt.Println("The pointer is ", *pointer)
	fmt.Println("The value is ", value)
}
