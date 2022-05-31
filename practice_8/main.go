package main

import "fmt"

func main() {
	doSomething()
	fmt.Println("The sum is", addValues(5, 197))

}

func doSomething() {
	fmt.Println("Enter doSomething")
}

func addValues(value1 int, value2 int) int { // addValues(value1, value2 int) int
	return value1 + value2
}
